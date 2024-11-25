package pkg

import (
	"github.com/CloudNativeGame/structured-filter-go/pkg/document"
	"github.com/CloudNativeGame/structured-filter-go/pkg/errors"
	"github.com/CloudNativeGame/structured-filter-go/pkg/factory"
	"github.com/CloudNativeGame/structured-filter-go/pkg/filters/scene_filter"
	"sync"
)

type FilterService[T any] struct {
	filterFactory   factory.FilterFactory[T]
	filterDocuments sync.Map
}

func NewFilterService[T any]() *FilterService[T] {
	return &FilterService[T]{
		filterFactory: *factory.NewFilterFactory[T](),
	}
}

type SceneFilterCreator[T any] func(filterFactory *factory.FilterFactory[T]) scene_filter.ISceneFilter[T]

func (f *FilterService[T]) WithSceneFilter(sceneFilterCreator SceneFilterCreator[T]) *FilterService[T] {
	f.filterFactory.WithSceneFilter(sceneFilterCreator(&f.filterFactory))
	return f
}

func (f *FilterService[T]) WithSceneFilters(sceneFilterCreators []SceneFilterCreator[T]) *FilterService[T] {
	for _, sceneCreator := range sceneFilterCreators {
		f.WithSceneFilter(sceneCreator)
	}
	return f
}

func (f *FilterService[T]) Match(rawFilter string, matchTarget T) errors.FilterError {
	if rawFilter == "" {
		return nil
	}

	filterDocument, err := f.loadOrAddFilterDocument(rawFilter)
	if err != nil {
		return err
	}

	for k, v := range filterDocument.Enumerate() {
		filter, err := f.filterFactory.Get(k)
		if err != nil {
			return err
		}
		err = filter.Match(v, matchTarget)
		if err != nil {
			return err
		}
	}

	return nil
}

func (f *FilterService[T]) loadOrAddFilterDocument(rawFilter string) (*document.FilterDocument[T], errors.FilterError) {
	if filterDocument, ok := f.filterDocuments.Load(rawFilter); ok {
		return filterDocument.(*document.FilterDocument[T]), nil
	}
	filterDocument, err := document.NewFilterDocument[T](rawFilter, f.filterFactory)
	if err != nil {
		return nil, err
	}
	f.filterDocuments.Store(rawFilter, filterDocument)
	return filterDocument, nil
}

func (f *FilterService[T]) FilterOut(rawFilter string, matchTargets []T) []T {
	filterResults := make([]T, 0)
	for _, filter := range matchTargets {
		err := f.Match(rawFilter, filter)
		if err == nil {
			filterResults = append(filterResults, filter)
		}
	}
	return filterResults
}
