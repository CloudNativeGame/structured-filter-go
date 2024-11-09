package string_filter

import (
	internaltypes "github.com/CloudNativeGame/structured-filter-go/internal/types"
	"github.com/CloudNativeGame/structured-filter-go/pkg/errors"
)

type StringFilterFactory struct {
	stringFilters map[string]IStringFilter
}

func NewStringFilterFactory(stringFilters []IStringFilter) StringFilterFactory {
	stringFilterFactory := StringFilterFactory{
		stringFilters: make(map[string]IStringFilter, len(stringFilters)),
	}
	for _, filter := range stringFilters {
		stringFilterFactory.stringFilters[filter.GetKey()] = filter
	}
	return stringFilterFactory
}

func (n StringFilterFactory) Get(key string) (IStringFilter, errors.FilterError) {
	if filter, ok := n.stringFilters[key]; ok {
		return filter, nil
	}

	return nil, internaltypes.NewKeyNotFoundError(key)
}
