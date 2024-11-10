package logic_filter

import (
	"github.com/CloudNativeGame/structured-filter-go/internal/consts"
	"github.com/CloudNativeGame/structured-filter-go/internal/filters/scene_filter"
	internaltypes "github.com/CloudNativeGame/structured-filter-go/internal/types"
	"github.com/CloudNativeGame/structured-filter-go/pkg/checkers"
	"github.com/CloudNativeGame/structured-filter-go/pkg/errors"
	"github.com/CloudNativeGame/structured-filter-go/pkg/types"
)

type OrFilter[T any] struct {
	sceneFilterFactory scene_filter.SceneFilterFactory[T]
}

func NewOrFilter[T any](sceneFilterFactory scene_filter.SceneFilterFactory[T]) *OrFilter[T] {
	return &OrFilter[T]{
		sceneFilterFactory: sceneFilterFactory,
	}
}

func (o OrFilter[T]) GetKey() string {
	return consts.OrKey
}

func (o OrFilter[T]) Valid(element types.JsonElement) errors.FilterError {
	return checkers.CheckIsValidObjectArray(o, element, func(propertyKey string, propertyValue interface{}) errors.FilterError {
		filter, err := o.sceneFilterFactory.Get(propertyKey)
		if err != nil {
			return err
		}
		return filter.Valid(propertyValue)
	})
}

func (o OrFilter[T]) Match(element types.JsonElement, matchTarget T) errors.FilterError {
	var lastErr errors.FilterError
	for _, filterObject := range element.([]interface{}) {
		for propertyKey, propertyValue := range filterObject.(map[string]interface{}) {
			filter, _ := o.sceneFilterFactory.Get(propertyKey)
			err := filter.Match(propertyValue, matchTarget)
			if err == nil {
				return nil
			}
			lastErr = err
		}
	}

	return internaltypes.NewNotMatchError(o, matchTarget, element, lastErr)
}
