package bool_filter

import "github.com/CloudNativeGame/structured-filter-go/internal/filters/basic/common"

func NewBoolNeFilter() IBoolFilter {
	return common.NewNeFilter[bool]()
}
