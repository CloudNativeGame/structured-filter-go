package bool_filter

import (
	"github.com/CloudNativeGame/structured-filter-go/internal/filters/basic/common/value_filter"
)

func NewBoolNeFilter() IBoolFilter {
	return value_filter.NewNeFilter[bool]()
}
