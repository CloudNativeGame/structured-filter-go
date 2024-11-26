package number_filter

import (
	"github.com/CloudNativeGame/structured-filter-go/internal/filters/basic/common"
)

func NewNumberGtFilter() INumberFilter {
	return common.NewGtFilter[float64]()
}
