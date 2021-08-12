package errutil_test

import (
	"errors"
	"testing"

	"github.com/carlmjohnson/errutil"
	herr "github.com/hashicorp/go-multierror"
)

func TestHashicorp(t *testing.T) {
	err1 := errors.New("1")
	err2 := errors.New("2")
	errs := herr.Append(err1, err2)
	s := errutil.AsSlice(errs)
	if len(s) != 2 {
		t.Fatalf("len(s) == %d", len(s))
	}
	if s[0] != err1 {
		t.Fatalf("s[0] == %v", s[0])
	}
	if s[1] != err2 {
		t.Fatalf("s[1] == %v", s[1])
	}
}
