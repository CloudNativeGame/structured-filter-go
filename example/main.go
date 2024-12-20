package main

import (
	"encoding/json"
	"fmt"
	"github.com/CloudNativeGame/structured-filter-go/example/builder"
	"github.com/CloudNativeGame/structured-filter-go/example/models"
	"github.com/CloudNativeGame/structured-filter-go/example/scenes"
	"github.com/CloudNativeGame/structured-filter-go/pkg"
	filterbuilder "github.com/CloudNativeGame/structured-filter-go/pkg/builder"
	"github.com/CloudNativeGame/structured-filter-go/pkg/errors"
	"github.com/CloudNativeGame/structured-filter-go/pkg/factory"
	"github.com/CloudNativeGame/structured-filter-go/pkg/filters/scene_filter"
	"log/slog"
)

func checkErrAndPrint(err errors.FilterError, filter, matchTarget string) {
	if err != nil {
		slog.Error(fmt.Sprintf("[%s] filter %s match %s error: %v", err.Type(), filter, matchTarget, err))
	} else {
		slog.Info(fmt.Sprintf("filter %s match %s ok", filter, matchTarget))
	}
}

func main() {
	p := models.Player{
		Id: 1,
		User: models.User{
			Name:   "Scott",
			IsMale: true,
		},
		Ids:    []int{1, 2, 3, 4, 5},
		Level:  10,
		Labels: []string{"season 1", "region a"},
	}
	playerJsonBytes, _ := json.Marshal(p)
	playerJson := string(playerJsonBytes)

	filterService := pkg.NewFilterService[*models.Player]().
		WithSceneFilters([]pkg.SceneFilterCreator[*models.Player]{
			func(factory *factory.FilterFactory[*models.Player]) scene_filter.ISceneFilter[*models.Player] {
				return scenes.NewIsMaleFilter(factory)
			},
			func(factory *factory.FilterFactory[*models.Player]) scene_filter.ISceneFilter[*models.Player] {
				return scenes.NewUserNameFilter(factory)
			},
			func(factory *factory.FilterFactory[*models.Player]) scene_filter.ISceneFilter[*models.Player] {
				return scenes.NewLevelFilter(factory)
			},
			func(factory *factory.FilterFactory[*models.Player]) scene_filter.ISceneFilter[*models.Player] {
				return scenes.NewPlayerIdsFilter(factory)
			},
			func(factory *factory.FilterFactory[*models.Player]) scene_filter.ISceneFilter[*models.Player] {
				return scenes.NewPlayerLabelsFilter(factory)
			},
		})

	// should match
	filterBuilder := builder.NewPlayerFilterBuilder()
	filter := filterBuilder.
		IsMaleObject(filterbuilder.Eq(true)).
		Build()
	filterBuilder.Reset()
	//filter := "{\"isMale\": {\"$eq\": true}}"
	err := filterService.Match(filter, &p)
	checkErrAndPrint(err, filter, playerJson)

	// FilterDocument will be cached after first parse
	//filter := "{\"isMale\": {\"$eq\": true}}"
	err = filterService.Match(filter, &p)
	checkErrAndPrint(err, filter, playerJson)

	players := []*models.Player{
		&p,
		{
			User: models.User{
				Name:   "Alice",
				IsMale: false,
			},
			Level: 25,
		},
	}
	//filter := "{\"isMale\": {\"$eq\": true}}"
	filteredPlayers := filterService.FilterOut(filter, players)
	for _, player := range filteredPlayers {
		slog.Info(fmt.Sprintf("after FilterOut, filter %s match %+#v ok", filter, player))
	}

	filter = filterBuilder.UserNameObject(filterbuilder.Eq("Scott")).Build()
	filterBuilder.Reset()
	//filter = "{\"userName\": {\"$eq\": \"Scott\"}}"
	err = filterService.Match(filter, &p)
	checkErrAndPrint(err, filter, playerJson)

	filter = filterBuilder.UserNameObject(filterbuilder.StringRange([]string{"Sam", "Stella"})).Build()
	filterBuilder.Reset()
	//filter = "{\"userName\": {\"$range\": [\"Sam\", \"Stella\"]}}"
	err = filterService.Match(filter, &p)
	checkErrAndPrint(err, filter, playerJson)

	filter = filterBuilder.PlayerIds([]int{1, 2, 3, 4, 5}).Build()
	filterBuilder.Reset()
	//filter = "{\"playerIds\": [1, 2, 3, 4, 5]}"
	err = filterService.Match(filter, &p)
	checkErrAndPrint(err, filter, playerJson)

	filter = filterBuilder.PlayerIdsObject(filterbuilder.NumberArrayEq([]int{1, 2, 3, 4, 5})).Build()
	filterBuilder.Reset()
	//filter = "{\"playerIds\": {\"$eq\": [1, 2, 3, 4, 5]}}"
	err = filterService.Match(filter, &p)
	checkErrAndPrint(err, filter, playerJson)

	filter = filterBuilder.PlayerIdsObject(filterbuilder.NumberArrayAll([]int{2, 3})).Build()
	filterBuilder.Reset()
	//filter = "{\"playerIds\": {\"$all\": [2, 3]}}"
	err = filterService.Match(filter, &p)
	checkErrAndPrint(err, filter, playerJson)

	filter = filterBuilder.PlayerLabels([]string{"season 1", "region a"}).Build()
	filterBuilder.Reset()
	//filter = "{\"playerLabels\": [\"season 1\", \"region a\"]}"
	err = filterService.Match(filter, &p)
	checkErrAndPrint(err, filter, playerJson)

	filter = filterBuilder.PlayerLabelsObject(filterbuilder.StringArrayEq([]string{"season 1", "region a"})).Build()
	filterBuilder.Reset()
	//filter = "{\"playerIds\": {\"$eq\": [\"season 1\", \"region a\"]}}"
	err = filterService.Match(filter, &p)
	checkErrAndPrint(err, filter, playerJson)

	filter = filterBuilder.PlayerLabelsObject(filterbuilder.StringArrayAll([]string{"season 1"})).Build()
	filterBuilder.Reset()
	//filter = "{\"playerIds\": {\"$all\": [\"season 1\"]}}"
	err = filterService.Match(filter, &p)
	checkErrAndPrint(err, filter, playerJson)

	filter = filterBuilder.UserNameObject(filterbuilder.Regex("^S")).Build()
	filterBuilder.Reset()
	//filter = "{\"userName\": {\"$regex\": \"^S\"}}"
	err = filterService.Match(filter, &p)
	checkErrAndPrint(err, filter, playerJson)

	filter = filterBuilder.UserNameObject(filterbuilder.Ne("Bob")).Build()
	filterBuilder.Reset()
	//filter = "{\"userName\": {\"$ne\": \"Bob\"}}"
	err = filterService.Match(filter, &p)
	checkErrAndPrint(err, filter, playerJson)

	filter = filterBuilder.UserNameObject(filterbuilder.StringIn([]string{"Bob", "Scott", "Alice"})).Build()
	filterBuilder.Reset()
	//filter = "{\"userName\": {\"$in\": [\"Bob\", \"Scott\", \"Alice\"]}}"
	err = filterService.Match(filter, &p)
	checkErrAndPrint(err, filter, playerJson)

	filter = filterBuilder.LevelObject(filterbuilder.Eq(10)).Build()
	filterBuilder.Reset()
	//filter = "{\"level\": {\"$eq\": 10}}"
	err = filterService.Match(filter, &p)
	checkErrAndPrint(err, filter, playerJson)

	filter = filterBuilder.LevelObject(filterbuilder.Le(10)).Build()
	filterBuilder.Reset()
	//filter = "{\"level\": {\"$le\": 10}}"
	err = filterService.Match(filter, &p)
	checkErrAndPrint(err, filter, playerJson)

	filter = filterBuilder.LevelObject(filterbuilder.NumberRange([]int{5, 15})).Build()
	filterBuilder.Reset()
	//filter = "{\"level\": {\"$range\": [5, 15]}}"
	err = filterService.Match(filter, &p)
	checkErrAndPrint(err, filter, playerJson)

	filter = filterBuilder.LevelObject(filterbuilder.NumberIn([]int{10, 20, 30})).Build()
	filterBuilder.Reset()
	//filter = "{\"level\": {\"$in\": [10, 20, 30]}}"
	err = filterService.Match(filter, &p)
	checkErrAndPrint(err, filter, playerJson)

	filter = filterBuilder.LevelObject(filterbuilder.Ne(20)).Build()
	filterBuilder.Reset()
	//filter = "{\"level\": {\"$ne\": 20}}"
	err = filterService.Match(filter, &p)
	checkErrAndPrint(err, filter, playerJson)

	// support k:v, would transform to k:{"$eq":v} internal
	filter = filterBuilder.IsMale(true).Build()
	filterBuilder.Reset()
	//filter = "{\"isMale\": true}"
	err = filterService.Match(filter, &p)
	checkErrAndPrint(err, filter, playerJson)

	// support logic operator
	filter = filterBuilder.And().IsMaleObject(filterbuilder.Eq(true)).
		UserNameObject(filterbuilder.Eq("Scott")).Build()
	filterBuilder.Reset()
	//filter = "{\"$and\": [{\"isMale\": {\"$eq\": true}}, {\"userName\": {\"$eq\": \"Scott\"}}]}"
	err = filterService.Match(filter, &p)
	checkErrAndPrint(err, filter, playerJson)

	filter = filterBuilder.Or().IsMaleObject(filterbuilder.Eq(false)).
		UserNameObject(filterbuilder.Eq("Scott")).Build()
	filterBuilder.Reset()
	//filter = "{\"$or\": [{\"isMale\": {\"$eq\": false}}, {\"userName\": {\"$eq\": \"Scott\"}}]}"
	err = filterService.Match(filter, &p)
	checkErrAndPrint(err, filter, playerJson)

	filter = filterBuilder.Or().IsMale(false).
		UserName("Scott").Build()
	filterBuilder.Reset()
	//filter = "{\"$or\": [{\"isMale\": false}, {\"userName\": \"Scott\"}]}"
	err = filterService.Match(filter, &p)
	checkErrAndPrint(err, filter, playerJson)

	// should return not match FilterError
	filter = filterBuilder.IsMaleObject(filterbuilder.Eq(false)).Build()
	filterBuilder.Reset()
	//filter = "{\"isMale\": {\"$eq\": false}}"
	err = filterService.Match(filter, &p)
	checkErrAndPrint(err, filter, playerJson)

	filter = filterBuilder.And().IsMaleObject(filterbuilder.Eq(true)).
		UserNameObject(filterbuilder.Eq("Tom")).Build()
	filterBuilder.Reset()
	//filter = "{\"$and\": [{\"isMale\": {\"$eq\": true}}, {\"userName\": {\"$eq\": \"Tom\"}}]}"
	err = filterService.Match(filter, &p)
	checkErrAndPrint(err, filter, playerJson)

	filter = filterBuilder.Or().IsMaleObject(filterbuilder.Eq(false)).
		UserNameObject(filterbuilder.Eq("Tom")).Build()
	filterBuilder.Reset()
	//filter = "{\"$or\": [{\"isMale\": {\"$eq\": false}}, {\"userName\": {\"$eq\": \"Tom\"}}]}"
	err = filterService.Match(filter, &p)
	checkErrAndPrint(err, filter, playerJson)

	filter = filterBuilder.LevelObject(filterbuilder.Lt(10)).Build()
	filterBuilder.Reset()
	//filter = "{\"level\": {\"$lt\": 10}}"
	err = filterService.Match(filter, &p)
	checkErrAndPrint(err, filter, playerJson)

	filter = filterBuilder.UserNameObject(filterbuilder.StringRange([]string{"Zarah", "Zero"})).Build()
	filterBuilder.Reset()
	//filter = "{\"userName\": {\"$range\": [\"Zarah\", \"Zero\"]}}"
	err = filterService.Match(filter, &p)
	checkErrAndPrint(err, filter, playerJson)

	filter = filterBuilder.PlayerIds([]int{1, 2, 3, 4}).Build()
	filterBuilder.Reset()
	//filter = "{\"playerIds\": [1, 2, 3, 4]}"
	err = filterService.Match(filter, &p)
	checkErrAndPrint(err, filter, playerJson)

	filter = filterBuilder.PlayerIds([]int{5, 2, 3, 4, 1}).Build()
	filterBuilder.Reset()
	//filter = "{\"playerIds\": [5, 2, 3, 4, 1]}"
	err = filterService.Match(filter, &p)
	checkErrAndPrint(err, filter, playerJson)

	filter = filterBuilder.PlayerIdsObject(filterbuilder.NumberArrayEq([]int{2, 3, 4, 5})).Build()
	filterBuilder.Reset()
	//filter = "{\"playerIds\": {\"$eq\": [2, 3, 4, 5]}}"
	err = filterService.Match(filter, &p)
	checkErrAndPrint(err, filter, playerJson)

	filter = filterBuilder.PlayerIdsObject(filterbuilder.NumberArrayAll([]int{2, 3, 6})).Build()
	filterBuilder.Reset()
	//filter = "{\"playerIds\": {\"$all\": [2, 3, 6]}}"
	err = filterService.Match(filter, &p)
	checkErrAndPrint(err, filter, playerJson)

	filter = filterBuilder.PlayerLabels([]string{"season 1"}).Build()
	filterBuilder.Reset()
	//filter = "{\"playerLabels\": [\"season 1\"]}"
	err = filterService.Match(filter, &p)
	checkErrAndPrint(err, filter, playerJson)

	filter = filterBuilder.PlayerLabels([]string{"season 1", "region b"}).Build()
	filterBuilder.Reset()
	//filter = "{\"playerLabels\": [\"season 1\", \"region b\"]}"
	err = filterService.Match(filter, &p)
	checkErrAndPrint(err, filter, playerJson)

	filter = filterBuilder.PlayerLabelsObject(filterbuilder.StringArrayEq([]string{"region a", "season 1"})).Build()
	filterBuilder.Reset()
	//filter = "{\"playerIds\": {\"$eq\": [\"region a\", \"season 1\"]}}"
	err = filterService.Match(filter, &p)
	checkErrAndPrint(err, filter, playerJson)

	filter = filterBuilder.PlayerLabelsObject(filterbuilder.StringArrayAll([]string{"season 1", "region b"})).Build()
	filterBuilder.Reset()
	//filter = "{\"playerIds\": {\"$all\": [\"season 1\", \"region b\"]}}"
	err = filterService.Match(filter, &p)
	checkErrAndPrint(err, filter, playerJson)

	// should return InvalidFilter error because Logic Filter array should only contain object elements
	filter = "{\"$or\": [false, \"Scott\"]}"
	err = filterService.Match(filter, &p)
	checkErrAndPrint(err, filter, playerJson)

	// should return InvalidFilter error because of wrong range syntax
	filter = filterBuilder.UserNameObject(filterbuilder.StringRange([]string{"Sam", "Stella", "Tom"})).Build()
	filterBuilder.Reset()
	//filter = "{\"userName\": {\"$range\": [\"Sam\", \"Stella\", \"Tom\"]}}"
	err = filterService.Match(filter, &p)
	checkErrAndPrint(err, filter, playerJson)

	// should return InvalidFilter error because of wrong range syntax
	filter = filterBuilder.UserNameObject(filterbuilder.StringRange([]string{"Stella", "Sam"})).Build()
	filterBuilder.Reset()
	//filter = "{\"userName\": {\"$range\": [\"Stella\", \"Sam\"]}}"
	err = filterService.Match(filter, &p)
	checkErrAndPrint(err, filter, playerJson)

	// should return InvalidFilter error because of wrong value type
	filter = "{\"isMale\": [true, false]}"
	err = filterService.Match(filter, &p)
	checkErrAndPrint(err, filter, playerJson)

	// should return InvalidFilter error because of wrong value type
	filter = "{\"isMale\": {\"$eq\": 1}}"
	err = filterService.Match(filter, &p)
	checkErrAndPrint(err, filter, playerJson)
}
