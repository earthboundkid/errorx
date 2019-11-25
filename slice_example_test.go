package errutil_test

import (
	"fmt"

	"github.com/carlmjohnson/errutil"
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
	var errs errutil.Slice

	// Do something that returns an error sometimes
	err := someFunc()
	errs.Push(err)

	// Now merging them to produces <nil> error
	fmt.Println(errs.Merge())

	// But if we add non-nil errors...
	err = someFunc()
	errs.Push(err)
	errs.Push(err)

	// Merge returns non-nil
	fmt.Println(errs.Merge())
	// Output:
	// <nil>
	// 2 errors: even error!; even error!
}

func ExampleSlice_extendedFormat() {
	var errs errutil.Slice

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
	fmt.Printf("%+v", errs.Merge())

	// Output:
	// 3 errors:
	// 	error 1: error 1
	// 	error 2: error 2
	// 	error 3: error 3
}

func ExampleMerge() {
	// A function that sometimes returns an error
	called := 0
	someFunc := func() error {
		called++
		if called%2 == 0 {
			return fmt.Errorf("even error: %d!", called)
		}
		return nil
	}

	// We do a series of operations that might return an error.
	err := someFunc()

	// This time, it didn't return an error.
	fmt.Printf("%+v\n", err)

	// After each operation, we merge it into our existing error variable
	// then do the next operation.
	err = errutil.Merge(err, someFunc())
	err = errutil.Merge(err, someFunc())
	err = errutil.Merge(err, someFunc())

	// Finally, we return the result
	fmt.Printf("%+v", err)
	// Output:
	// <nil>
	// 2 errors:
	// 	error 1: even error: 2!
	// 	error 2: even error: 4!
}
