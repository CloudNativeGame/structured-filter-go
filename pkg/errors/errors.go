package errors

type FilterErrorType string

const (
	InvalidFilter FilterErrorType = "InvalidFilter"
	NotMatch      FilterErrorType = "NotMatch"
)

type FilterError interface {
	Error() string

	Type() FilterErrorType
}
