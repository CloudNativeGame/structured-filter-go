package number_filter

import (
	"github.com/CloudNativeGame/structured-filter-go/internal/filters/basic/common"
)

func NewNumberEqFilter() INumberFilter {
	return common.NewEqFilter[float64]()
}
