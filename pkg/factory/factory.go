package factory

import (
	"github.com/CloudNativeGame/structured-filter-go/internal/filters/basic"
	"github.com/CloudNativeGame/structured-filter-go/internal/filters/logic"
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
	StringFilterFactory basic.StringFilterFactory
	SceneFilterFactory  internalscene.SceneFilterFactory[T]
	LogicFilterFactory  logic.LogicFilterFactory[T]
}

func NewFilterFactory[T any]() *FilterFactory[T] {
	filterFactory := &FilterFactory[T]{}
	filterFactory.BoolFilterFactory = basic.NewBoolFilterFactory([]basic.IBoolFilter{
		basic.NewBoolEqFilter(),
	})
	filterFactory.NumberFilterFactory = basic.NewNumberFilterFactory([]basic.INumberFilter{
		basic.NewNumberEqFilter(),
	})
	filterFactory.StringFilterFactory = basic.NewStringFilterFactory([]basic.IStringFilter{
		basic.NewStringEqFilter(),
	})
	filterFactory.SceneFilterFactory = internalscene.NewSceneFilterFactory([]scene.ISceneFilter[T]{})
	filterFactory.LogicFilterFactory = logic.NewLogicFilterFactory([]logic.ILogicFilter[T]{
		logic.NewAndFilter(filterFactory.SceneFilterFactory),
		logic.NewOrFilter(filterFactory.SceneFilterFactory),
	})
	return filterFactory
}

func (f *FilterFactory[T]) Get(key string) (types.IFilter[T], errors.FilterError) {
	filter, err := f.LogicFilterFactory.Get(key)
	if err != nil {
		filter, err = f.SceneFilterFactory.Get(key)
		if err != nil {
			return nil, err
		}
	}
	return filter, nil
}

func (f *FilterFactory[T]) WithSceneFilters(sceneFilters []scene.ISceneFilter[T]) *FilterFactory[T] {
	for _, filter := range sceneFilters {
		f.SceneFilterFactory.RegisterFilter(filter)
	}
	return f
}
