package grpc_test

import (
	stderrors "errors"
	"context"
	"testing"

	connectgo "github.com/bufbuild/connect-go"
	pkggrpc "github.com/goauthnz/pkg/grpc"
	"github.com/goauthnz/pkg/errors"
)

var ctx = context.Background()

func TestTranslateToGRPCError_Nil(t *testing.T) {
	if got := pkggrpc.TranslateToGRPCError(ctx, nil); got != nil {
		t.Errorf("expected nil, got %v", got)
	}
}

func TestTranslateFromGRPCError_Nil(t *testing.T) {
	if got := pkggrpc.TranslateFromGRPCError(ctx, nil); got != nil {
		t.Errorf("expected nil, got %v", got)
	}
}

func TestTranslateToGRPCError_Codes(t *testing.T) {
	cases := []struct {
		name     string
		err      error
		wantCode connectgo.Code
	}{
		{"NotFound", errors.NewNotFoundError("k"), connectgo.CodeNotFound},
		{"ResourceAlreadyCreated", errors.NewResourceAlreadyCreatedError("k"), connectgo.CodeAlreadyExists},
		{"BadRequest", errors.NewBadRequestError("k"), connectgo.CodeInvalidArgument},
		{"Unauthorized", errors.NewUnauthorizedError("k"), connectgo.CodeUnauthenticated},
		{"ExpiredResource", errors.NewExpiredResourceError("k"), connectgo.CodeFailedPrecondition},
		{"OutdatedResource", errors.NewOutdatedResourceError("k"), connectgo.CodeAborted},
		{"InternalServer", errors.NewInternalServerError("k"), connectgo.CodeInternal},
		{"UnknownError", stderrors.New("unknown"), connectgo.CodeInternal},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := pkggrpc.TranslateToGRPCError(ctx, tc.err)
			if got == nil {
				t.Fatal("expected non-nil error")
			}
			if connectgo.CodeOf(got) != tc.wantCode {
				t.Errorf("code = %v, want %v", connectgo.CodeOf(got), tc.wantCode)
			}
		})
	}
}

func TestTranslateFromGRPCError_DomainTypes(t *testing.T) {
	cases := []struct {
		name    string
		grpcErr error
		check   func(error) bool
	}{
		{
			"CodeNotFound",
			connectgo.NewError(connectgo.CodeNotFound, stderrors.New("not found")),
			errors.IsNotFoundError,
		},
		{
			"CodeAlreadyExists",
			connectgo.NewError(connectgo.CodeAlreadyExists, stderrors.New("exists")),
			errors.IsResourceAlreadyCreatedError,
		},
		{
			"CodeInvalidArgument",
			connectgo.NewError(connectgo.CodeInvalidArgument, stderrors.New("bad")),
			errors.IsBadRequestError,
		},
		{
			"CodeUnauthenticated",
			connectgo.NewError(connectgo.CodeUnauthenticated, stderrors.New("unauth")),
			errors.IsUnauthorizedError,
		},
		{
			"CodeFailedPrecondition",
			connectgo.NewError(connectgo.CodeFailedPrecondition, stderrors.New("expired")),
			errors.IsExpiredResourceError,
		},
		{
			"CodeAborted",
			connectgo.NewError(connectgo.CodeAborted, stderrors.New("outdated")),
			errors.IsOutdatedResourceError,
		},
		{
			"CodeInternal",
			connectgo.NewError(connectgo.CodeInternal, stderrors.New("internal")),
			errors.IsInternalServerError,
		},
		{
			"CodeUnknown_fallsToInternal",
			connectgo.NewError(connectgo.CodeUnknown, stderrors.New("unknown")),
			errors.IsInternalServerError,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := pkggrpc.TranslateFromGRPCError(ctx, tc.grpcErr)
			if got == nil {
				t.Fatal("expected non-nil error")
			}
			if !tc.check(got) {
				t.Errorf("wrong domain error type for %s: %v", tc.name, got)
			}
		})
	}
}

func TestRoundTrip_DomainToGRPCToDomain(t *testing.T) {
	originals := []struct {
		name  string
		err   error
		check func(error) bool
	}{
		{"NotFound", errors.NewNotFoundError("k"), errors.IsNotFoundError},
		{"ResourceAlreadyCreated", errors.NewResourceAlreadyCreatedError("k"), errors.IsResourceAlreadyCreatedError},
		{"BadRequest", errors.NewBadRequestError("k"), errors.IsBadRequestError},
		{"Unauthorized", errors.NewUnauthorizedError("k"), errors.IsUnauthorizedError},
		{"ExpiredResource", errors.NewExpiredResourceError("k"), errors.IsExpiredResourceError},
		{"OutdatedResource", errors.NewOutdatedResourceError("k"), errors.IsOutdatedResourceError},
		{"InternalServer", errors.NewInternalServerError("k"), errors.IsInternalServerError},
	}
	for _, tc := range originals {
		t.Run(tc.name, func(t *testing.T) {
			grpcErr := pkggrpc.TranslateToGRPCError(ctx, tc.err)
			domainErr := pkggrpc.TranslateFromGRPCError(ctx, grpcErr)
			if !tc.check(domainErr) {
				t.Errorf("round-trip failed: got %v", domainErr)
			}
		})
	}
}
