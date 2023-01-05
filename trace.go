package errorx

import (
	"fmt"
	"path/filepath"
	"runtime"
)

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
