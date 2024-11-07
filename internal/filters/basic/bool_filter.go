package basic

import (
	"github.com/CloudNativeGame/structured-filter-go/internal/consts"
	internaltypes "github.com/CloudNativeGame/structured-filter-go/internal/types"
	"github.com/CloudNativeGame/structured-filter-go/pkg/checkers"
	"github.com/CloudNativeGame/structured-filter-go/pkg/errors"
	"github.com/CloudNativeGame/structured-filter-go/pkg/types"
)

type IBoolFilter interface {
	types.IFilter[bool]
}

type BoolEqFilter struct {
}

func (b BoolEqFilter) GetKey() string {
	return consts.EqKey
}

func (b BoolEqFilter) Valid(element types.JsonElement) errors.FilterError {
	return checkers.CheckIsValidBool(b, element)
}

func (b BoolEqFilter) Match(element types.JsonElement, matchTarget bool) errors.FilterError {
	if element.(bool) == matchTarget {
		return nil
	}

	return internaltypes.NewNotMatchError(b, matchTarget, element)
}
