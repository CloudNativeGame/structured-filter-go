package number_filter

import (
	"github.com/CloudNativeGame/structured-filter-go/internal/filters/basic/common/value_filter"
)

func NewNumberLeFilter() INumberFilter {
	return value_filter.NewLeFilter[float64]()
}
