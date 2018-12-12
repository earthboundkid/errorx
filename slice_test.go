package errors_test

import (
	"fmt"

	"github.com/carlmjohnson/errors"
)

func ExampleSlice() {
	// A function that sometimes returns an error
	called := 0
	someFunc := func() error {
		called++
		if called%2 == 0 {
			return fmt.Errorf("even error!")
		}
		return nil
	}

	// The empty value can be used
	var errs errors.Slice

	// Do something that returns an error sometimes
	err := someFunc()
	errs.Push(err)

	// Now merging them to produces <nil> error
	fmt.Println(errs.Merge())

	// But if we add a non nil error...
	err = someFunc()
	errs.Push(err)

	// Merge returns non-nil
	fmt.Println(errs.Merge())
	// Output:
	// <nil>
	// 1 error: even error!
}
