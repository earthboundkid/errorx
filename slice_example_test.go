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

func ExampleSlice_extendedFormat() {
	var errs errors.Slice

	// Collect several errors
	err := fmt.Errorf("error 1")
	errs.Push(err)

	err = fmt.Errorf("error 2")
	errs.Push(err)

	// ...and a nil error
	err = nil
	errs.Push(err)

	// ...then a real error again
	err = fmt.Errorf("error 3")
	errs.Push(err)

	// Now merge and output them in extended format
	fmt.Printf("%#v", errs.Merge())

	// Output:
	// 3 errors:
	// 	error 1: error 1
	// 	error 2: error 2
	// 	error 3: error 3
}
