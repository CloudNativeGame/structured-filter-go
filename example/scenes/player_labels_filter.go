package scenes

import (
	"github.com/CloudNativeGame/structured-filter-go/example/models"
	"github.com/CloudNativeGame/structured-filter-go/pkg/factory"
	"github.com/CloudNativeGame/structured-filter-go/pkg/filters/scene_filter"
	"github.com/CloudNativeGame/structured-filter-go/pkg/filters/scenes"
)

func NewPlayerLabelsFilter(filterFactory *factory.FilterFactory[*models.Player]) scene_filter.ISceneFilter[*models.Player] {
	return scenes.NewStringArraySceneFilter[*models.Player]("playerLabels", func(p *models.Player) []string {
		return p.Labels
	}, filterFactory)
}
