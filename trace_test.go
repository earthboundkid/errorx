package errorx_test

import (
	"errors"
	"testing"

	"github.com/earthboundkid/errorx/v2"
)

func TestTrace(t *testing.T) {
	cause := errors.New("cause")
	e := cause
	errorx.Trace(&e)
	if e == cause {
		t.Fatal("error should not be cause (is of identity)")
	}
	if !errors.Is(e, cause) {
		t.Fatal("error should be a cause (is of predication)")
	}
}
