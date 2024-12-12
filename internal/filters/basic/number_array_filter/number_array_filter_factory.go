package number_array_filter

import (
	internaltypes "github.com/CloudNativeGame/structured-filter-go/internal/types"
	"github.com/CloudNativeGame/structured-filter-go/pkg/errors"
)

type NumberArrayFilterFactory struct {
	numberArrayFilters map[string]INumberArrayFilter
}

func NewNumberArrayFilterFactory(numberArrayFilters []INumberArrayFilter) NumberArrayFilterFactory {
	numberArrayFilterFactory := NumberArrayFilterFactory{
		numberArrayFilters: make(map[string]INumberArrayFilter, len(numberArrayFilters)),
	}
	for _, filter := range numberArrayFilters {
		numberArrayFilterFactory.numberArrayFilters[filter.GetKey()] = filter
	}
	return numberArrayFilterFactory
}

func (n NumberArrayFilterFactory) Get(key string) (INumberArrayFilter, errors.FilterError) {
	if filter, ok := n.numberArrayFilters[key]; ok {
		return filter, nil
	}

	return nil, internaltypes.NewKeyNotFoundError(key)
}
