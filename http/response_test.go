package http_test

import (
	"net/http"
	"testing"

	pkghttp "github.com/goauthnz/pkg/http"
)

func TestNewHTTPResponse_ErrorFlagForSuccessCodes(t *testing.T) {
	successCodes := []int{
		http.StatusOK,
		http.StatusCreated,
		http.StatusAccepted,
		http.StatusNoContent,
	}
	for _, code := range successCodes {
		resp := pkghttp.NewHTTPResponse(code, "OK", nil)
		if resp.Status.Error {
			t.Errorf("code %d: expected Error=false, got true", code)
		}
	}
}

func TestNewHTTPResponse_ErrorFlagForErrorCodes(t *testing.T) {
	errorCodes := []int{
		http.StatusBadRequest,
		http.StatusUnauthorized,
		http.StatusNotFound,
		http.StatusConflict,
		http.StatusGone,
		http.StatusInternalServerError,
	}
	for _, code := range errorCodes {
		resp := pkghttp.NewHTTPResponse(code, "error", nil)
		if !resp.Status.Error {
			t.Errorf("code %d: expected Error=true, got false", code)
		}
	}
}

func TestNewHTTPResponse_FieldsSet(t *testing.T) {
	data := map[string]string{"id": "acnt_01J"}
	resp := pkghttp.NewHTTPResponse(http.StatusOK, pkghttp.MessageSuccess, data)

	if resp.Status.Code != http.StatusOK {
		t.Errorf("Code = %d, want %d", resp.Status.Code, http.StatusOK)
	}
	if resp.Status.Message != pkghttp.MessageSuccess {
		t.Errorf("Message = %q, want %q", resp.Status.Message, pkghttp.MessageSuccess)
	}
	if resp.Data == nil {
		t.Error("expected Data to be set")
	}
}
