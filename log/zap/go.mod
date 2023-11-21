module github.com/devexps/go-micro/log/zap/v2

go 1.19

replace github.com/devexps/go-micro/v2 => ../../

require (
	github.com/devexps/go-micro/v2 v2.0.5
	go.uber.org/zap v1.26.0
)

require go.uber.org/multierr v1.11.0 // indirect
