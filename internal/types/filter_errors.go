package types

import (
	internalerrors "github.com/CloudNativeGame/structured-filter-go/internal/errors"
	"github.com/CloudNativeGame/structured-filter-go/pkg/errors"
	"github.com/CloudNativeGame/structured-filter-go/pkg/types"
	"reflect"
)

func NewWrongFilterValueTypeError[T any](filter types.IFilter[T], element types.JsonElement, expectedType reflect.Type) errors.FilterError {
	return internalerrors.NewFilterError(errors.InvalidFilter, "%v value type is %v, not expected %v", reflect.TypeOf(filter), reflect.TypeOf(element), expectedType)
}

func NewKeyNotFoundError(key string) errors.FilterError {
	return internalerrors.NewFilterError(errors.InvalidFilter, "filter key not found: %s", key)
}

func NewNotMatchError[T any](filter types.IFilter[T], value T, element types.JsonElement) errors.FilterError {
	return internalerrors.NewFilterError(errors.NotMatch, "%v value %v does not match filter {%s: %v}", reflect.TypeOf(filter), value, filter.GetKey(), element)
}
