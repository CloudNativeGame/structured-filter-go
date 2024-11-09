package string_filter

import (
	"github.com/CloudNativeGame/structured-filter-go/pkg/types"
)

type IStringFilter interface {
	types.IFilter[string]
}
