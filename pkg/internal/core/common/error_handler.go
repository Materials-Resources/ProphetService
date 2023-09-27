package common

import (
	"errors"

	"github.com/materials-resources/s_prophet/pkg/internal/core/port/repository"
	"google.golang.org/genproto/googleapis/rpc/code"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GetRpcError(err error) error {
	if errors.Is(
		repository.NotFound,
		err,
	) {
		return status.Error(
			codes.Code(code.Code_NOT_FOUND),
			"could not find",
		)
	}
	return status.Error(
		codes.Unknown,
		"a unknown error has occurred",
	)
}
