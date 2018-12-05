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
	if err != nil {
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
	a := make([]string, 0, len(s))
	for _, err := range s {
		if s != nil {
			a = append(a, err.Error())
		}
	}
	if len(a) == 0 {
		return "<empty error slice>"
	}
	return fmt.Sprintf("%d errors: %s", len(a), strings.Join(a, "; "))
}
