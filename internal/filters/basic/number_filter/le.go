package number_filter

import (
	"github.com/CloudNativeGame/structured-filter-go/internal/filters/basic/common"
)

func NewNumberLeFilter() INumberFilter {
	return common.NewLeFilter[float64]()
}
