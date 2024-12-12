package number_filter

import (
	"github.com/CloudNativeGame/structured-filter-go/internal/filters/basic/common/value_filter"
)

func NewNumberInFilter() INumberFilter {
	return value_filter.NewInFilter[float64]()
}
