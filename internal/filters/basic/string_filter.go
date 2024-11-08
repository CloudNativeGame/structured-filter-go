package basic

import (
	"github.com/CloudNativeGame/structured-filter-go/internal/consts"
	internaltypes "github.com/CloudNativeGame/structured-filter-go/internal/types"
	"github.com/CloudNativeGame/structured-filter-go/pkg/checkers"
	"github.com/CloudNativeGame/structured-filter-go/pkg/errors"
	"github.com/CloudNativeGame/structured-filter-go/pkg/types"
)

type IStringFilter interface {
	types.IFilter[string]
}

type StringEqFilter struct {
}

func NewStringEqFilter() StringEqFilter {
	return StringEqFilter{}
}

func (b StringEqFilter) GetKey() string {
	return consts.EqKey
}

func (b StringEqFilter) Valid(element types.JsonElement) errors.FilterError {
	return checkers.CheckIsValidString(b, element)
}

func (b StringEqFilter) Match(element types.JsonElement, matchTarget string) errors.FilterError {
	if element.(string) == matchTarget {
		return nil
	}

	return internaltypes.NewNotMatchError(b, matchTarget, element, nil)
}
