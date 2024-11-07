package main

import (
	"github.com/CloudNativeGame/structured-filter-go/example/models"
	"github.com/CloudNativeGame/structured-filter-go/example/scenes"
	"github.com/CloudNativeGame/structured-filter-go/pkg"
	"github.com/CloudNativeGame/structured-filter-go/pkg/factory"
	"github.com/CloudNativeGame/structured-filter-go/pkg/filters/scene"
	"log"
)

func main() {
	p := models.Player{
		User: models.User{
			Name:   "Scott",
			IsMale: true,
		},
	}

	filterFactory := factory.NewFilterFactory[*models.Player]()
	filterService := pkg.NewFilterService[*models.Player](filterFactory.WithSceneFilters([]scene.ISceneFilter[*models.Player]{
		scenes.NewIsMaleFilter(filterFactory),
	}))

	filter := "{\"isMale\": {\"$eq\": true}}"
	err := filterService.MatchFilter(filter, &p)
	if err != nil {
		log.Println("match error: ", err)
	}

	filter = "{\"isMale\": true}"
	err = filterService.MatchFilter(filter, &p)
	if err != nil {
		log.Println("match error: ", err)
	}

	filter = "{\"isMale\": [true, false]}"
	err = filterService.MatchFilter(filter, &p)
	if err != nil {
		log.Println("match error: ", err)
	}

	filter = "{\"isMale\": {\"$eq\": 1}}"
	err = filterService.MatchFilter(filter, &p)
	if err != nil {
		log.Println("match error: ", err)
	}

	filter = "{\"isMale\": {\"$eq\": false}}"
	err = filterService.MatchFilter(filter, &p)
	if err != nil {
		log.Println("match error: ", err)
	}
}
