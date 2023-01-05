package errorx_test

import (
	"errors"
	"fmt"
	"time"

	"github.com/carlmjohnson/errorx"
)

func ExampleExecParallel() {
	start := time.Now()
	err := errorx.ExecParallel(func() error {
		time.Sleep(1 * time.Second)
		return nil
	}, func() error {
		time.Sleep(1 * time.Second)
		return errors.New("one error")
	}, func() error {
		time.Sleep(1 * time.Second)
		panic("ahhh")
	})
	fmt.Println("ran parallel?", time.Since(start) < 2*time.Second)
	for i, suberr := range errorx.AsSlice(err) {
		fmt.Printf("error %d: %v\n", i+1, suberr)
	}
	// Output:
	// ran parallel? true
	// error 1: one error
	// error 2: panic: ahhh
}
