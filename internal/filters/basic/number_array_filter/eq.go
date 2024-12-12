package number_array_filter

import (
	"github.com/CloudNativeGame/structured-filter-go/internal/filters/basic/common/array_filter"
)

func NewNumberArrayEqFilter() INumberArrayFilter {
	return array_filter.NewEqFilter[[]float64, float64]()
}
