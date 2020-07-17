package errutil_test

import (
	"errors"
	"fmt"
	"time"

	"github.com/carlmjohnson/errutil"
)

func ExampleExecParallel() {
	start := time.Now()
	err := errutil.ExecParallel(func() error {
		time.Sleep(1 * time.Second)
		return nil
	}, func() error {
		time.Sleep(1 * time.Second)
		return errors.New("one error")
	})
	fmt.Println("ran parallel?", time.Since(start) < 2*time.Second)
	fmt.Printf("err == %v\n", err)
	// Output:
	// ran parallel? true
	// err == one error
}
