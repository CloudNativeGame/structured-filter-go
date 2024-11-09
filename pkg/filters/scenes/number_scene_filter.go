package scenes

import (
	"github.com/CloudNativeGame/structured-filter-go/pkg/checkers"
	"github.com/CloudNativeGame/structured-filter-go/pkg/types"

	"github.com/CloudNativeGame/structured-filter-go/pkg/errors"
	"github.com/CloudNativeGame/structured-filter-go/pkg/factory"
)

type numberValueGetter[T any] func(matchTarget T) float64

type NumberSceneFilter[T any] struct {
	filterFactory *factory.FilterFactory[T]
	key           string
	valueGetter   numberValueGetter[T]
}

func (n *NumberSceneFilter[T]) GetKey() string {
	return n.key
}

func (n *NumberSceneFilter[T]) Valid(element types.JsonElement) errors.FilterError {
	return checkers.CheckIsValidObject(n, element, func(propertyKey string, propertyValue interface{}) errors.FilterError {
		filter, err := n.filterFactory.NumberFilterFactory.Get(propertyKey)
		if err != nil {
			return err
		}
		return filter.Valid(propertyValue)
	})
}

func (n *NumberSceneFilter[T]) Match(element types.JsonElement, matchTarget T) errors.FilterError {
	for k, v := range element.(map[string]interface{}) {
		filter, err := n.filterFactory.NumberFilterFactory.Get(k)
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

func NewNumberSceneFilter[T any](key string, valueGetter numberValueGetter[T], filterFactory *factory.FilterFactory[T]) *NumberSceneFilter[T] {
	return &NumberSceneFilter[T]{
		filterFactory: filterFactory,
		key:           key,
		valueGetter:   valueGetter,
	}
}
