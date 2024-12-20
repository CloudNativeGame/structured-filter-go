# structured-filter-go

A general-purpose, business-agnostic structured filter library.

## How to use

* Suppose you need to filter out Players that meet certain conditions from a Player list.

```go
package main

type Player struct {
	User  User `json:"user"`
	Level int  `json:"level"`
	Tags []string `json:"tags"`
}

type User struct {
	Name   string `json:"name"`
	IsMale bool   `json:"isMale"`
}

func main() {
	players := []Player{
		{
			User: User{
				Name:   "Scott",
				IsMale: true,
			},
			Level: 10,
			Tags: []string{"season 1", "region a"},
		},
		{
			User: User{
				Name:   "Alice",
				IsMale: false,
			},
			Level: 25,
            Tags: []string{"season 1", "region b"},
		},
	}
}
```

* Define Scene filters for the object members you expect to match.

```go
func NewIsMaleFilter(filterFactory *factory.FilterFactory[*models.Player]) scene_filter.ISceneFilter[*models.Player] {
    return scenes.NewBoolSceneFilter[*models.Player]("isMale", func(p *models.Player) bool {
        return p.User.IsMale
    }, filterFactory)
}

func NewLevelFilter(filterFactory *factory.FilterFactory[*models.Player]) scene_filter.ISceneFilter[*models.Player] {
    return scenes.NewNumberSceneFilter[*models.Player]("level", func(p *models.Player) float64 {
        return float64(p.Level)
    }, filterFactory)
}

func NewUserNameFilter(filterFactory *factory.FilterFactory[*models.Player]) scene_filter.ISceneFilter[*models.Player] {
    return scenes.NewStringSceneFilter[*models.Player]("userName", func(p *models.Player) string {
        return p.User.Name
    }, filterFactory)
}
```

* Create your `FilterService[*models.Player]` and register the just defined filters.

```go
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
    })
```

* Use filter JSON string to filter your object list.

```go
// should match
filterBuilder := builder.NewPlayerFilterBuilder()
filter := filterBuilder.IsMaleObject(filterbuilder.Eq(true)).Build()
filterBuilder.Reset()
//filter := "{\"isMale\": {\"$eq\": true}}"
filteredPlayers := filterService.FilterOut(filter, players)
for _, player := range filteredPlayers {
    slog.Info(fmt.Sprintf("after FilterOut, filter %s match %+#v ok", filter, player))
}
// INFO after FilterOut, filter {"isMale":{"$eq":true}} match &models.Player{User:models.User{Name:"Scott", IsMale:true}, Level:10} ok
// INFO after FilterOut, filter {"isMale":{"$eq":true}} not match &models.Player{User:models.User{Name:"Alice", IsMale:false}, Level:25}
```

## Data types for StructuredFilter

### Object

* Object is a JSON object but allows exactly one key-value pair.
* Examples: `{"userName": "Scott"}`, `{"userName": {"$eq": "Scott"}}`
* Invalid: `{"userName": "Scott", "age": 20}`, `{"userName": {"$eq": "Scott"}, "age": {"$eq": 20}}`

### Array

* Array is a JSON array containing at least one element, and all elements need to be of the same type.
* Examples: `[{"userName": "Scott"}, {"age": 20}]`, `[1, 2, 3]`, `["a", "b", "c"]`
* Invalid: `[]`, `[{"userName": "Scott"}, 20]`, `[1, 2, "c"]`

### Range

* Range is a StructuredFilter Array containing two elements, the first element of the Array is less than or equal to the second element, and the elements need to be one of the following types:
    * string
    * double
    * Version
* The value of Range uses a closed interval, that is, the match is successful when it is greater than or equal to the first element and less than or equal to the second element.
* Examples: `[1, 2]`、`["a", "z"]`, `[Version.Parse("1.0.0"), Version.Parse("1.6.0")]`
* Invalid: `[]`, `[1]`, `[1, 2, 3]`, `[3, 1]`

## Filter types

### Logic filters

* Logic filters currently does not support nested Logic filters.

| Key  | Value Type                 | Description                                                          | Filter Examples                                                      |
|------|----------------------------|----------------------------------------------------------------------|----------------------------------------------------------------------|
| $and | Array with Object Elements | Match successfully when all elements of the array match successfully | `{"$and": [{"pid": {"$lt": 1000}}, {"userName": {"$eq": "Scott"}}]}` |
| $or  | Array with Object Elements | Match successfully when any element of the array match successfully  | `{"$or": [{"pid": {"$lt": 1000}}, {"userName": {"$eq": "Scott"}}]}`  |

### Basic filters

#### Bool filters

| Key | Value Type | Description                                                                             | Filter Examples             |
|-----|------------|-----------------------------------------------------------------------------------------|-----------------------------|
| $eq | bool       | Match successfully when the value of the matching object and filter value are equal     | `{"isMale": {"$eq": true}}` |
| $ne | bool       | Match successfully when the value of the matching object and filter value are not equal | `{"isMale": {"$ne": true}}` |

#### Number filters

