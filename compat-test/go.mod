module github.com/carlmjohnson/errutil/compat-test

go 1.18

require (
	github.com/carlmjohnson/be v0.22.4
	github.com/carlmjohnson/errutil v0.0.0
	github.com/hashicorp/go-multierror v1.1.1
	go.uber.org/multierr v1.8.0
)

require (
	github.com/hashicorp/errwrap v1.1.0 // indirect
	go.uber.org/atomic v1.9.0 // indirect
)

replace github.com/carlmjohnson/errutil => ../
