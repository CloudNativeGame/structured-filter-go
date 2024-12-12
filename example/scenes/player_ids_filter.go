package scenes

import (
	"github.com/CloudNativeGame/structured-filter-go/example/models"
	"github.com/CloudNativeGame/structured-filter-go/pkg/factory"
	"github.com/CloudNativeGame/structured-filter-go/pkg/filters/scene_filter"
	"github.com/CloudNativeGame/structured-filter-go/pkg/filters/scenes"
)

func NewPlayerIdsFilter(filterFactory *factory.FilterFactory[*models.Player]) scene_filter.ISceneFilter[*models.Player] {
	return scenes.NewNumberArraySceneFilter[*models.Player]("playerIds", func(p *models.Player) []float64 {
		ids := make([]float64, 0, len(p.Ids))
		for _, id := range p.Ids {
			ids = append(ids, float64(id))
		}
		return ids
	}, filterFactory)
}
