module coolsim

go 1.18

require (
	github.com/google/uuid v1.3.0
	github.com/rppala3/broadcast-channel v1.0.0
	go.uber.org/zap v1.21.0
)

require (
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.8.0 // indirect
)

replace github.com/rppala3/broadcast-channel v1.0.0 => github.com/rppala3/broadcast-go v1.0.0
