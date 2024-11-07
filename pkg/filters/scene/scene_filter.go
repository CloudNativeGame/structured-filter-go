package scene

import "github.com/CloudNativeGame/structured-filter-go/pkg/types"

type ISceneFilter[T any] interface {
	types.IFilter[T]
}
