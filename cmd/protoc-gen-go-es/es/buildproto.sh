protoc -I/usr/local/include -I. \
  -I$GOPATH/src \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway \
  --go_out=. --go-grpc_opt=paths=source_relative --go_opt=paths=source_relative --go-grpc_out=. --grpc-gateway_out=logtostderr=true:. \
  *.proto
