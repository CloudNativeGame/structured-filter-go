package number_filter

import (
	"github.com/CloudNativeGame/structured-filter-go/internal/filters/basic/common"
)

func NewNumberInFilter() INumberFilter {
	return common.NewInFilter[float64]()
}
