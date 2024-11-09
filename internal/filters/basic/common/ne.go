package common

import (
	"github.com/CloudNativeGame/structured-filter-go/internal/consts"
	internaltypes "github.com/CloudNativeGame/structured-filter-go/internal/types"
	"github.com/CloudNativeGame/structured-filter-go/pkg/checkers"
	"github.com/CloudNativeGame/structured-filter-go/pkg/errors"
	"github.com/CloudNativeGame/structured-filter-go/pkg/types"
)

type NeFilter[T comparable] struct {
}

func NewNeFilter[T comparable]() NeFilter[T] {
	return NeFilter[T]{}
}

func (b NeFilter[T]) GetKey() string {
	return consts.NeKey
}

func (b NeFilter[T]) Valid(element types.JsonElement) errors.FilterError {
	return checkers.CheckElementType(b, element)
}

func (b NeFilter[T]) Match(element types.JsonElement, matchTarget T) errors.FilterError {
	if element.(T) != matchTarget {
		return nil
	}

	return internaltypes.NewNotMatchError(b, matchTarget, element, nil)
}
