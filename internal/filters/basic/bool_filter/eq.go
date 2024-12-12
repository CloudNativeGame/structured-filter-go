package bool_filter

import (
	"github.com/CloudNativeGame/structured-filter-go/internal/filters/basic/common/value_filter"
)

func NewBoolEqFilter() IBoolFilter {
	return value_filter.NewEqFilter[bool]()
}
