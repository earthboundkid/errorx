module github.com/carlmjohnson/errutil/compat-test

go 1.16

require (
	github.com/carlmjohnson/errutil v0.0.0
	github.com/hashicorp/go-multierror v1.1.1
	go.uber.org/multierr v1.7.0
)

replace github.com/carlmjohnson/errutil => ../
