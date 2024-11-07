package errors

import (
	"fmt"
	"github.com/CloudNativeGame/structured-filter-go/pkg/errors"
)

type filterErrorImpl struct {
	errorType errors.FilterErrorType
	msg       string
}

func (e filterErrorImpl) Error() string {
	return e.msg
}

func (e filterErrorImpl) Type() errors.FilterErrorType {
	return e.errorType
}

func NewFilterError(errorType errors.FilterErrorType, msg string, args ...interface{}) errors.FilterError {
	return filterErrorImpl{
		errorType: errorType,
		msg:       fmt.Sprintf(msg, args...),
	}
}

func ToFilterError(err error, errorType errors.FilterErrorType) errors.FilterError {
	if err == nil {
		return nil
	}
	return filterErrorImpl{
		errorType: errorType,
		msg:       err.Error(),
	}
}
