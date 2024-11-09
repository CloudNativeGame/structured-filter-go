package string_filter

import (
	"github.com/CloudNativeGame/structured-filter-go/internal/filters/basic/common"
)

func NewStringNeFilter() IStringFilter {
	return common.NewNeFilter[string]()
}
