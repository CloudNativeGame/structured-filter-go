package string_array_filter

import (
	"github.com/CloudNativeGame/structured-filter-go/internal/filters/basic/common/array_filter"
)

func NewStringArrayAllFilter() IStringArrayFilter {
	return array_filter.NewAllFilter[[]string, string]()
}
