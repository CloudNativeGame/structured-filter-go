package string_filter

import (
	"github.com/CloudNativeGame/structured-filter-go/internal/consts"
	internaltypes "github.com/CloudNativeGame/structured-filter-go/internal/types"
	"github.com/CloudNativeGame/structured-filter-go/pkg/checkers"
	"github.com/CloudNativeGame/structured-filter-go/pkg/errors"
	"github.com/CloudNativeGame/structured-filter-go/pkg/types"
	"regexp"
)

type RegexFilter struct {
}

func NewRegexFilter() RegexFilter {
	return RegexFilter{}
}

func (b RegexFilter) GetKey() string {
	return consts.RegexKey
}

func (b RegexFilter) Valid(element types.JsonElement) errors.FilterError {
	return checkers.CheckElementType(b, element)
}

func (b RegexFilter) Match(element types.JsonElement, matchTarget string) errors.FilterError {
	match, _ := regexp.MatchString(element.(string), matchTarget)
	if match {
		return nil
	}

	return internaltypes.NewNotMatchError(b, matchTarget, element, nil)
}
