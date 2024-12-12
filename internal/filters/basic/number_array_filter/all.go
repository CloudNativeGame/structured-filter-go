package number_array_filter

import (
	"github.com/CloudNativeGame/structured-filter-go/internal/filters/basic/common/array_filter"
)

func NewNumberArrayAllFilter() INumberArrayFilter {
	return array_filter.NewAllFilter[[]float64, float64]()
}
