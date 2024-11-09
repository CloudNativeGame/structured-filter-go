package scene_filter

import (
	internalerrors "github.com/CloudNativeGame/structured-filter-go/internal/errors"
	"github.com/CloudNativeGame/structured-filter-go/pkg/errors"
	"github.com/CloudNativeGame/structured-filter-go/pkg/filters/scene_filter"
)

type SceneFilterFactory[T any] struct {
	sceneFilters map[string]scene_filter.ISceneFilter[T]
}

func NewSceneFilterFactory[T any](sceneFilters []scene_filter.ISceneFilter[T]) SceneFilterFactory[T] {
	sceneFilterFactory := SceneFilterFactory[T]{
		sceneFilters: make(map[string]scene_filter.ISceneFilter[T], len(sceneFilters)),
	}
	for _, filter := range sceneFilters {
		sceneFilterFactory.sceneFilters[filter.GetKey()] = filter
	}
	return sceneFilterFactory
}

func (s SceneFilterFactory[T]) Get(key string) (scene_filter.ISceneFilter[T], errors.FilterError) {
	if filter, ok := s.sceneFilters[key]; ok {
		return filter, nil
	}

	return nil, internalerrors.NewFilterError(errors.InvalidFilter, "filter key not found: %s", key)
}

func (s SceneFilterFactory[T]) RegisterFilter(filter scene_filter.ISceneFilter[T]) {
	s.sceneFilters[filter.GetKey()] = filter
}
