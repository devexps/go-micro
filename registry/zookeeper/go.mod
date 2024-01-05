module github.com/devexps/go-micro/registry/zookeeper/v2

go 1.19

replace github.com/devexps/go-micro/v2 => ../../

require (
	github.com/devexps/go-micro/v2 v2.0.7
	github.com/go-zookeeper/zk v1.0.3
	golang.org/x/sync v0.5.0
)
