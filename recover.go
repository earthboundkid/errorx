package errutil

import "fmt"

// Recover catches any panics and converts them to errors when deferred.
func Recover(errp *error) {
	o := recover()
	if o == nil {
		return
	}
	err, ok := o.(error)
	if !ok {
		err = fmt.Errorf("panic: %v", o)
	}
	*errp = err
}
