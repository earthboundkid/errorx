package errorx

// Defer is for use when defering a function call that can return an error.
// If the referenced error is nil but the callback returns a non-nil error,
// it sets the reference to the value of the returned error.
func Defer(errp *error, f func() error) {
	newErr := f()
	if *errp == nil {
		*errp = newErr
	}
}
