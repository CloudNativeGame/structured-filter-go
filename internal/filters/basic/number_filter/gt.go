package number_filter

import (
	"github.com/CloudNativeGame/structured-filter-go/internal/filters/basic/common/value_filter"
)

func NewNumberGtFilter() INumberFilter {
	return value_filter.NewGtFilter[float64]()
}
