module github.com/devexps/go-micro/log/logrus/v2

go 1.18

replace github.com/devexps/go-micro/v2 => ../../

require (
	github.com/devexps/go-micro/v2 v2.0.0-20230823132135-27ba0739d0d2
	github.com/sirupsen/logrus v1.8.1
)

require golang.org/x/sys v0.10.0 // indirect
