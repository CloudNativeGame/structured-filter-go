package array_filter

import (
	"github.com/CloudNativeGame/structured-filter-go/internal/consts"
	internaltypes "github.com/CloudNativeGame/structured-filter-go/internal/types"
	"github.com/CloudNativeGame/structured-filter-go/internal/utils"
	"github.com/CloudNativeGame/structured-filter-go/pkg/checkers"
	"github.com/CloudNativeGame/structured-filter-go/pkg/errors"
	"github.com/CloudNativeGame/structured-filter-go/pkg/types"
)

type AllFilter[TFilter any, TFilterElement comparable] struct {
}

func NewAllFilter[TFilter any, TFilterElement comparable]() AllFilter[TFilter, TFilterElement] {
	return AllFilter[TFilter, TFilterElement]{}
}

func (a AllFilter[TFilter, TFilterElement]) GetKey() string {
	return consts.AllKey
}

func (a AllFilter[TFilter, TFilterElement]) Valid(element types.JsonElement) errors.FilterError {
	return checkers.CheckIsValidArrayForArrayFilter[TFilter, TFilterElement](a, element, nil, true)
}

func (a AllFilter[TFilter, TFilterElement]) Match(element types.JsonElement, matchTarget []TFilterElement) errors.FilterError {
	s := utils.NewSet[TFilterElement]().FromSlice(matchTarget)
	for _, val := range element.([]interface{}) {
		if !s.Contains(val.(TFilterElement)) {
			return internaltypes.NewNotMatchError(a, matchTarget, element, nil)
		}
	}

	return nil
}
