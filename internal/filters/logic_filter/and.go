package logic_filter

import (
	"github.com/CloudNativeGame/structured-filter-go/internal/consts"
	"github.com/CloudNativeGame/structured-filter-go/internal/filters/scene_filter"
	internaltypes "github.com/CloudNativeGame/structured-filter-go/internal/types"
	"github.com/CloudNativeGame/structured-filter-go/pkg/checkers"
	"github.com/CloudNativeGame/structured-filter-go/pkg/errors"
	"github.com/CloudNativeGame/structured-filter-go/pkg/types"
)

type AndFilter[T any] struct {
	sceneFilterFactory scene_filter.SceneFilterFactory[T]
}

func NewAndFilter[T any](sceneFilterFactory scene_filter.SceneFilterFactory[T]) *AndFilter[T] {
	return &AndFilter[T]{
		sceneFilterFactory: sceneFilterFactory,
	}
}

func (a AndFilter[T]) GetKey() string {
	return consts.AndKey
}

func (a AndFilter[T]) Valid(element types.JsonElement) errors.FilterError {
	return checkers.CheckIsValidObjectArray(a, element, func(propertyKey string, propertyValue interface{}) errors.FilterError {
		filter, err := a.sceneFilterFactory.Get(propertyKey)
		if err != nil {
			return err
		}
		return filter.Valid(propertyValue)
	})
}

func (a AndFilter[T]) Match(element types.JsonElement, matchTarget T) errors.FilterError {
	for _, filterObject := range element.([]interface{}) {
		for propertyKey, propertyValue := range filterObject.(map[string]interface{}) {
			filter, _ := a.sceneFilterFactory.Get(propertyKey)
			err := filter.Match(propertyValue, matchTarget)
			if err != nil {
				return internaltypes.NewNotMatchError(a, matchTarget, element, err)
			}
		}
	}

	return nil
}
