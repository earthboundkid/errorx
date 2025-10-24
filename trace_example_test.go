package errorx_test

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/earthboundkid/errorx/v2"
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

func printErr(err error) {
	if err == nil {
		fmt.Println(nil)
		return
	}
	// For test stability, replace line numbers
	re := regexp.MustCompile(`:(\d+)`)
	fmt.Println(re.ReplaceAllString(err.Error(), ":XX"))
}

func ExampleTrace() {
	printErr(traceErr1(true))
	printErr(traceErr1(false))
	printErr(traceErr2(1, -1))
	printErr(traceErr2(1, 1))
	// Output:
	// <nil>
	// @github.com/earthboundkid/errorx/v2_test.traceErr1 (trace_example_test.go:XX)
	// oh no!
	// <nil>
	// @github.com/earthboundkid/errorx/v2_test.traceErr2 (trace_example_test.go:XX)
	// uh oh!
}
