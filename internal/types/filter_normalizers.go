package types

import (
	"github.com/CloudNativeGame/structured-filter-go/internal/consts"
	"strings"
)

func NormalizeFilter(filterMap map[string]interface{}) {
	valuesTransformToObjects(filterMap, 0)
}

func valuesTransformToObjects(filter map[string]interface{}, layer int) {
	itemsToTrans := make(map[string]interface{})
	for k, v := range filter {
		switch v.(type) {
		case map[string]interface{}:
			valuesTransformToObjects(v.(map[string]interface{}), layer+1)
		case []interface{}:
			if arrayValuesTransformToObjects(v.([]interface{}), layer+1) && layer == 0 {
				itemsToTrans[k] = v
			}
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

// 如果是值类型的数组，则返回 true，需要标准化为 $eq: []
func arrayValuesTransformToObjects(array []interface{}, layer int) bool {
	isValueArray := true
	for _, v := range array {
		switch v.(type) {
		case map[string]interface{}:
			valuesTransformToObjects(v.(map[string]interface{}), layer+1)
			isValueArray = false
		case []interface{}:
			arrayValuesTransformToObjects(v.([]interface{}), layer+1)
			isValueArray = false
		}
	}
	return isValueArray
}
