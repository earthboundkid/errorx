package errutil

import (
	"sync"
)

// ExecParallel runs the functions in separate goroutines
// and then merges the returned errors, if any.
func ExecParallel(fs ...func() error) error {
	var (
		size = len(fs)
		wg   sync.WaitGroup
		errs = make(Slice, size)
	)
	wg.Add(size)
	for i := range fs {
		go func(i int) {
			errs[i] = fs[i]()
			wg.Done()
		}(i)
	}
	wg.Wait()
	return errs.Merge()
}
