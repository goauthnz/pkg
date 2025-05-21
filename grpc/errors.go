package grpc

import (
	"context"

	connectgo "github.com/bufbuild/connect-go"
	"github.com/goauthnz/pkg/errors"
)

// TranslateFromGRPCError translates an error from a gRPC service to a errors.
// If no error is passed, it returns nil.
func TranslateFromGRPCError(_ context.Context, err error) error {
	// check if error is nil
	if err == nil {
		return nil
	}

	switch connectgo.CodeOf(err) {
	case connectgo.CodeNotFound:
		return errors.NewNotFoundError(err.Error())
	case connectgo.CodeAlreadyExists:
		return errors.NewResourceAlreadyCreatedError(err.Error())
	case connectgo.CodeInvalidArgument:
		return errors.NewBadRequestError(err.Error())
	case connectgo.CodeInternal:
		return errors.NewInternalServerError(err.Error())
	default:
		return errors.NewInternalServerError(err.Error())
	}
}

// TranslateToGRPCError translates an error from errors to a gRPC service.
// If no error is passed, it returns nil.
func TranslateToGRPCError(_ context.Context, err error) error {
	if err == nil {
		return nil
	}

	switch {
	case errors.IsNotFoundError(err):
		return connectgo.NewError(connectgo.CodeNotFound, err)
	case errors.IsResourceAlreadyCreatedError(err):
		return connectgo.NewError(connectgo.CodeAlreadyExists, err)
	case errors.IsBadRequestError(err):
		return connectgo.NewError(connectgo.CodeInvalidArgument, err)
	case errors.IsInternalServerError(err):
		return connectgo.NewError(connectgo.CodeInternal, err)
	default:
		return connectgo.NewError(connectgo.CodeInternal, err)
	}
}
