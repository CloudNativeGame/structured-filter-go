package builder

import "github.com/CloudNativeGame/structured-filter-go/internal/consts"

type FilterBuilderObject map[string]interface{}

func Eq[T comparable](value T) FilterBuilderObject {
	m := map[string]interface{}{
		consts.EqKey: value,
	}
	return m
}

func Ne[T comparable](value T) FilterBuilderObject {
	m := map[string]interface{}{
		consts.NeKey: value,
	}
	return m
}
