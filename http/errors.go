package http

import (
	"context"
	"net/http"

	pkgerrors "github.com/goauthnz/pkg/errors"
)

// TranslateError translates an error to an HTTP status code and a response.
func TranslateError(ctx context.Context, err error) (int, any) {
	switch {
	case pkgerrors.IsNotFoundError(err):
		return http.StatusNotFound, NewHTTPResponse(http.StatusNotFound, err.Error(), nil)
	case pkgerrors.IsResourceAlreadyCreatedError(err):
		return http.StatusConflict, NewHTTPResponse(http.StatusConflict, err.Error(), nil)
	case pkgerrors.IsBadRequestError(err):
		return http.StatusBadRequest, NewHTTPResponse(http.StatusBadRequest, err.Error(), nil)
	case pkgerrors.IsUnauthorizedError(err):
		return http.StatusUnauthorized, NewHTTPResponse(http.StatusUnauthorized, err.Error(), nil)
	default:
		return http.StatusInternalServerError, NewHTTPResponse(http.StatusInternalServerError, MessageInternalServerError, nil)
	}
}
