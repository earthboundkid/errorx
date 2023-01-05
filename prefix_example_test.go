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
	// problem in github.com/carlmjohnson/errorx_test.traceErr1 (prefix_example_test.go:13): oh no!
	// <nil>
	// problem in github.com/carlmjohnson/errorx_test.traceErr2 (prefix_example_test.go:21): uh oh!
}

func ExamplePrefix() {
	maybeErr1 := func(ok bool) (err error) {
		defer errorx.Prefix(&err, "maybeErr1")
		if !ok {
			return errors.New("oh no!")
		}
		return nil
	}

	maybeErr2 := func(x, y int) (err error) {
		defer errorx.Prefix(&err, "maybeErr2(%d, %d)", x, y)
		if x+y > 1 {
			return errors.New("uh oh!")
		}
		return nil
	}
	fmt.Println(maybeErr1(true))
	fmt.Println(maybeErr1(false))
	fmt.Println(maybeErr2(1, -1))
	fmt.Println(maybeErr2(1, 1))
	// Output:
	// <nil>
	// maybeErr1: oh no!
	// <nil>
	// maybeErr2(1, 1): uh oh!
}
