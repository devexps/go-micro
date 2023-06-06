## About GoMicro

Go-Micro is a set of lightweight Go microservice framework, including a large number of microservice related frameworks and tools, such as:

- The communication protocol is based on the HTTP/gRPC through the definition of Protobuf.
- Abstract transport layer support: HTTP / gRPC.
- Powerful middleware design, support: Tracing (OpenTelemetry), Metrics (Prometheus is default), Recovery and more.
- Registry interface able to be connected with various other centralized registries through plug-ins.
- The standard log interfaces ease the integration of the third-party log libs.
- Automatically support the selection of the content encoding with Accept and Content-Type.
- Multiple data sources are supported for configurations and dynamic configurations (use atomic operations).
- In the protocol of HTTP/gRPC, use the uniform metadata transfer method.
- You can define errors in protos and generate enums with protoc-gen-go.
- You can define verification rules in Protobuf supported by the HTTP/gRPC service.

GoMicro is accessible, powerful, and provides tools required for large, robust applications.

## License

The GoMicro framework is open-sourced software licensed under the [MIT license](./LICENSE).