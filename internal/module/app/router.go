package app

import (
	"io/fs"
	"net/http"
	"path"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/usesnipet/go-template/config"
	errorhandler "github.com/usesnipet/go-template/internal/module/app/error-handler"
	"github.com/usesnipet/go-template/web"
)

type HandlerFunc = errorhandler.HandlerFunc

func NewRouter(cfg *config.Config) (http.Handler, chi.Router, func(HandlerFunc) http.HandlerFunc, error) {
	builder := errorhandler.NewErrorHandlerBuilder()
	builder.AddMapper(errorhandler.GormMapper)
	serve := builder.Serve()

	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type"},
	}))
	r.Use(responseTime)
	r.Use(middleware.Compress(1, "text/html"))

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
	))

	api := chi.NewRouter()
	r.Mount(config.APIPrefix, api)

	dist, err := fs.Sub(web.Dist, "dist")
	if err != nil {
		return nil, nil, nil, err
	}

	spa := newSPAHandler(dist)
	r.Get("/", spa.ServeHTTP)
	r.Handle("/*", spa)

	return r, api, serve, nil
}

func newSPAHandler(fsys fs.FS) http.Handler {
	fileServer := http.StripPrefix("/", http.FileServer(http.FS(fsys)))

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, config.APIPrefix) {
			http.NotFound(w, r)
			return
		}

		name := strings.TrimPrefix(path.Clean(r.URL.Path), "/")
		if name == "" || name == "." {
			http.ServeFileFS(w, r, fsys, "index.html")
			return
		}

		if _, err := fs.Stat(fsys, name); err == nil {
			fileServer.ServeHTTP(w, r)
			return
		}

		http.ServeFileFS(w, r, fsys, "index.html")
	})
}

type responseTimeWriter struct {
	http.ResponseWriter
	start time.Time
	path  string
}

func (w *responseTimeWriter) WriteHeader(statusCode int) {
	if !strings.HasSuffix(w.path, ".html") {
		w.Header().Set("X-Response-Time", time.Since(w.start).String())
	}
	w.ResponseWriter.WriteHeader(statusCode)
}

func responseTime(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(&responseTimeWriter{
			ResponseWriter: w,
			start:          time.Now(),
			path:           r.URL.Path,
		}, r)
	})
}
