package string_array_filter

import (
	internaltypes "github.com/CloudNativeGame/structured-filter-go/internal/types"
	"github.com/CloudNativeGame/structured-filter-go/pkg/errors"
)

type StringArrayFilterFactory struct {
	stringArrayFilters map[string]IStringArrayFilter
}

func NewStringArrayFilterFactory(stringArrayFilters []IStringArrayFilter) StringArrayFilterFactory {
	stringArrayFilterFactory := StringArrayFilterFactory{
		stringArrayFilters: make(map[string]IStringArrayFilter, len(stringArrayFilters)),
	}
	for _, filter := range stringArrayFilters {
		stringArrayFilterFactory.stringArrayFilters[filter.GetKey()] = filter
	}
	return stringArrayFilterFactory
}

func (n StringArrayFilterFactory) Get(key string) (IStringArrayFilter, errors.FilterError) {
	if filter, ok := n.stringArrayFilters[key]; ok {
		return filter, nil
	}

	return nil, internaltypes.NewKeyNotFoundError(key)
}
