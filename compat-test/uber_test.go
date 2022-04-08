package errutil_test

import (
	"errors"
	"testing"

	"github.com/carlmjohnson/be"
	"github.com/carlmjohnson/errutil"
	"go.uber.org/multierr"
)

func TestUber(t *testing.T) {
	err1 := errors.New("1")
	err2 := errors.New("2")
	errs := multierr.Append(err1, err2)
	s := errutil.AsSlice(errs)

	be.Equal(t, 2, len(s))
	be.DeepEqual(t, err1, s[0])
	be.DeepEqual(t, err2, s[1])
}
