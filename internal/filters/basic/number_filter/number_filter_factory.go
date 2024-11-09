package number_filter

import (
	internaltypes "github.com/CloudNativeGame/structured-filter-go/internal/types"
	"github.com/CloudNativeGame/structured-filter-go/pkg/errors"
)

type NumberFilterFactory struct {
	numberFilters map[string]INumberFilter
}

func NewNumberFilterFactory(numberFilters []INumberFilter) NumberFilterFactory {
	numberFilterFactory := NumberFilterFactory{
		numberFilters: make(map[string]INumberFilter, len(numberFilters)),
	}
	for _, filter := range numberFilters {
		numberFilterFactory.numberFilters[filter.GetKey()] = filter
	}
	return numberFilterFactory
}

func (n NumberFilterFactory) Get(key string) (INumberFilter, errors.FilterError) {
	if filter, ok := n.numberFilters[key]; ok {
		return filter, nil
	}

	return nil, internaltypes.NewKeyNotFoundError(key)
}
