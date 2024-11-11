package utils

import (
	"fmt"
	"reflect"
)

func ToFloat64(num interface{}) float64 {
	numType := reflect.TypeOf(num)
	numValue := reflect.ValueOf(num)

	switch numType.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float64(numValue.Int())
	case reflect.Float32, reflect.Float64:
		return numValue.Float()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return float64(numValue.Uint())
	default:
		panic(fmt.Errorf("unsupported number type %v for %v", numType, numValue))
	}
}
