package scenes

import (
	"github.com/CloudNativeGame/structured-filter-go/pkg/checkers"
	"github.com/CloudNativeGame/structured-filter-go/pkg/errors"
	"github.com/CloudNativeGame/structured-filter-go/pkg/factory"
	"github.com/CloudNativeGame/structured-filter-go/pkg/types"
)

type stringArrayValueGetter[T any] func(matchTarget T) []string

type StringArraySceneFilter[T any] struct {
	filterFactory *factory.FilterFactory[T]
	key           string
	valueGetter   stringArrayValueGetter[T]
}

func (n *StringArraySceneFilter[T]) GetKey() string {
	return n.key
}

func (n *StringArraySceneFilter[T]) Valid(element types.JsonElement) errors.FilterError {
	return checkers.CheckIsValidObject(n, element, func(propertyKey string, propertyValue interface{}) errors.FilterError {
		filter, err := n.filterFactory.StringArrayFilterFactory.Get(propertyKey)
		if err != nil {
			return err
		}
		return filter.Valid(propertyValue)
	})
}

func (n *StringArraySceneFilter[T]) Match(element types.JsonElement, matchTarget T) errors.FilterError {
	for k, v := range element.(map[string]interface{}) {
		filter, err := n.filterFactory.StringArrayFilterFactory.Get(k)
		if err != nil {
			return err
		}
		err = filter.Match(v, n.valueGetter(matchTarget))
		if err != nil {
			return err
		}
	}
	return nil
}

func NewStringArraySceneFilter[T any](key string, valueGetter stringArrayValueGetter[T], filterFactory *factory.FilterFactory[T]) *StringArraySceneFilter[T] {
	return &StringArraySceneFilter[T]{
		filterFactory: filterFactory,
		key:           key,
		valueGetter:   valueGetter,
	}
}
