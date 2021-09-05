package errutil

import (
	"fmt"
	"path/filepath"
	"runtime"
)

// Prefix will prefix an error with a fixed string if it is non-nil.
func Prefix(errp *error, prefixformat string, a ...interface{}) {
	if err := *errp; err != nil {
		prefix := fmt.Sprintf(prefixformat, a...)
		*errp = fmt.Errorf("%s: %w", prefix, err)
	}
}

// Trace prefixes an error with caller information if the error is not nil.
func Trace(errp *error) {
	if err := *errp; err != nil {
		pc, file, line, ok := runtime.Caller(1)
		prefix := "debug information not available"
		if ok {
			f := runtime.FuncForPC(pc)
			file = filepath.Base(file)
			prefix = fmt.Sprintf("problem in %s (%s:%d)",
				f.Name(), file, line)
		}
		*errp = fmt.Errorf("%s: %w", prefix, err)
	}
}
