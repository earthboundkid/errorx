package errorx

import (
	"errors"
	"fmt"
	"path/filepath"
	"runtime"
)

// Trace joins the error with caller information if the error is not nil.
func Trace(errp *error) {
	if err := *errp; err != nil {
		pc, file, line, ok := runtime.Caller(1)
		te := traceErr{pc, file, line, ok}
		*errp = errors.Join(&te, err)
	}
}

type traceErr struct {
	pc   uintptr
	file string
	line int
	ok   bool
}

func (te traceErr) Error() string {
	if !te.ok {
		return "debug information not available"
	}
	f := runtime.FuncForPC(te.pc)
	file := filepath.Base(te.file)

	return fmt.Sprintf("@%s (%s:%d)",
		f.Name(), file, te.line)
}
