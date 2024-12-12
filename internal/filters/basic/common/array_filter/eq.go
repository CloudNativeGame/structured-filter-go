package array_filter

import (
	"github.com/CloudNativeGame/structured-filter-go/internal/consts"
	internaltypes "github.com/CloudNativeGame/structured-filter-go/internal/types"
	"github.com/CloudNativeGame/structured-filter-go/pkg/checkers"
	"github.com/CloudNativeGame/structured-filter-go/pkg/errors"
	"github.com/CloudNativeGame/structured-filter-go/pkg/types"
)

type EqFilter[TFilter any, TFilterElement comparable] struct {
}

func NewEqFilter[TFilter any, TFilterElement comparable]() EqFilter[TFilter, TFilterElement] {
	return EqFilter[TFilter, TFilterElement]{}
}

func (a EqFilter[TFilter, TFilterElement]) GetKey() string {
	return consts.EqKey
}

func (a EqFilter[TFilter, TFilterElement]) Valid(element types.JsonElement) errors.FilterError {
	return checkers.CheckIsValidArrayForArrayFilter[TFilter, TFilterElement](a, element, nil, true)
}

func (a EqFilter[TFilter, TFilterElement]) Match(element types.JsonElement, matchTarget []TFilterElement) errors.FilterError {
	filterArr := element.([]interface{})
	if len(matchTarget) != len(filterArr) {
		return internaltypes.NewNotMatchError(a, matchTarget, element, nil)
	}

	for i, val := range filterArr {
		if matchTarget[i] != val {
			return internaltypes.NewNotMatchError(a, matchTarget, element, nil)
		}
	}

	return nil
}
