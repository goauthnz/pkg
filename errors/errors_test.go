package errors_test

import (
	stderrors "errors"
	"testing"

	"github.com/goauthnz/pkg/errors"
)

func TestErrorKey(t *testing.T) {
	key := "some.resource.key"
	cases := []struct {
		name    string
		err     error
		wantMsg string
	}{
		{"NotFound", errors.NewNotFoundError(key), key},
		{"BadRequest", errors.NewBadRequestError(key), key},
		{"ExpiredResource", errors.NewExpiredResourceError(key), key},
		{"InternalServer", errors.NewInternalServerError(key), key},
		{"Unauthorized", errors.NewUnauthorizedError(key), key},
		{"ResourceAlreadyCreated", errors.NewResourceAlreadyCreatedError(key), key},
		{"OutdatedResource", errors.NewOutdatedResourceError(key), key},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.err == nil {
				t.Fatal("expected non-nil error")
			}
			if tc.err.Error() != tc.wantMsg {
				t.Errorf("Error() = %q, want %q", tc.err.Error(), tc.wantMsg)
			}
		})
	}
}

func TestIsErrorFunctions_MatchOwnType(t *testing.T) {
	cases := []struct {
		name  string
		err   error
		check func(error) bool
	}{
		{"IsNotFoundError", errors.NewNotFoundError("k"), errors.IsNotFoundError},
		{"IsBadRequestError", errors.NewBadRequestError("k"), errors.IsBadRequestError},
		{"IsExpiredResourceError", errors.NewExpiredResourceError("k"), errors.IsExpiredResourceError},
		{"IsInternalServerError", errors.NewInternalServerError("k"), errors.IsInternalServerError},
		{"IsUnauthorizedError", errors.NewUnauthorizedError("k"), errors.IsUnauthorizedError},
		{"IsResourceAlreadyCreatedError", errors.NewResourceAlreadyCreatedError("k"), errors.IsResourceAlreadyCreatedError},
		{"IsOutdatedResourceError", errors.NewOutdatedResourceError("k"), errors.IsOutdatedResourceError},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if !tc.check(tc.err) {
				t.Errorf("%s returned false for its own error type", tc.name)
			}
		})
	}
}

func TestIsErrorFunctions_RejectOtherTypes(t *testing.T) {
	allErrors := []error{
		errors.NewNotFoundError("k"),
		errors.NewBadRequestError("k"),
		errors.NewExpiredResourceError("k"),
		errors.NewInternalServerError("k"),
		errors.NewUnauthorizedError("k"),
		errors.NewResourceAlreadyCreatedError("k"),
		errors.NewOutdatedResourceError("k"),
	}
	allChecks := []struct {
		name  string
		check func(error) bool
		own   int // index into allErrors that this check owns
	}{
		{"IsNotFoundError", errors.IsNotFoundError, 0},
		{"IsBadRequestError", errors.IsBadRequestError, 1},
		{"IsExpiredResourceError", errors.IsExpiredResourceError, 2},
		{"IsInternalServerError", errors.IsInternalServerError, 3},
		{"IsUnauthorizedError", errors.IsUnauthorizedError, 4},
		{"IsResourceAlreadyCreatedError", errors.IsResourceAlreadyCreatedError, 5},
		{"IsOutdatedResourceError", errors.IsOutdatedResourceError, 6},
	}

	for _, chk := range allChecks {
		for i, err := range allErrors {
			if i == chk.own {
				continue
			}
			if chk.check(err) {
				t.Errorf("%s returned true for error at index %d (wrong type)", chk.name, i)
			}
		}
	}
}

func TestIsErrorFunctions_RejectNilAndPlainErrors(t *testing.T) {
	plain := stderrors.New("plain error")
	checks := []struct {
		name  string
		check func(error) bool
	}{
		{"IsNotFoundError", errors.IsNotFoundError},
		{"IsBadRequestError", errors.IsBadRequestError},
		{"IsExpiredResourceError", errors.IsExpiredResourceError},
		{"IsInternalServerError", errors.IsInternalServerError},
		{"IsUnauthorizedError", errors.IsUnauthorizedError},
		{"IsResourceAlreadyCreatedError", errors.IsResourceAlreadyCreatedError},
		{"IsOutdatedResourceError", errors.IsOutdatedResourceError},
	}
	for _, chk := range checks {
		t.Run(chk.name+"/plain", func(t *testing.T) {
			if chk.check(plain) {
				t.Errorf("%s returned true for a plain stdlib error", chk.name)
			}
		})
	}
}
