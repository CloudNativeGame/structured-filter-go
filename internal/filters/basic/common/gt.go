package common

import (
	"cmp"
	"github.com/CloudNativeGame/structured-filter-go/internal/consts"
	internaltypes "github.com/CloudNativeGame/structured-filter-go/internal/types"
	"github.com/CloudNativeGame/structured-filter-go/pkg/checkers"
	"github.com/CloudNativeGame/structured-filter-go/pkg/errors"
	"github.com/CloudNativeGame/structured-filter-go/pkg/types"
)

type GtFilter[T cmp.Ordered] struct {
}

func NewGtFilter[T cmp.Ordered]() GtFilter[T] {
	return GtFilter[T]{}
}

func (g GtFilter[T]) GetKey() string {
	return consts.GtKey
}

func (g GtFilter[T]) Valid(element types.JsonElement) errors.FilterError {
	return checkers.CheckElementType(g, element)
}

func (g GtFilter[T]) Match(element types.JsonElement, matchTarget T) errors.FilterError {
	if element.(T) > matchTarget {
		return nil
	}

	return internaltypes.NewNotMatchError(g, matchTarget, element, nil)
}
