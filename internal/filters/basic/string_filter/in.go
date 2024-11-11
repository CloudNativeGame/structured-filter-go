package string_filter

import (
	"github.com/CloudNativeGame/structured-filter-go/internal/filters/basic/common"
)

func NewStringInFilter() IStringFilter {
	return common.NewInFilter[string]()
}
