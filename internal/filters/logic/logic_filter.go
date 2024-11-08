package logic

import (
	"github.com/CloudNativeGame/structured-filter-go/internal/filters/scene"
	internaltypes "github.com/CloudNativeGame/structured-filter-go/internal/types"
	"github.com/CloudNativeGame/structured-filter-go/pkg/checkers"
	"github.com/CloudNativeGame/structured-filter-go/pkg/errors"
	"github.com/CloudNativeGame/structured-filter-go/pkg/types"
)

type ILogicFilter[T any] interface {
	types.IFilter[T]
}

type AndFilter[T any] struct {
	sceneFilterFactory scene.SceneFilterFactory[T]
}

func NewAndFilter[T any](sceneFilterFactory scene.SceneFilterFactory[T]) *AndFilter[T] {
	return &AndFilter[T]{
		sceneFilterFactory: sceneFilterFactory,
	}
}

func (a AndFilter[T]) GetKey() string {
	return "$and"
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

type OrFilter[T any] struct {
	sceneFilterFactory scene.SceneFilterFactory[T]
}

func NewOrFilter[T any](sceneFilterFactory scene.SceneFilterFactory[T]) *OrFilter[T] {
	return &OrFilter[T]{
		sceneFilterFactory: sceneFilterFactory,
	}
}

func (o OrFilter[T]) GetKey() string {
	return "$or"
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
