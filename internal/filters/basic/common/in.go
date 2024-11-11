package common

import (
	"github.com/CloudNativeGame/structured-filter-go/internal/consts"
	internaltypes "github.com/CloudNativeGame/structured-filter-go/internal/types"
	"github.com/CloudNativeGame/structured-filter-go/pkg/checkers"
	"github.com/CloudNativeGame/structured-filter-go/pkg/errors"
	"github.com/CloudNativeGame/structured-filter-go/pkg/types"
)

type InFilter[T comparable] struct {
}

func NewInFilter[T comparable]() InFilter[T] {
	return InFilter[T]{}
}

func (b InFilter[T]) GetKey() string {
	return consts.InKey
}

func (b InFilter[T]) Valid(element types.JsonElement) errors.FilterError {
	return checkers.CheckIsValidArray(b, element, nil, true)
}

func (b InFilter[T]) Match(element types.JsonElement, matchTarget T) errors.FilterError {
	arr := element.([]interface{})
	for _, val := range arr {
		if matchTarget == val.(T) {
			return nil
		}
	}

	return internaltypes.NewNotMatchError(b, matchTarget, element, nil)
}
