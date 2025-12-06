package render

import (
	"context"
	"errors"
	"fmt"

	"github.com/alexeevayaan/portfolio-react-golang-todo/api/pkg/logger"
)

func Error(ctx context.Context, err error, message string) error {
	ctxError, ok := ctx.Value(logger.ContextErrKey{}).(*error)
	if !ok {
		*ctxError = fmt.Errorf("%s: %w", message, err)
	}

	err = unpack(err)

	return fmt.Errorf("%s: %w", message, err)
}

func unpack(err error) error {
	for {
		e := errors.Unwrap(err)
		if e == nil {
			break
		}
		err = e
	}

	return err
}
