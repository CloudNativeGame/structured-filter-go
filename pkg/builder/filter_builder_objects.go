package builder

import (
	"github.com/CloudNativeGame/structured-filter-go/internal/consts"
	"github.com/CloudNativeGame/structured-filter-go/internal/utils"
)

type FilterBuilderObject map[string]interface{}

func Eq[T comparable](value T) FilterBuilderObject {
	return filterBuilderObject(consts.EqKey, value)
}

func Ne[T comparable](value T) FilterBuilderObject {
	return filterBuilderObject(consts.NeKey, value)
}

func NumberRange[T comparable](value []T) FilterBuilderObject {
	float64Arr := make([]float64, 0, len(value))
	for _, v := range value {
		float64Arr = append(float64Arr, utils.ToFloat64(v))
	}
	return filterBuilderObject(consts.RangeKey, float64Arr)
}

func Regex(value string) FilterBuilderObject {
	return filterBuilderObject(consts.RegexKey, value)
}

func NumberIn[T comparable](value []T) FilterBuilderObject {
	float64Arr := make([]float64, 0, len(value))
	for _, v := range value {
		float64Arr = append(float64Arr, utils.ToFloat64(v))
	}
	return filterBuilderObject(consts.InKey, float64Arr)
}

func StringIn(value []string) FilterBuilderObject {
	return filterBuilderObject(consts.InKey, value)
}

func filterBuilderObject[T any](key string, value T) FilterBuilderObject {
	m := map[string]interface{}{
		key: value,
	}
	return m
}
