package string_filter

import (
	"github.com/CloudNativeGame/structured-filter-go/internal/filters/basic/common/value_filter"
)

func NewStringNeFilter() IStringFilter {
	return value_filter.NewNeFilter[string]()
}
