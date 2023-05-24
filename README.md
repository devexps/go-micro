## About GoMicro

GoMicro is a microservice-oriented governance framework implemented by golang, which offers convenient capabilities to help you quickly build a bulletproof application from scratch, such as:

- The communication protocol is based on the HTTP/gRPC through the definition of Protobuf.
- Abstract transport layer support: HTTP / gRPC.
- Powerful middleware design, support: Tracing (OpenTelemetry), Metrics (Prometheus is default), Recovery and more.
- Registry interface able to be connected with various other centralized registries through plug-ins.
- The standard log interfaces ease the integration of the third-party log libs with logs collected through the *Fluentd*.
- Automatically support the selection of the content encoding with Accept and Content-Type.
- Multiple data sources are supported for configurations and dynamic configurations (use atomic operations).
- In the protocol of HTTP/gRPC, use the uniform metadata transfer method.
- You can define errors in protos and generate enums with protoc-gen-go.
- You can define verification rules in Protobuf supported by the HTTP/gRPC service.

GoMicro is accessible, powerful, and provides tools required for large, robust applications.

### Principles

* **Simple**: Appropriate design with plain and easy code.
* **General**: Cover the various utilities for business development.
* **Highly efficient**: Speeding up the efficiency of businesses upgrading.
* **Stable**: The base libs validated in the production environment have the characteristics of high testability, high coverage as well as high security and reliability.
* **Robust**: Eliminating misusing through high quality of the base libs.
* **High-performance**: Optimal performance excluding the optimization of hacking in case of *unsafe*. 
* **Expandability**: Properly designed interfaces where you can expand utilities such as base libs to meet your further requirements.
* **Fault-tolerance**: Designed against failure, enhance the understanding and exercising of SRE within GoMicro to achieve more robustness.
* **Toolchain**: Includes an extensive toolchain, such as the code generation of cache, the lint tool, and so forth.

## License

The GoMicro framework is open-sourced software licensed under the [MIT license](./LICENSE).