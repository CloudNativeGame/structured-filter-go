package number_array_filter

import (
	"github.com/CloudNativeGame/structured-filter-go/pkg/types"
)

type INumberArrayFilter interface {
	types.IArrayFilter[[]float64, float64]
}
