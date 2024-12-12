package number_filter

import (
	"github.com/CloudNativeGame/structured-filter-go/internal/filters/basic/common/value_filter"
)

func NewNumberGeFilter() INumberFilter {
	return value_filter.NewGeFilter[float64]()
}
