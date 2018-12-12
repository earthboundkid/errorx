package errors

import (
	"fmt"
	"strings"
)

func Merge(errs ...error) error {
	s := Slice(errs)
	return s.Merge()
}

type Slice []error

func (s *Slice) Push(err error) {
	if s2, ok := err.(Slice); ok {
		*s = append(*s, s2...)
	} else if err != nil {
		*s = append(*s, err)
	}
}

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
