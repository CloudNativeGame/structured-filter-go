package string_filter

import (
	"github.com/CloudNativeGame/structured-filter-go/internal/filters/basic/common/value_filter"
)

func NewStringInFilter() IStringFilter {
	return value_filter.NewInFilter[string]()
}
