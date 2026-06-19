package errorhandler

import (
	"net/http"

	"github.com/usesnipet/go-template/internal/api"
	"gorm.io/gorm"
)

func GormMapper(err error) (error, bool) {
	switch err {
	case gorm.ErrRecordNotFound:
		return api.NewHTTPError(http.StatusNotFound, "record not found"), true
	case gorm.ErrInvalidTransaction,
		gorm.ErrNotImplemented,
		gorm.ErrMissingWhereClause,
		gorm.ErrUnsupportedRelation,
		gorm.ErrPrimaryKeyRequired,
		gorm.ErrModelValueRequired,
		gorm.ErrModelAccessibleFieldsRequired,
		gorm.ErrInvalidData,
		gorm.ErrInvalidDB,
		gorm.ErrInvalidField,
		gorm.ErrInvalidValue:
		return api.NewHTTPError(http.StatusInternalServerError, "internal server error"), true
	default:
		return err, false
	}
}
