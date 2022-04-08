package errutil_test

import (
	"errors"
	"testing"

	"github.com/carlmjohnson/be"
	"github.com/carlmjohnson/errutil"
	herr "github.com/hashicorp/go-multierror"
)

func TestHashicorp(t *testing.T) {
	err1 := errors.New("1")
	err2 := errors.New("2")
	errs := herr.Append(err1, err2)
	s := errutil.AsSlice(errs)

	be.Equal(t, 2, len(s))
	be.DeepEqual(t, err1, s[0])
	be.DeepEqual(t, err2, s[1])
}
