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

var arrayType = reflect.TypeOf(make([]interface{}, 0))

func checkIsValidArray[T any](filter types.IFilter[T], element types.JsonElement) errors.FilterError {
	filterArray, ok := element.([]interface{})
	if !ok {
		return internaltypes.NewWrongFilterValueTypeError(filter, element, arrayType)
	}

	if len(filterArray) == 0 {
		return internalerrors.NewFilterError(errors.InvalidFilter, "array elements count should be more than 0, %v value %v has 0",
			reflect.TypeOf(element), filterArray)
	}

	return nil
}

func CheckIsValidObjectArray[T any](filter types.IFilter[T], element types.JsonElement, checker jsonPropertyChecker) errors.FilterError {
	err := checkIsValidArray(filter, element)
	if err != nil {
		return err
	}

	for _, filterObject := range element.([]interface{}) {
		err = CheckIsValidObject(filter, filterObject, checker)
		if err != nil {
			return err
		}
	}

	return nil
}

func CheckElementType[T any](filter types.IFilter[T], element types.JsonElement) errors.FilterError {
	var t T
	if _, ok := element.(T); !ok {
		return internaltypes.NewWrongFilterValueTypeError(filter, element, reflect.TypeOf(t))
	}
	return nil
}
