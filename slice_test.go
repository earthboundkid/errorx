package errutil

import (
	"fmt"
	"testing"
)

func TestPush(t *testing.T) {
	var errs Slice

	errs.Push(nil)
	checkLength(t, errs, 0)

	err := fmt.Errorf("")
	errs.Push(err)
	checkLength(t, errs, 1)

	errs.Push(nil)
	errs.Push(err)
	checkLength(t, errs, 2)

	errs.Push(errs.Merge())
	checkLength(t, errs, 4)
}

func checkLength(t *testing.T, errs Slice, want int) {
	t.Helper()
	if len(errs) != want {
		t.Errorf("wrong length for Slice: want %d, got %d", want, len(errs))
	}
}
