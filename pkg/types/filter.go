package types

import (
	"github.com/CloudNativeGame/structured-filter-go/pkg/errors"
)

type IFilter[T any] interface {
	GetKey() string
	Valid(element JsonElement) errors.FilterError
	Match(element JsonElement, matchTarget T) errors.FilterError
}
