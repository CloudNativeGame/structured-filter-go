package bool_filter

import (
	"github.com/CloudNativeGame/structured-filter-go/pkg/types"
)

type IBoolFilter interface {
	types.IFilter[bool]
}
