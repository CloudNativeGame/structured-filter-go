package number_filter

import (
	"github.com/CloudNativeGame/structured-filter-go/internal/filters/basic/common/value_filter"
)

func NewNumberEqFilter() INumberFilter {
	return value_filter.NewEqFilter[float64]()
}
