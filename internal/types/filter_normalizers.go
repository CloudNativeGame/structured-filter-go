package types

import (
	"github.com/CloudNativeGame/structured-filter-go/internal/consts"
	"strings"
)

func NormalizeFilter(filterMap map[string]interface{}) {
	valuesTransformToObjects(filterMap)
}

func valuesTransformToObjects(filter map[string]interface{}) {
	itemsToTrans := make(map[string]interface{})
	for k, v := range filter {
		switch v.(type) {
		case map[string]interface{}:
			valuesTransformToObjects(v.(map[string]interface{}))
		case []interface{}:
			arrayValuesTransformToObjects(v.([]interface{}))
		default:
			if !strings.HasPrefix(k, "$") {
				itemsToTrans[k] = v
			}
		}
	}

	for k, v := range itemsToTrans {
		delete(filter, k)
		filter[k] = map[string]interface{}{
			consts.EqKey: v,
		}
	}
}

func arrayValuesTransformToObjects(array []interface{}) {
	for _, v := range array {
		switch v.(type) {
		case map[string]interface{}:
			valuesTransformToObjects(v.(map[string]interface{}))
		case []interface{}:
			arrayValuesTransformToObjects(v.([]interface{}))
		}
	}
}
