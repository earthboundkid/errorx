package errutil_test

import (
	"fmt"

	"github.com/carlmjohnson/errutil"
)

type closer struct{}

func (c closer) Close() error {
	return fmt.Errorf("<had problem closing!>")
}

func openThingie() (c closer, err error) { return }

// Calling Close() on an io.WriteCloser can return an important error
// encountered while flushing to disk. Don't risk missing them by
// using a plain defer w.Close(). Use errutil.Defer to capture the return value.
func ExampleDefer() {
	// If you just defer a close call, you can miss an error
	return1 := func() error {
		thing, err := openThingie()
		if err != nil {
			return err
		}
		defer thing.Close() // oh no, this returned an error!
		// do stuff...
		return nil
	}()
	fmt.Println(return1) // == <nil>

	// Use errutil.Defer and a named return to capture the error
	return2 := func() (err error) {
		thing, err := openThingie()
		if err != nil {
			return err
		}
		defer errutil.Defer(&err, thing.Close)
		// do stuff...
		return nil
	}()
	fmt.Println(return2) // == <had problem closing!>

	// Output:
	// <nil>
	// <had problem closing!>
}
