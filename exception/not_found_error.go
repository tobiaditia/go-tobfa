package exception

/**
* mengikuti kontrak interface error go
* type error interface {
* 	Error() string
* }
 */
type NotFoundError struct {
	Error string
}

func NewNotFoundError(error string) NotFoundError {
	return NotFoundError{Error: error}
}
