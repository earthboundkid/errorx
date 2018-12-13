package errors

import (
	"fmt"
	"strings"
)

// Merge is a convenience method for making a Slice of errors and calling the Merge method.
func Merge(errs ...error) error {
	s := Slice(errs)
	return s.Merge()
}

// Slice is a slice of errors that implements the error interface itself.
type Slice []error

// Push extends a Slice with an error if the error is non-nil.
//
// If a Slice is passed to Push, the result is flattened.
func (s *Slice) Push(err error) {
	if s2, ok := err.(Slice); ok {
		*s = append(*s, s2...)
	} else if err != nil {
		*s = append(*s, err)
	}
}

// Merge removes any nil errors from the Slice, and then either
// returns nil if the length of the Slice is zero or returns
// the Slice if it is non-zero.
func (s *Slice) Merge() error {
	errsFiltered := (*s)[:0]
	for _, err := range *s {
		if err != nil {
			errsFiltered = append(errsFiltered, err)
		}
	}
	*s = errsFiltered
	if len(errsFiltered) < 1 {
		return nil
	}
	return s
}

var _ error = Slice{}

// Error implements the error interface.
func (s Slice) Error() string {
	a := s.errors()
	if len(a) == 0 {
		return "<empty error slice>"
	}
	plural := "s"
	if len(a) == 1 {
		plural = ""
	}
	return fmt.Sprintf("%d error%s: %s", len(a), plural, strings.Join(a, "; "))
}

func (s Slice) errors() []string {
	a := make([]string, 0, len(s))
	for _, err := range s {
		if s != nil {
			a = append(a, err.Error())
		}
	}
	return a
}

var _ fmt.Formatter = Slice{}

// Format implements fmt.Formatter.
func (s Slice) Format(state fmt.State, verb rune) {
	switch verb {
	case 's', 'q', 'v':
		if verb == 'v' && state.Flag('#') {
			a := s.errors()
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
			msg := s.Error()
			if verb == 'q' {
				msg = fmt.Sprintf("%q", msg)
			}
			fmt.Fprint(state, msg)
		}
	}
}
