package value_filter

import (
	"cmp"
	"github.com/CloudNativeGame/structured-filter-go/internal/consts"
	internaltypes "github.com/CloudNativeGame/structured-filter-go/internal/types"
	"github.com/CloudNativeGame/structured-filter-go/pkg/checkers"
	"github.com/CloudNativeGame/structured-filter-go/pkg/errors"
	"github.com/CloudNativeGame/structured-filter-go/pkg/types"
)

type GeFilter[T cmp.Ordered] struct {
}

func NewGeFilter[T cmp.Ordered]() GeFilter[T] {
	return GeFilter[T]{}
}

func (g GeFilter[T]) GetKey() string {
	return consts.GeKey
}

func (g GeFilter[T]) Valid(element types.JsonElement) errors.FilterError {
	return checkers.CheckElementType(g, element)
}

func (g GeFilter[T]) Match(element types.JsonElement, matchTarget T) errors.FilterError {
	if element.(T) <= matchTarget {
		return nil
	}

	return internaltypes.NewNotMatchError(g, matchTarget, element, nil)
}
