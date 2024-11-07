package factory

import (
	"github.com/CloudNativeGame/structured-filter-go/internal/filters/basic"
	internalscene "github.com/CloudNativeGame/structured-filter-go/internal/filters/scene"
	"github.com/CloudNativeGame/structured-filter-go/pkg/errors"
	"github.com/CloudNativeGame/structured-filter-go/pkg/filters/scene"
	"github.com/CloudNativeGame/structured-filter-go/pkg/types"
)

type IFilterFactory[T any] interface {
	Get(key string) (types.IFilter[T], errors.FilterError)
}

type FilterFactory[T any] struct {
	BoolFilterFactory   basic.BoolFilterFactory
	NumberFilterFactory basic.NumberFilterFactory
	SceneFilterFactory  internalscene.SceneFilterFactory[T]
}

func NewFilterFactory[T any]() *FilterFactory[T] {
	filterFactory := &FilterFactory[T]{}
	filterFactory.BoolFilterFactory = basic.NewBoolFilterFactory([]basic.IBoolFilter{basic.BoolEqFilter{}})
	filterFactory.NumberFilterFactory = basic.NewNumberFilterFactory([]basic.INumberFilter{})
	filterFactory.SceneFilterFactory = internalscene.NewSceneFilterFactory([]scene.ISceneFilter[T]{})
	return filterFactory
}

func (f *FilterFactory[T]) Get(key string) (types.IFilter[T], errors.FilterError) {
	filter, err := f.SceneFilterFactory.Get(key)
	if err != nil {
		return nil, err
	}
	return filter, nil
}

func (f *FilterFactory[T]) WithSceneFilters(sceneFilters []scene.ISceneFilter[T]) *FilterFactory[T] {
	for _, filter := range sceneFilters {
		f.SceneFilterFactory.RegisterFilter(filter)
	}
	return f
}
