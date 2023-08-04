package errorx_test

import (
	"errors"
	"fmt"

	"github.com/carlmjohnson/errorx"
)

func traceErr1(ok bool) (err error) {
	defer errorx.Trace(&err)
	if !ok {
		return errors.New("oh no!")
	}
	return nil
}

func traceErr2(x, y int) (err error) {
	defer errorx.Trace(&err)
	if x+y > 1 {
		return errors.New("uh oh!")
	}
	return nil
}

func ExampleTrace() {
	fmt.Println(traceErr1(true))
	fmt.Println(traceErr1(false))
	fmt.Println(traceErr2(1, -1))
	fmt.Println(traceErr2(1, 1))
	// Output:
	// <nil>
	// @github.com/carlmjohnson/errorx_test.traceErr1 (trace_example_test.go:13)
	// oh no!
	// <nil>
	// @github.com/carlmjohnson/errorx_test.traceErr2 (trace_example_test.go:21)
	// uh oh!
}
