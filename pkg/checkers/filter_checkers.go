package checkers

import (
	internalerrors "github.com/CloudNativeGame/structured-filter-go/internal/errors"
	internaltypes "github.com/CloudNativeGame/structured-filter-go/internal/types"
	"github.com/CloudNativeGame/structured-filter-go/internal/utils"
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

func checkIsValidArray[T any](filter types.IFilter[T], element types.JsonElement, elementNumber *int) ([]interface{}, errors.FilterError) {
	filterArray, ok := element.([]interface{})
	if !ok {
		return nil, internaltypes.NewWrongFilterValueTypeError(filter, element, arrayType)
	}

	if len(filterArray) == 0 {
		return nil, internalerrors.NewFilterError(errors.InvalidFilter, "array elements count should be more than 0, %v value %v has 0",
			reflect.TypeOf(element), filterArray)
	}

	if elementNumber != nil {
		if len(filterArray) != *elementNumber {
			return nil, internalerrors.NewFilterError(errors.InvalidFilter, "array elements count should be %d, %v value %v has %d",
				*elementNumber, reflect.TypeOf(element), filterArray, len(filterArray))
		}
	}

	return filterArray, nil
}

func CheckIsValidArray[T any](filter types.IFilter[T], element types.JsonElement, elementNumber *int, checkType bool) errors.FilterError {
	filterArray, err := checkIsValidArray(filter, element, elementNumber)
	if err != nil {
		return err
	}

	if checkType {
		for _, val := range filterArray {
			err := CheckElementType(filter, val)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func CheckIsValidArrayForArrayFilter[TFilter, TFilterElement any](filter types.IArrayFilter[TFilter, TFilterElement], element types.JsonElement, elementNumber *int, checkType bool) errors.FilterError {
	filterArray, err := checkIsValidArray(filter, element, elementNumber)
	if err != nil {
		return err
	}

	if checkType {
		for _, val := range filterArray {
			err := checkElementTypeForArrayFilter[TFilter, TFilterElement](filter, val)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func CheckIsValidObjectArray[T any](filter types.IFilter[T], element types.JsonElement, checker jsonPropertyChecker) errors.FilterError {
	err := CheckIsValidArray(filter, element, nil, false)
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

var rangeElementNumber = 2

func CheckIsValidNumberRange(filter types.IFilter[float64], element types.JsonElement) errors.FilterError {
	err := checkIsValidRange(filter, element)
	if err != nil {
		return err
	}

	filterArray, ok := element.([]interface{})
	if !ok {
		return internaltypes.NewWrongFilterValueTypeError(filter, element, arrayType)
	}

	if utils.NumberToFloat64(filterArray[1])-utils.NumberToFloat64(filterArray[0]) < 0 {
		return internalerrors.NewFilterError(errors.InvalidFilter,
			"the second element of the range %f is not >= the first element %f",
			filterArray[1], filterArray[0])
	}

	return nil
}

func checkIsValidRange[T comparable](filter types.IFilter[T], element types.JsonElement) errors.FilterError {
	err := CheckIsValidArray(filter, element, &rangeElementNumber, true)
	if err != nil {
		return err
	}

	return nil
}

func CheckIsValidStringRange(filter types.IFilter[string], element types.JsonElement) errors.FilterError {
	err := checkIsValidRange(filter, element)
	if err != nil {
		return err
	}

	filterArray, ok := element.([]interface{})
	if !ok {
		return internaltypes.NewWrongFilterValueTypeError(filter, element, arrayType)
	}

	firstElem, ok := filterArray[0].(string)
	if !ok {
		return internaltypes.NewWrongFilterValueTypeError(filter, element, arrayType)
	}

	secondElem, ok := filterArray[1].(string)
	if !ok {
		return internaltypes.NewWrongFilterValueTypeError(filter, element, arrayType)
	}

	if secondElem < firstElem {
		return internalerrors.NewFilterError(errors.InvalidFilter,
			"the second element of the range %s is not >= the first element %s",
			filterArray[1], filterArray[0])
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

func checkElementTypeForArrayFilter[TFilter, TFilterElement any](filter types.IArrayFilter[TFilter, TFilterElement], element types.JsonElement) errors.FilterError {
	var t TFilterElement
	if _, ok := element.(TFilterElement); !ok {
		return internaltypes.NewWrongFilterValueTypeError(filter, element, reflect.TypeOf(t))
	}
	return nil
}
