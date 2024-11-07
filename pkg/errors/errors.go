package errors

type FilterErrorType int

const (
	InvalidFilter FilterErrorType = 1
	NotMatch      FilterErrorType = 2
)

type FilterError interface {
	Error() string

	Type() FilterErrorType
}
