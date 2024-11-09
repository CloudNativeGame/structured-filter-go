package string_filter

import (
	"github.com/CloudNativeGame/structured-filter-go/internal/filters/basic/common"
)

func NewStringEqFilter() IStringFilter {
	return common.NewEqFilter[string]()
}
