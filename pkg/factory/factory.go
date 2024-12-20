package factory

import (
	"github.com/CloudNativeGame/structured-filter-go/internal/filters/basic/bool_filter"
	"github.com/CloudNativeGame/structured-filter-go/internal/filters/basic/number_array_filter"
	"github.com/CloudNativeGame/structured-filter-go/internal/filters/basic/number_filter"
	"github.com/CloudNativeGame/structured-filter-go/internal/filters/basic/string_array_filter"
	"github.com/CloudNativeGame/structured-filter-go/internal/filters/basic/string_filter"
	"github.com/CloudNativeGame/structured-filter-go/internal/filters/logic_filter"
	internalscenefilter "github.com/CloudNativeGame/structured-filter-go/internal/filters/scene_filter"
	"github.com/CloudNativeGame/structured-filter-go/pkg/errors"
	"github.com/CloudNativeGame/structured-filter-go/pkg/filters/scene_filter"
	"github.com/CloudNativeGame/structured-filter-go/pkg/types"
)

type IFilterFactory[T any] interface {
	Get(key string) (types.IFilter[T], errors.FilterError)
}

type FilterFactory[T any] struct {
	BoolFilterFactory        bool_filter.BoolFilterFactory
	NumberFilterFactory      number_filter.NumberFilterFactory
	NumberArrayFilterFactory number_array_filter.NumberArrayFilterFactory
	StringFilterFactory      string_filter.StringFilterFactory
	StringArrayFilterFactory string_array_filter.StringArrayFilterFactory
	SceneFilterFactory       internalscenefilter.SceneFilterFactory[T]
	LogicFilterFactory       logic_filter.LogicFilterFactory[T]
}

func NewFilterFactory[T any]() *FilterFactory[T] {
	filterFactory := &FilterFactory[T]{}
	filterFactory.BoolFilterFactory = bool_filter.NewBoolFilterFactory([]bool_filter.IBoolFilter{
		bool_filter.NewBoolEqFilter(),
		bool_filter.NewBoolNeFilter(),
	})
	filterFactory.NumberFilterFactory = number_filter.NewNumberFilterFactory([]number_filter.INumberFilter{
		number_filter.NewNumberEqFilter(),
		number_filter.NewNumberNeFilter(),
		number_filter.NewNumberRangeFilter(),
		number_filter.NewNumberInFilter(),
		number_filter.NewNumberGtFilter(),
		number_filter.NewNumberLtFilter(),
		number_filter.NewNumberGeFilter(),
		number_filter.NewNumberLeFilter(),
	})
	filterFactory.NumberArrayFilterFactory = number_array_filter.NewNumberArrayFilterFactory([]number_array_filter.INumberArrayFilter{
		number_array_filter.NewNumberArrayAllFilter(),
		number_array_filter.NewNumberArrayEqFilter(),
	})
	filterFactory.StringFilterFactory = string_filter.NewStringFilterFactory([]string_filter.IStringFilter{
		string_filter.NewStringEqFilter(),
		string_filter.NewStringNeFilter(),
		string_filter.NewRegexFilter(),
		string_filter.NewStringInFilter(),
		string_filter.NewStringRangeFilter(),
	})
	filterFactory.StringArrayFilterFactory = string_array_filter.NewStringArrayFilterFactory([]string_array_filter.IStringArrayFilter{
		string_array_filter.NewStringArrayAllFilter(),
		string_array_filter.NewStringArrayEqFilter(),
	})
	filterFactory.SceneFilterFactory = internalscenefilter.NewSceneFilterFactory([]scene_filter.ISceneFilter[T]{})
	filterFactory.LogicFilterFactory = logic_filter.NewLogicFilterFactory([]logic_filter.ILogicFilter[T]{
		logic_filter.NewAndFilter(filterFactory.SceneFilterFactory),
		logic_filter.NewOrFilter(filterFactory.SceneFilterFactory),
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

func (f *FilterFactory[T]) WithSceneFilter(sceneFilter scene_filter.ISceneFilter[T]) *FilterFactory[T] {
	f.SceneFilterFactory.RegisterFilter(sceneFilter)

	return f
}

func (f *FilterFactory[T]) WithSceneFilters(sceneFilters []scene_filter.ISceneFilter[T]) *FilterFactory[T] {
	for _, filter := range sceneFilters {
		f.WithSceneFilter(filter)
	}
	return f
}