| Key    | Value Type   | Description                                                                                                                                                                                   | Filter Examples                  |
|--------|--------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|----------------------------------|
| $eq    | double       | Match successfully when the value of the matching object and filter value are equal                                                                                                           | `{"age": {"$eq": 20}}`           |
| $ne    | double       | Match successfully when the value of the matching object and filter value are not equal                                                                                                       | `{"age": {"$ne": 20}}`           |
| $in    | double Array | Match successfully when the value of the matching object is equal to any of the filter values                                                                                                 | `{"age": {"$in": [20, 21, 22]}}` |
| $lt    | double       | Match successfully when the value of the matching object is less than the filter value                                                                                                        | `{"age": {"$lt": 20}}`           |
| $gt    | double       | Match successfully when the value of the matching object is greater than the filter value                                                                                                     | `{"age": {"$gt": 20}}`           |
| $le    | double       | Match successfully when the value of the matching object is less than or equal to the filter value                                                                                            | `{"age": {"$le": 20}}`           |
| $ge    | double       | Match successfully when the value of the matching object is greater than or equal to the filter value                                                                                         | `{"age": {"$ge": 20}}`           |
| $range | double Range | Match successfully when the value of the matching object is greater than or equal to the first element in the filter values and less than or equal to the second element in the filter values | `{"age": {"$range": [20, 30]}}`  |

#### Number Array filters

| Key  | Value Type   | Description                                                                                                                           | Filter Examples                            |
|------|--------------|---------------------------------------------------------------------------------------------------------------------------------------|--------------------------------------------|
| $eq  | double Array | Match successfully when each element of the matching array is equal to the element at each corresponding position of the filter value | `{"itemIds": {"$eq": [1000, 1001, 1002]}}` |
| $all | double Array | Match successfully when the matching array contains every element of the filter value                                                 | `{"itemIds": {"$all": [1000, 1005]}}`      |

#### String filters

| Key    | Value Type                 | Description                                                                                                                                                                                   | Filter Examples                                              |
|--------|----------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|--------------------------------------------------------------|
| $eq    | string                     | Match successfully when the value of the matching object and filter value are equal                                                                                                           | `{"userName": {"$eq": "Scott"}}`                             |
| $ne    | string                     | Match successfully when the value of the matching object and filter value are not equal                                                                                                       | `{"userName": {"$ne": "Scott"}}`                             |
| $in    | string Array               | Match successfully when the value of the matching object is equal to any of the filter values                                                                                                 | `{"userName": {"$in": ["Scott", "Tom", "Bob"]}}`             |
| $range | string Range               | Match successfully when the value of the matching object is greater than or equal to the first element in the filter values and less than or equal to the second element in the filter values | `{"serialNumber": {"$range": ["abcde00001", "abcde99999"]}}` |
| $regex | string(regular expression) | Match successfully when the value of the matching object match filter value as regular expression                                                                                             | `{"userName": {"$regex": "^S"}}`                             |

#### String Array filters

| Key  | Value Type   | Description                                                                                                                           | Filter Examples                               |
|------|--------------|---------------------------------------------------------------------------------------------------------------------------------------|-----------------------------------------------|
| $eq  | string Array | Match successfully when each element of the matching array is equal to the element at each corresponding position of the filter value | `{"tags": {"$eq": ["season 1", "region c"]}}` |
| $all | string Array | Match successfully when the matching array contains every element of the filter value                                                 | `{"tags": {"$all": ["season 3"]}}`            |

### Scene filters

* Scene filters are defined by the user and can be nested as subordinates of Logic filters or superior of Basic filters.
* The following classes of scene filters are provided to simplify the writing of your scene filters:
    * BoolSceneFilter
    * NumberSceneFilter
    * StringSceneFilter

## Filter Builders

Filter Builders are used to construct the filter string by chaining method calls.

```go
type IFilterBuilder interface {
  // Or And must be placed before any other method call, and only one of them can be called, at most once.
  Or() IFilterBuilder
  And() IFilterBuilder
  // KStringV KStringArrayV KBoolV KNumberV KNumberArrayV are used to build `$eq` filters of corresponding data types.
  KStringV(key string, value string) IFilterBuilder
  KStringArrayV(key string, value []string) IFilterBuilder
  KBoolV(key string, value bool) IFilterBuilder
  KNumberV(key string, value float64) IFilterBuilder
  KNumberArrayV(key string, value []float64) IFilterBuilder
  // KObjectV is used to build filters other than `$eq`.
  KObjectV(key string, value FilterBuilderObject) IFilterBuilder
  // Build is usually placed at the end of the chaining method calls and returns the filter string. It is reentrant.
  Build() string
  // Reset is used to clean all the filters in the `IFilterBuilder` if you need to build another filter string.
  Reset()
}
```

* `Or`/`And` must be placed before any other method call, and only one of them can be called, at most once.
* `KStringV`/`KStringArrayV`/`KBoolV`/`KNumberV`/`KNumberArrayV` are used to build `$eq` filters of corresponding data types when `KObjectV` is used to build other filters like `$lt`, `$range`, `$in` and all other filters mentioned above.
* `Build` is usually placed at the end of the chaining method calls and returns the filter string. `Build` is reentrant and you can call `Reset` to clean all the filters in the `IFilterBuilder` if you need to build another filter string.
