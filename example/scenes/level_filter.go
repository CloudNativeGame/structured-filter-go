package scenes

import (
	"github.com/CloudNativeGame/structured-filter-go/example/models"
	"github.com/CloudNativeGame/structured-filter-go/pkg/checkers"
	"github.com/CloudNativeGame/structured-filter-go/pkg/errors"
	"github.com/CloudNativeGame/structured-filter-go/pkg/factory"
	"github.com/CloudNativeGame/structured-filter-go/pkg/types"
)

type LevelFilter struct {
	filterFactory *factory.FilterFactory[*models.Player]
}

func (u *LevelFilter) GetKey() string {
	return "level"
}

func (u *LevelFilter) Valid(element types.JsonElement) errors.FilterError {
	return checkers.CheckIsValidObject(u, element, func(propertyKey string, propertyValue interface{}) errors.FilterError {
		filter, err := u.filterFactory.NumberFilterFactory.Get(propertyKey)
		if err != nil {
			return err
		}
		return filter.Valid(propertyValue)
	})
}

func (u *LevelFilter) Match(element types.JsonElement, matchTarget *models.Player) errors.FilterError {
	for k, v := range element.(map[string]interface{}) {
		filter, err := u.filterFactory.NumberFilterFactory.Get(k)
		if err != nil {
			return err
		}
		err = filter.Match(v, float64(matchTarget.Level))
		if err != nil {
			return err
		}
	}
	return nil
}

func NewLevelFilter(filterFactory *factory.FilterFactory[*models.Player]) *LevelFilter {
	return &LevelFilter{
		filterFactory: filterFactory,
	}
}
