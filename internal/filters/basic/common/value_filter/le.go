package value_filter

import (
	"cmp"
	"github.com/CloudNativeGame/structured-filter-go/internal/consts"
	internaltypes "github.com/CloudNativeGame/structured-filter-go/internal/types"
	"github.com/CloudNativeGame/structured-filter-go/pkg/checkers"
	"github.com/CloudNativeGame/structured-filter-go/pkg/errors"
	"github.com/CloudNativeGame/structured-filter-go/pkg/types"
)

type LeFilter[T cmp.Ordered] struct {
}

func NewLeFilter[T cmp.Ordered]() LeFilter[T] {
	return LeFilter[T]{}
}

func (l LeFilter[T]) GetKey() string {
	return consts.LeKey
}

func (l LeFilter[T]) Valid(element types.JsonElement) errors.FilterError {
	return checkers.CheckElementType(l, element)
}

func (l LeFilter[T]) Match(element types.JsonElement, matchTarget T) errors.FilterError {
	if element.(T) >= matchTarget {
		return nil
	}

	return internaltypes.NewNotMatchError(l, matchTarget, element, nil)
}
