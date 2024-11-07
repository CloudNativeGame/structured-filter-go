package document

import (
	"encoding/json"
	internalerrors "github.com/CloudNativeGame/structured-filter-go/internal/errors"
	internaltypes "github.com/CloudNativeGame/structured-filter-go/internal/types"
	"github.com/CloudNativeGame/structured-filter-go/pkg/errors"
	"github.com/CloudNativeGame/structured-filter-go/pkg/factory"
	"github.com/CloudNativeGame/structured-filter-go/pkg/types"
)

type FilterDocument[T any] struct {
	filterJson types.JsonElement
}

func NewFilterDocument[T any](rawFilter string, filterFactory factory.FilterFactory[T]) (*FilterDocument[T], errors.FilterError) {
	filterDocument := FilterDocument[T]{}

	var filterJson types.JsonElement
	err := json.Unmarshal([]byte(rawFilter), &filterJson)
	if err != nil {
		return nil, internalerrors.ToFilterError(err, errors.InvalidFilter)
	}

	filterMap, ok := filterJson.(map[string]interface{})
	if !ok {
		return nil, internalerrors.NewFilterError(errors.InvalidFilter, "filter should be map[string]interface{}")
	}
	if len(filterMap) != 1 {
		return nil, internalerrors.NewFilterError(errors.InvalidFilter, "filter should contain exact one element")
	}

	internaltypes.NormalizeFilter(filterMap)

	for k, v := range filterMap {
		filter, err := filterFactory.Get(k)
		if err != nil {
			return nil, err
		}
		err = filter.Valid(v)
		if err != nil {
			return nil, err
		}
	}

	filterDocument.filterJson = filterMap
	return &filterDocument, nil
}

func (f *FilterDocument[T]) Enumerate() map[string]interface{} {
	return f.filterJson.(map[string]interface{})
}
