package basic

import (
	"github.com/CloudNativeGame/structured-filter-go/internal/consts"
	internaltypes "github.com/CloudNativeGame/structured-filter-go/internal/types"
	"github.com/CloudNativeGame/structured-filter-go/pkg/checkers"
	"github.com/CloudNativeGame/structured-filter-go/pkg/errors"
	"github.com/CloudNativeGame/structured-filter-go/pkg/types"
)

type INumberFilter interface {
	types.IFilter[float64]
}

type NumberEqFilter struct {
}

func NewNumberEqFilter() NumberEqFilter {
	return NumberEqFilter{}
}

func (b NumberEqFilter) GetKey() string {
	return consts.EqKey
}

func (b NumberEqFilter) Valid(element types.JsonElement) errors.FilterError {
	return checkers.CheckIsValidNumber(b, element)
}

func (b NumberEqFilter) Match(element types.JsonElement, matchTarget float64) errors.FilterError {
	if element.(float64) == matchTarget {
		return nil
	}

	return internaltypes.NewNotMatchError(b, matchTarget, element, nil)
}
