package exception

/**
* mengikuti kontrak interface error go
* type error interface {
* 	Error() string
* }
 */
type AlreadyExistsError struct {
	Error string
}

func NewAlreadyExistsError(error string) AlreadyExistsError {
	return AlreadyExistsError{Error: error}
}
