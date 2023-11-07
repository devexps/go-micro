module github.com/devexps/go-micro/log/logrus/v2

go 1.18

replace github.com/devexps/go-micro/v2 => ../../

require (
	github.com/devexps/go-micro/v2 v2.0.2
	github.com/sirupsen/logrus v1.8.1
)

require golang.org/x/sys v0.10.0 // indirect
