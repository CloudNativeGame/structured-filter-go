package string_filter

import (
	"github.com/CloudNativeGame/structured-filter-go/internal/consts"
	internaltypes "github.com/CloudNativeGame/structured-filter-go/internal/types"
	"github.com/CloudNativeGame/structured-filter-go/pkg/checkers"
	"github.com/CloudNativeGame/structured-filter-go/pkg/errors"
	"github.com/CloudNativeGame/structured-filter-go/pkg/types"
)

type StringRangeFilter[T float64] struct {
}

func NewStringRangeFilter() StringRangeFilter[float64] {
	return StringRangeFilter[float64]{}
}

func (s StringRangeFilter[T]) GetKey() string {
	return consts.RangeKey
}

func (s StringRangeFilter[T]) Valid(element types.JsonElement) errors.FilterError {
	return checkers.CheckIsValidStringRange(s, element)
}

func (s StringRangeFilter[T]) Match(element types.JsonElement, matchTarget string) errors.FilterError {
	filterRange := element.([]interface{})
	if matchTarget >= filterRange[0].(string) && matchTarget <= filterRange[1].(string) {
		return nil
	}

	return internaltypes.NewNotMatchError(s, matchTarget, element, nil)
}
