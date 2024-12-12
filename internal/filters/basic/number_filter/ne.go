package number_filter

import (
	"github.com/CloudNativeGame/structured-filter-go/internal/filters/basic/common/value_filter"
)

func NewNumberNeFilter() INumberFilter {
	return value_filter.NewNeFilter[float64]()
}
