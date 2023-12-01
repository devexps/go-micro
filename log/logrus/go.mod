module github.com/devexps/go-micro/log/logrus/v2

go 1.19

replace github.com/devexps/go-micro/v2 => ../../

require (
	github.com/devexps/go-micro/v2 v2.0.6
	github.com/sirupsen/logrus v1.9.3
)

require golang.org/x/sys v0.15.0 // indirect
