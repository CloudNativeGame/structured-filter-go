package number_filter

import (
	"github.com/CloudNativeGame/structured-filter-go/internal/filters/basic/common"
)

func NewNumberGeFilter() INumberFilter {
	return common.NewGeFilter[float64]()
}
