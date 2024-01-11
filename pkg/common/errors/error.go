package errors

type ErrorCode int

const (
	NotFound = iota
	Unauthorized
)

type ApplicationError struct {
	ErrorCode
	Message string
}
