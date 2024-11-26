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

func (i InFilter[T]) GetKey() string {
	return consts.InKey
}

func (i InFilter[T]) Valid(element types.JsonElement) errors.FilterError {
	return checkers.CheckIsValidArray(i, element, nil, true)
}

func (i InFilter[T]) Match(element types.JsonElement, matchTarget T) errors.FilterError {
	arr := element.([]interface{})
	for _, val := range arr {
		if matchTarget == val.(T) {
			return nil
		}
	}

	return internaltypes.NewNotMatchError(i, matchTarget, element, nil)
}
