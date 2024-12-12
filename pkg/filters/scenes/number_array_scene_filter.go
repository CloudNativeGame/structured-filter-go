package scenes

import (
	"github.com/CloudNativeGame/structured-filter-go/pkg/checkers"
	"github.com/CloudNativeGame/structured-filter-go/pkg/errors"
	"github.com/CloudNativeGame/structured-filter-go/pkg/factory"
	"github.com/CloudNativeGame/structured-filter-go/pkg/types"
)

type numberArrayValueGetter[T any] func(matchTarget T) []float64

type NumberArraySceneFilter[T any] struct {
	filterFactory *factory.FilterFactory[T]
	key           string
	valueGetter   numberArrayValueGetter[T]
}

func (n *NumberArraySceneFilter[T]) GetKey() string {
	return n.key
}

func (n *NumberArraySceneFilter[T]) Valid(element types.JsonElement) errors.FilterError {
	return checkers.CheckIsValidObject(n, element, func(propertyKey string, propertyValue interface{}) errors.FilterError {
		filter, err := n.filterFactory.NumberArrayFilterFactory.Get(propertyKey)
		if err != nil {
			return err
		}
		return filter.Valid(propertyValue)
	})
}

func (n *NumberArraySceneFilter[T]) Match(element types.JsonElement, matchTarget T) errors.FilterError {
	for k, v := range element.(map[string]interface{}) {
		filter, err := n.filterFactory.NumberArrayFilterFactory.Get(k)
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

func NewNumberArraySceneFilter[T any](key string, valueGetter numberArrayValueGetter[T], filterFactory *factory.FilterFactory[T]) *NumberArraySceneFilter[T] {
	return &NumberArraySceneFilter[T]{
		filterFactory: filterFactory,
		key:           key,
		valueGetter:   valueGetter,
	}
}
