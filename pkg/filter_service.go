package pkg

import (
	"github.com/CloudNativeGame/structured-filter-go/pkg/document"
	"github.com/CloudNativeGame/structured-filter-go/pkg/errors"
	"github.com/CloudNativeGame/structured-filter-go/pkg/factory"
)

type FilterService[T any] struct {
	filterFactory factory.FilterFactory[T]
}

func NewFilterService[T any](filterFactory *factory.FilterFactory[T]) *FilterService[T] {
	return &FilterService[T]{
		filterFactory: *filterFactory,
	}
}

func (f FilterService[T]) MatchFilter(rawFilter string, matchTarget T) errors.FilterError {
	if rawFilter == "" {
		return nil
	}

	filterDocument, err := document.NewFilterDocument[T](rawFilter, f.filterFactory)
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
