package http_test

import (
	stderrors "errors"
	"context"
	"net/http"
	"testing"

	pkghttp "github.com/goauthnz/pkg/http"
	"github.com/goauthnz/pkg/errors"
)

func TestTranslateError_StatusCodes(t *testing.T) {
	ctx := context.Background()
	cases := []struct {
		name       string
		err        error
		wantStatus int
	}{
		{"NotFound", errors.NewNotFoundError("k"), http.StatusNotFound},
		{"ResourceAlreadyCreated", errors.NewResourceAlreadyCreatedError("k"), http.StatusConflict},
		{"BadRequest", errors.NewBadRequestError("k"), http.StatusBadRequest},
		{"Unauthorized", errors.NewUnauthorizedError("k"), http.StatusUnauthorized},
		{"ExpiredResource", errors.NewExpiredResourceError("k"), http.StatusGone},
		{"OutdatedResource", errors.NewOutdatedResourceError("k"), http.StatusConflict},
		{"InternalServer", errors.NewInternalServerError("k"), http.StatusInternalServerError},
		{"UnknownError", stderrors.New("unknown"), http.StatusInternalServerError},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			status, resp := pkghttp.TranslateError(ctx, tc.err)
			if status != tc.wantStatus {
				t.Errorf("status = %d, want %d", status, tc.wantStatus)
			}
			httpResp, ok := resp.(pkghttp.HTTPResponse)
			if !ok {
				t.Fatalf("response is not HTTPResponse: %T", resp)
			}
			if httpResp.Status.Code != tc.wantStatus {
				t.Errorf("response body code = %d, want %d", httpResp.Status.Code, tc.wantStatus)
			}
			if !httpResp.Status.Error {
				t.Error("expected Status.Error = true for error responses")
			}
		})
	}
}

func TestTranslateError_MessagePreserved(t *testing.T) {
	ctx := context.Background()
	key := "user.not.found"
	_, resp := pkghttp.TranslateError(ctx, errors.NewNotFoundError(key))
	httpResp := resp.(pkghttp.HTTPResponse)
	if httpResp.Status.Message != key {
		t.Errorf("message = %q, want %q", httpResp.Status.Message, key)
	}
}

func TestTranslateError_InternalServerErrorMessageObfuscated(t *testing.T) {
	ctx := context.Background()
	_, resp := pkghttp.TranslateError(ctx, stderrors.New("secret db error"))
	httpResp := resp.(pkghttp.HTTPResponse)
	if httpResp.Status.Message == "secret db error" {
		t.Error("internal error message should be obfuscated, not returned verbatim")
	}
	if httpResp.Status.Message != pkghttp.MessageInternalServerError {
		t.Errorf("message = %q, want %q", httpResp.Status.Message, pkghttp.MessageInternalServerError)
	}
}
