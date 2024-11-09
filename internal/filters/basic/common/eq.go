package common

import (
	"github.com/CloudNativeGame/structured-filter-go/internal/consts"
	internaltypes "github.com/CloudNativeGame/structured-filter-go/internal/types"
	"github.com/CloudNativeGame/structured-filter-go/pkg/checkers"
	"github.com/CloudNativeGame/structured-filter-go/pkg/errors"
	"github.com/CloudNativeGame/structured-filter-go/pkg/types"
)

type EqFilter[T comparable] struct {
}

func NewEqFilter[T comparable]() EqFilter[T] {
	return EqFilter[T]{}
}

func (b EqFilter[T]) GetKey() string {
	return consts.EqKey
}

func (b EqFilter[T]) Valid(element types.JsonElement) errors.FilterError {
	return checkers.CheckElementType(b, element)
}

func (b EqFilter[T]) Match(element types.JsonElement, matchTarget T) errors.FilterError {
	if element.(T) == matchTarget {
		return nil
	}

	return internaltypes.NewNotMatchError(b, matchTarget, element, nil)
}
