package number_filter

import (
	"github.com/CloudNativeGame/structured-filter-go/internal/filters/basic/common/value_filter"
)

func NewNumberLtFilter() INumberFilter {
	return value_filter.NewLtFilter[float64]()
}
