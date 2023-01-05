package errorx

import "errors"

// Defer is for use when defering a function call that can return an error.
// The error value of errp and the return value of f are joined with errors.Join.
func Defer(errp *error, f func() error) {
	newErr := f()
	*errp = errors.Join(*errp, newErr)
}
