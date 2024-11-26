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

func (r RegexFilter) GetKey() string {
	return consts.RegexKey
}

func (r RegexFilter) Valid(element types.JsonElement) errors.FilterError {
	return checkers.CheckElementType(r, element)
}

func (r RegexFilter) Match(element types.JsonElement, matchTarget string) errors.FilterError {
	match, _ := regexp.MatchString(element.(string), matchTarget)
	if match {
		return nil
	}

	return internaltypes.NewNotMatchError(r, matchTarget, element, nil)
}
