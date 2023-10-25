package xcode

import (
	"context"
	"github.com/pkg/errors"
)

type Status struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
	Details any    `json:"details,omitempty"`
}

func CodeFromError(err error) XCode {
	err = errors.Cause(err)
	var code XCode
	if errors.As(err, &code) {
		return code
	}

	switch {
	case errors.Is(err, context.Canceled):
		return Canceled
	case errors.Is(err, context.DeadlineExceeded):
		return Deadline
	}

	return ServerErr
}
