package errutil_test

import (
	"errors"
	"fmt"

	"github.com/carlmjohnson/errutil"
)

func ExamplePrefix() {
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

func maybeErr1(ok bool) (err error) {
	defer errutil.Prefix(&err, "maybeErr1")
	if !ok {
		return errors.New("oh no!")
	}
	return nil
}

func maybeErr2(x, y int) (err error) {
	defer errutil.Prefix(&err, "maybeErr2(%d, %d)", x, y)
	if x+y > 1 {
		return errors.New("uh oh!")
	}
	return nil
}
