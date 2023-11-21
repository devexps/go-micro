module github.com/devexps/go-micro/transport/kafka/v2

go 1.19

replace (
	github.com/devexps/go-micro/broker/kafka/v2 => ../../broker/kafka/
	github.com/devexps/go-micro/v2 => ../../
)

require (
	github.com/devexps/go-micro/broker/kafka/v2 v2.0.5
	github.com/devexps/go-micro/v2 v2.0.5
	github.com/segmentio/kafka-go v0.4.45
	go.opentelemetry.io/otel v1.21.0
	go.opentelemetry.io/otel/trace v1.21.0
)

require (
	github.com/go-logr/logr v1.3.0 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-playground/form/v4 v4.2.1 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/uuid v1.4.0 // indirect
	github.com/klauspost/compress v1.17.3 // indirect
	github.com/kr/pretty v0.1.0 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/pierrec/lz4/v4 v4.1.18 // indirect
	go.opentelemetry.io/otel/metric v1.21.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20231120223509-83a465c0220f // indirect
	google.golang.org/grpc v1.59.0 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
