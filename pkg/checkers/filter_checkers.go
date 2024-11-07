package checkers

import (
	internalerrors "github.com/CloudNativeGame/structured-filter-go/internal/errors"
	internaltypes "github.com/CloudNativeGame/structured-filter-go/internal/types"
	"github.com/CloudNativeGame/structured-filter-go/pkg/errors"
	"github.com/CloudNativeGame/structured-filter-go/pkg/types"
	"reflect"
)

type jsonPropertyChecker func(propertyKey string, propertyValue interface{}) errors.FilterError

var objectType = reflect.TypeOf(make(map[string]interface{}))

func CheckIsValidObject[T any](filter types.IFilter[T], element types.JsonElement, checker jsonPropertyChecker) errors.FilterError {
	filterObject, ok := element.(map[string]interface{})
	if !ok {

		return internaltypes.NewWrongFilterValueTypeError(filter, element, objectType)
	}

	// limit kv count to 1
	if len(filterObject) != 1 {
		return internalerrors.NewFilterError(errors.InvalidFilter, "object kv count should be 1, %v value %v has %d",
			reflect.TypeOf(element), filterObject, len(filterObject))
	}

	for k, v := range filterObject {
		if err := checker(k, v); err != nil {
			return err
		}
	}

	return nil
}

var boolType = reflect.TypeOf(true)

func CheckIsValidBool[T any](filter types.IFilter[T], element types.JsonElement) errors.FilterError {
	if _, ok := element.(bool); !ok {
		return internaltypes.NewWrongFilterValueTypeError(filter, element, boolType)
	}
	return nil
}
