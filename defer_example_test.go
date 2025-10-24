package errorx_test

import (
	"fmt"

	"github.com/earthboundkid/errorx/v2"
)

type closer struct{}

func (c closer) Close() error {
	return fmt.Errorf("<had problem closing!>")
}

func openThingie() (c closer, err error) { return }

// Calling Close() on an io.WriteCloser can return an important error
// encountered while flushing to disk. Don't risk missing them by
// using a plain defer w.Close(). Use errorx.Defer to capture the return value.
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

	// Use errorx.Defer and a named return to capture the error
	return2 := func() (err error) {
		thing, err := openThingie()
		if err != nil {
			return err
		}
		defer errorx.Defer(&err, thing.Close)
		// do stuff...
		return nil
	}()
	fmt.Println(return2) // == <had problem closing!>

	// Output:
	// <nil>
	// <had problem closing!>
}
