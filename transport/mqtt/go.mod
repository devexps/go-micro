module github.com/devexps/go-micro/transport/mqtt/v2

go 1.19

replace github.com/devexps/go-micro/broker/mqtt/v2 => ../../broker/mqtt/

replace github.com/devexps/go-micro/v2 => ../../

require (
	github.com/devexps/go-micro/broker/mqtt/v2 v2.0.0-20240105082415-89bf106e3d8b
	github.com/devexps/go-micro/v2 v2.0.7
)

require (
	github.com/devexps/go-pkg/v2 v2.0.2 // indirect
	github.com/eclipse/paho.mqtt.golang v1.4.3 // indirect
	github.com/go-logr/logr v1.3.0 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-playground/form/v4 v4.2.1 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/uuid v1.4.0 // indirect
	github.com/gorilla/mux v1.8.1 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	go.opentelemetry.io/otel v1.21.0 // indirect
	go.opentelemetry.io/otel/metric v1.21.0 // indirect
	go.opentelemetry.io/otel/trace v1.21.0 // indirect
	golang.org/x/net v0.19.0 // indirect
	golang.org/x/sync v0.5.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20231127180814-3a041ad873d4 // indirect
	google.golang.org/grpc v1.59.0 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
