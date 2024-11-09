package scenes

import (
	"github.com/CloudNativeGame/structured-filter-go/pkg/checkers"
	"github.com/CloudNativeGame/structured-filter-go/pkg/types"

	"github.com/CloudNativeGame/structured-filter-go/pkg/errors"
	"github.com/CloudNativeGame/structured-filter-go/pkg/factory"
)

type stringValueGetter[T any] func(matchTarget T) string

type StringSceneFilter[T any] struct {
	filterFactory *factory.FilterFactory[T]
	key           string
	valueGetter   stringValueGetter[T]
}

func (s *StringSceneFilter[T]) GetKey() string {
	return s.key
}

func (s *StringSceneFilter[T]) Valid(element types.JsonElement) errors.FilterError {
	return checkers.CheckIsValidObject(s, element, func(propertyKey string, propertyValue interface{}) errors.FilterError {
		filter, err := s.filterFactory.StringFilterFactory.Get(propertyKey)
		if err != nil {
			return err
		}
		return filter.Valid(propertyValue)
	})
}

func (s *StringSceneFilter[T]) Match(element types.JsonElement, matchTarget T) errors.FilterError {
	for k, v := range element.(map[string]interface{}) {
		filter, err := s.filterFactory.StringFilterFactory.Get(k)
		if err != nil {
			return err
		}
		err = filter.Match(v, s.valueGetter(matchTarget))
		if err != nil {
			return err
		}
	}
	return nil
}

func NewStringSceneFilter[T any](key string, valueGetter stringValueGetter[T], filterFactory *factory.FilterFactory[T]) *StringSceneFilter[T] {
	return &StringSceneFilter[T]{
		filterFactory: filterFactory,
		key:           key,
		valueGetter:   valueGetter,
	}
}
