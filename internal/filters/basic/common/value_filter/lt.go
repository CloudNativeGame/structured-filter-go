package value_filter

import (
	"cmp"
	"github.com/CloudNativeGame/structured-filter-go/internal/consts"
	internaltypes "github.com/CloudNativeGame/structured-filter-go/internal/types"
	"github.com/CloudNativeGame/structured-filter-go/pkg/checkers"
	"github.com/CloudNativeGame/structured-filter-go/pkg/errors"
	"github.com/CloudNativeGame/structured-filter-go/pkg/types"
)

type LtFilter[T cmp.Ordered] struct {
}

func NewLtFilter[T cmp.Ordered]() LtFilter[T] {
	return LtFilter[T]{}
}

func (l LtFilter[T]) GetKey() string {
	return consts.LtKey
}

func (l LtFilter[T]) Valid(element types.JsonElement) errors.FilterError {
	return checkers.CheckElementType(l, element)
}

func (l LtFilter[T]) Match(element types.JsonElement, matchTarget T) errors.FilterError {
	if element.(T) > matchTarget {
		return nil
	}

	return internaltypes.NewNotMatchError(l, matchTarget, element, nil)
}
