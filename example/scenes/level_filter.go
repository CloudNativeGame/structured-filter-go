package scenes

import (
	"github.com/CloudNativeGame/structured-filter-go/example/models"
	"github.com/CloudNativeGame/structured-filter-go/pkg/factory"
	"github.com/CloudNativeGame/structured-filter-go/pkg/filters/scene_filter"
	"github.com/CloudNativeGame/structured-filter-go/pkg/filters/scenes"
)

func NewLevelFilter(filterFactory *factory.FilterFactory[*models.Player]) scene_filter.ISceneFilter[*models.Player] {
	return scenes.NewNumberSceneFilter[*models.Player]("level", func(p *models.Player) float64 {
		return float64(p.Level)
	}, filterFactory)
}
