package number_filter

import (
	"github.com/CloudNativeGame/structured-filter-go/pkg/types"
)

type INumberFilter interface {
	types.IFilter[float64]
}
