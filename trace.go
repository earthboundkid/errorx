package errorx

import (
	"fmt"
	"path/filepath"
	"runtime"
)

// Trace joins the error with caller information if the error is not nil.
func Trace(errp *error) {
	if err := *errp; err != nil {
		pc, file, line, ok := runtime.Caller(1)
		te := traceErr{pc, file, line, ok, err}
		*errp = te
	}
}

type traceErr struct {
	pc    uintptr
	file  string
	line  int
	ok    bool
	cause error
}

func (te traceErr) Error() string {
	if !te.ok {
		return "debug information not available"
	}
	f := runtime.FuncForPC(te.pc)
	file := filepath.Base(te.file)

	return fmt.Sprintf("@%s (%s:%d)\n%v",
		f.Name(), file, te.line, te.cause)
}

func (te traceErr) Unwrap() error {
	return te.cause
}
