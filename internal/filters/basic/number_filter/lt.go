package number_filter

import (
	"github.com/CloudNativeGame/structured-filter-go/internal/filters/basic/common"
)

func NewNumberLtFilter() INumberFilter {
	return common.NewLtFilter[float64]()
}
