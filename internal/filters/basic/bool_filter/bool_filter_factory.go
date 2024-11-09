package bool_filter

import (
	internaltypes "github.com/CloudNativeGame/structured-filter-go/internal/types"
	"github.com/CloudNativeGame/structured-filter-go/pkg/errors"
)

type BoolFilterFactory struct {
	boolFilters map[string]IBoolFilter
}

func NewBoolFilterFactory(boolFilters []IBoolFilter) BoolFilterFactory {
	boolFilterFactory := BoolFilterFactory{
		boolFilters: make(map[string]IBoolFilter, len(boolFilters)),
	}
	for _, filter := range boolFilters {
		boolFilterFactory.boolFilters[filter.GetKey()] = filter
	}
	return boolFilterFactory
}

func (b BoolFilterFactory) Get(key string) (IBoolFilter, errors.FilterError) {
	if filter, ok := b.boolFilters[key]; ok {
		return filter, nil
	}

	return nil, internaltypes.NewKeyNotFoundError(key)
}
