package errors

import (
	"fmt"

	"github.com/pkg/errors"
)

func Wrap(err error, message string) error {
	return errors.Wrap(err, message)
}

func New(message string, value ...interface{}) error {
	return errors.New(fmt.Sprintf(message, value...))
}

func Error(e error) *HTTPError {
	switch e.(type) {
	case *HTTPError:
		return e.(*HTTPError)
	}
	return InternalErrorHappened.Error(e)
}
