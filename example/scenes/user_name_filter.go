package scenes

import (
	"github.com/CloudNativeGame/structured-filter-go/example/models"
	"github.com/CloudNativeGame/structured-filter-go/pkg/checkers"
	"github.com/CloudNativeGame/structured-filter-go/pkg/errors"
	"github.com/CloudNativeGame/structured-filter-go/pkg/factory"
	"github.com/CloudNativeGame/structured-filter-go/pkg/types"
)

type UserNameFilter struct {
	filterFactory *factory.FilterFactory[*models.Player]
}

func (u *UserNameFilter) GetKey() string {
	return "userName"
}

func (u *UserNameFilter) Valid(element types.JsonElement) errors.FilterError {
	return checkers.CheckIsValidObject(u, element, func(propertyKey string, propertyValue interface{}) errors.FilterError {
		filter, err := u.filterFactory.StringFilterFactory.Get(propertyKey)
		if err != nil {
			return err
		}
		return filter.Valid(propertyValue)
	})
}

func (u *UserNameFilter) Match(element types.JsonElement, matchTarget *models.Player) errors.FilterError {
	for k, v := range element.(map[string]interface{}) {
		filter, err := u.filterFactory.StringFilterFactory.Get(k)
		if err != nil {
			return err
		}
		err = filter.Match(v, matchTarget.User.Name)
		if err != nil {
			return err
		}
	}
	return nil
}

func NewUserNameFilter(filterFactory *factory.FilterFactory[*models.Player]) *UserNameFilter {
	return &UserNameFilter{
		filterFactory: filterFactory,
	}
}
