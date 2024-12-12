package string_array_filter

import (
	"github.com/CloudNativeGame/structured-filter-go/pkg/types"
)

type IStringArrayFilter interface {
	types.IArrayFilter[[]string, string]
}
