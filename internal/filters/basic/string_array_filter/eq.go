package string_array_filter

import (
	"github.com/CloudNativeGame/structured-filter-go/internal/filters/basic/common/array_filter"
)

func NewStringArrayEqFilter() IStringArrayFilter {
	return array_filter.NewEqFilter[[]string, string]()
}
