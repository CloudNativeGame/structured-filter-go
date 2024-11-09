package scenes

import (
	"github.com/CloudNativeGame/structured-filter-go/pkg/checkers"
	"github.com/CloudNativeGame/structured-filter-go/pkg/types"

	"github.com/CloudNativeGame/structured-filter-go/pkg/errors"
	"github.com/CloudNativeGame/structured-filter-go/pkg/factory"
)

type boolValueGetter[T any] func(matchTarget T) bool

type BoolSceneFilter[T any] struct {
	filterFactory *factory.FilterFactory[T]
	key           string
	valueGetter   boolValueGetter[T]
}

func (b *BoolSceneFilter[T]) GetKey() string {
	return b.key
}

func (b *BoolSceneFilter[T]) Valid(element types.JsonElement) errors.FilterError {
	return checkers.CheckIsValidObject(b, element, func(propertyKey string, propertyValue interface{}) errors.FilterError {
		filter, err := b.filterFactory.BoolFilterFactory.Get(propertyKey)
		if err != nil {
			return err
		}
		return filter.Valid(propertyValue)
	})
}

func (b *BoolSceneFilter[T]) Match(element types.JsonElement, matchTarget T) errors.FilterError {
	for k, v := range element.(map[string]interface{}) {
		filter, err := b.filterFactory.BoolFilterFactory.Get(k)
		if err != nil {
			return err
		}
		err = filter.Match(v, b.valueGetter(matchTarget))
		if err != nil {
			return err
		}
	}
	return nil
}

func NewBoolSceneFilter[T any](key string, valueGetter boolValueGetter[T], filterFactory *factory.FilterFactory[T]) *BoolSceneFilter[T] {
	return &BoolSceneFilter[T]{
		filterFactory: filterFactory,
		key:           key,
		valueGetter:   valueGetter,
	}
}
