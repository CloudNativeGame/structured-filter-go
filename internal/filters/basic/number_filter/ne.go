package number_filter

import (
	"github.com/CloudNativeGame/structured-filter-go/internal/filters/basic/common"
)

func NewNumberNeFilter() INumberFilter {
	return common.NewNeFilter[float64]()
}
