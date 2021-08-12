package errutil

import "fmt"

// Prefix will prefix an error with a fixed string if it is non-nil.
func Prefix(errp *error, prefixformat string, a ...interface{}) {
	if err := *errp; err != nil {
		prefix := fmt.Sprintf(prefixformat, a...)
		*errp = fmt.Errorf("%s: %w", prefix, err)
	}
}
