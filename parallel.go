package errutil

import (
	"sync"
)

// ExecParallel runs the functions in separate goroutines
// and then merges the returned errors, if any. Any panics
// in goroutines will be caught and converted into errors.
func ExecParallel(fs ...func() error) error {
	var (
		size = len(fs)
		wg   sync.WaitGroup
		errs = make(Slice, size)
	)
	wg.Add(size)
	for i := range fs {
		go func(i int) {
			defer wg.Done()
			defer Recover(&errs[i])
			errs[i] = fs[i]()
		}(i)
	}
	wg.Wait()
	return errs.Merge()
}
