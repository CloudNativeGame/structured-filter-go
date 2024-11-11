package number_filter

import (
	"github.com/CloudNativeGame/structured-filter-go/internal/consts"
	internaltypes "github.com/CloudNativeGame/structured-filter-go/internal/types"
	"github.com/CloudNativeGame/structured-filter-go/internal/utils"
	"github.com/CloudNativeGame/structured-filter-go/pkg/checkers"
	"github.com/CloudNativeGame/structured-filter-go/pkg/errors"
	"github.com/CloudNativeGame/structured-filter-go/pkg/types"
)

type NumberRangeFilter[T float64] struct {
}

func NewNumberRangeFilter() NumberRangeFilter[float64] {
	return NumberRangeFilter[float64]{}
}

func (b NumberRangeFilter[T]) GetKey() string {
	return consts.RangeKey
}

func (b NumberRangeFilter[T]) Valid(element types.JsonElement) errors.FilterError {
	return checkers.CheckIsValidNumberRange(b, element)
}

func (b NumberRangeFilter[T]) Match(element types.JsonElement, matchTarget float64) errors.FilterError {
	filterRange := element.([]interface{})
	if matchTarget >= utils.ToFloat64(filterRange[0]) && matchTarget <= utils.ToFloat64(filterRange[1]) {
		return nil
	}

	return internaltypes.NewNotMatchError(b, matchTarget, element, nil)
}
