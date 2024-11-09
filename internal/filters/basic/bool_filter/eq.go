package bool_filter

import "github.com/CloudNativeGame/structured-filter-go/internal/filters/basic/common"

func NewBoolEqFilter() IBoolFilter {
	return common.NewEqFilter[bool]()
}
