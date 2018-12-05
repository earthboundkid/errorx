package errors

func Defer(err *error, f func() error) {
	newErr := f()
	if *err == nil {
		*err = newErr
	}
}
