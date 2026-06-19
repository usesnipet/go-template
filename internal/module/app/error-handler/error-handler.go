package errorhandler

import (
	"net/http"

	"github.com/usesnipet/go-template/internal/api"
)

type HandlerFunc func(http.ResponseWriter, *http.Request) error

type ErrorHandlerBuilder struct {
	mappers []func(err error) (error, bool)
}

func (b *ErrorHandlerBuilder) AddMapper(mapper func(err error) (error, bool)) {
	b.mappers = append(b.mappers, mapper)
}

func (b *ErrorHandlerBuilder) mapError(err error) error {
	for _, mapper := range b.mappers {
		if mapped, ok := mapper(err); ok {
			return mapped
		}
	}
	return api.NewHTTPError(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
}

func (b *ErrorHandlerBuilder) Serve() func(HandlerFunc) http.HandlerFunc {
	return func(h HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			if err := h(w, r); err != nil {
				api.WriteError(w, b.mapError(err))
			}
		}
	}
}

func NewErrorHandlerBuilder() *ErrorHandlerBuilder {
	return &ErrorHandlerBuilder{
		mappers: []func(err error) (error, bool){},
	}
}
