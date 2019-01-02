package errors

import (
	"flag"
	"fmt"
	"os"
)

// ExitCoder is an optional interface that errors can implement to change
// what exit code they cause Execute to return.
type ExitCoder interface {
	ExitCode() int
}

// Execute is intended to be used as the top level function call of a CLI.
// It passes args to the callback. If args is nil, os.Args[1:] is substituted.
// If the callback returns a nil error, Execute does nothing and returns 0.
// Otherwise, Execute prints the error to stderr and returns non-zero. Specific
// error codes can be specified by implementing the ExitCoder optional interface.
func Execute(f func([]string) error, args []string) (exitCode int) {
	if args == nil {
		args = os.Args[1:]
	}
	err := f(args)
	if err == nil {
		return 0
	}
	// Special case: Let's assume flag.PrintDefaults() has already been called
	if err == flag.ErrHelp {
		return 2
	}
	if _, ok := err.(Slice); ok {
		fmt.Fprintf(os.Stderr, "Multiple runtime errors: %#v", err)
	} else {
		fmt.Fprintf(os.Stderr, "Runtime error: %v\n", err)
	}
	exitCode = 1
	if ec, ok := err.(ExitCoder); ok {
		exitCode = ec.ExitCode()
	}
	return exitCode
}
