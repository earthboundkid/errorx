package errutil

import "fmt"

// Prefix will prefix an error with a fixed string if it is non-nil.
func Prefix(errp *error, prefix string) {
	if err := *errp; err != nil {
		*errp = fmt.Errorf("%s: %w", prefix, err)
	}
}
