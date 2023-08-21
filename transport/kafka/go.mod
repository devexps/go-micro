module github.com/devexps/go-micro/transport/kafka/v2

go 1.18

replace github.com/devexps/go-micro/v2 => ../../

require (
	github.com/devexps/go-micro/broker/kafka/v2 v2.0.0-20230821111625-e154c96f19e4
	github.com/devexps/go-micro/v2 v2.0.0-00010101000000-000000000000
	github.com/segmentio/kafka-go v0.4.42
	go.opentelemetry.io/otel v1.7.0
	go.opentelemetry.io/otel/trace v1.7.0
)

require (
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-playground/form/v4 v4.2.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/klauspost/compress v1.15.9 // indirect
	github.com/kr/pretty v0.1.0 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/pierrec/lz4/v4 v4.1.15 // indirect
	google.golang.org/genproto v0.0.0-20220519153652-3a47de7e79bd // indirect
	google.golang.org/grpc v1.46.2 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
