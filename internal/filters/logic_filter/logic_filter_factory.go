package logic_filter

import (
	internaltypes "github.com/CloudNativeGame/structured-filter-go/internal/types"
	"github.com/CloudNativeGame/structured-filter-go/pkg/errors"
)

type LogicFilterFactory[T any] struct {
	logicFilters map[string]ILogicFilter[T]
}

func NewLogicFilterFactory[T any](logicFilters []ILogicFilter[T]) LogicFilterFactory[T] {
	logicFilterFactory := LogicFilterFactory[T]{
		logicFilters: make(map[string]ILogicFilter[T], len(logicFilters)),
	}
	for _, filter := range logicFilters {
		logicFilterFactory.logicFilters[filter.GetKey()] = filter
	}
	return logicFilterFactory
}

func (n LogicFilterFactory[T]) Get(key string) (ILogicFilter[T], errors.FilterError) {
	if filter, ok := n.logicFilters[key]; ok {
		return filter, nil
	}

	return nil, internaltypes.NewKeyNotFoundError(key)
}
