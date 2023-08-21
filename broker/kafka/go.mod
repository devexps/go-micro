module github.com/devexps/go-micro/broker/kafka/v2

go 1.18

replace github.com/devexps/go-micro/v2 => ../../

require (
	github.com/devexps/go-micro/v2 v2.0.0-00010101000000-000000000000
	github.com/google/uuid v1.3.0
	github.com/segmentio/kafka-go v0.4.42
	go.opentelemetry.io/otel v1.7.0
	go.opentelemetry.io/otel/trace v1.7.0
)

require (
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/klauspost/compress v1.15.9 // indirect
	github.com/pierrec/lz4/v4 v4.1.15 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
)
