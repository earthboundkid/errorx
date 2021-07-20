package errutil

import (
	"errors"
	"fmt"
	"strings"
)

// Merge is a convenience method for making a Slice of errors and calling the Merge method.
func Merge(errs ...error) error {
	s := Slice(errs)
	return s.Merge()
}

// Slice is a slice of errors. Use it to collect possible errors
// then create a Multierr with the Merge method.
type Slice []error

// AsSlice converts err into Slice. If err is nil, the slice has length 0.
// If the err is a Multierr, it returns the underlying Slice.
// All other errors become a slice of length 1.
func AsSlice(err error) Slice {
	if err == nil {
		return nil
	}
	if me := (Multierr)(nil); errors.As(err, &me) {
		return me.Errors()
	}
	return Slice{err}
}

// Push extends a Slice with an error if the error is non-nil.
//
// If a Multierr is passed to Push, the result is flattened.
func (s *Slice) Push(err error) {
	s2 := AsSlice(err)
	*s = append(*s, s2...)
}

// Merge first removes any nil errors from the Slice.
// If the resulting length of the Slice is zero, it returns nil.
// If there is only one error, it returns that error as is.
// If there are multiple errors, it returns a Multierr
// containing all the errors.
func (s *Slice) Merge() error {
	// Making a copy in case we need to flatten a nested Slice
	errsFiltered := make(Slice, 0, len(*s))
	for _, err := range *s {
		errsFiltered.Push(err)
	}
	*s = errsFiltered
	if len(errsFiltered) < 1 {
		return nil
	}
	if len(errsFiltered) == 1 {
		return (errsFiltered)[0]
	}
	return multierr{errsFiltered}
}

// Multierr is an interface allowing external types containing
// multiple errors (such as uber-go/multierr) to be treated as a Slice.
type Multierr interface {
	error
	Errors() []error
}

// multierr wraps multiple errors.
type multierr struct {
	s Slice
}

var _ Multierr = multierr{}

// Errors fulfills Multierr
func (m multierr) Errors() []error {
	return m.s
}

// Strings returns the strings from the underlying errors.
func (m multierr) Strings() []string {
	return errorsToStrings(m.s)
}

// Error implements the error interface.
func (m multierr) Error() string {
	a := m.Strings()
	if len(a) == 0 {
		return "<empty error slice>"
	}
	plural := "s"
	if len(a) == 1 {
		plural = ""
	}
	return fmt.Sprintf("%d error%s: %s", len(a), plural, strings.Join(a, "; "))
}

func errorsToStrings(s []error) []string {
	a := make([]string, 0, len(s))
	for _, err := range s {
		if s != nil {
			a = append(a, err.Error())
		}
	}
	return a
}

var _ fmt.Formatter = multierr{}

// Format implements fmt.Formatter. Adds %+v verb to print sub-errors.
func (m multierr) Format(state fmt.State, verb rune) {
	switch verb {
	case 's', 'q', 'v':
		if verb == 'v' && state.Flag('+') {
			a := m.Strings()
			if len(a) == 0 {
				fmt.Fprint(state, "<empty error slice>")
				return
			}

			plural := "s"
			if len(a) == 1 {
				plural = ""
			}
			fmt.Fprintf(state, "%d error%s:\n", len(a), plural)
			for i, err := range a {
				fmt.Fprintf(state, "\terror %d: %s\n", i+1, err)
			}
		} else {
			msg := m.Error()
			if verb == 'q' {
				msg = fmt.Sprintf("%q", msg)
			}
			fmt.Fprint(state, msg)
		}
	}
}
