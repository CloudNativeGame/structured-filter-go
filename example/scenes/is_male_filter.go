package scenes

import (
	"github.com/CloudNativeGame/structured-filter-go/example/models"
	"github.com/CloudNativeGame/structured-filter-go/pkg/checkers"
	"github.com/CloudNativeGame/structured-filter-go/pkg/types"

	"github.com/CloudNativeGame/structured-filter-go/pkg/errors"
	"github.com/CloudNativeGame/structured-filter-go/pkg/factory"
)

type IsMaleFilter struct {
	filterFactory *factory.FilterFactory[*models.Player]
}

func (i *IsMaleFilter) GetKey() string {
	return "isMale"
}

func (i *IsMaleFilter) Valid(element types.JsonElement) errors.FilterError {
	return checkers.CheckIsValidObject(i, element, func(propertyKey string, propertyValue interface{}) errors.FilterError {
		filter, err := i.filterFactory.BoolFilterFactory.Get(propertyKey)
		if err != nil {
			return err
		}
		return filter.Valid(propertyValue)
	})
}

func (i *IsMaleFilter) Match(element types.JsonElement, matchTarget *models.Player) errors.FilterError {
	for k, v := range element.(map[string]interface{}) {
		filter, err := i.filterFactory.BoolFilterFactory.Get(k)
		if err != nil {
			return err
		}
		err = filter.Match(v, matchTarget.User.IsMale)
		if err != nil {
			return err
		}
	}
	return nil
}

func NewIsMaleFilter(filterFactory *factory.FilterFactory[*models.Player]) *IsMaleFilter {
	return &IsMaleFilter{
		filterFactory: filterFactory,
	}
}
