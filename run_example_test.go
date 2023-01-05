package errorx_test

import (
	"errors"
	"fmt"
	"time"

	"github.com/carlmjohnson/errorx"
)

func ExampleRun() {
	start := time.Now()
	err := errorx.Run(func() error {
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
	fmt.Printf("error: %q", err)
	// Output:
	// ran parallel? true
	// error: "one error\npanic: ahhh"
}
