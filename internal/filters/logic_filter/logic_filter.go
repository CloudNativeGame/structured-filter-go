package logic_filter

import (
	"github.com/CloudNativeGame/structured-filter-go/pkg/types"
)

type ILogicFilter[T any] interface {
	types.IFilter[T]
}
