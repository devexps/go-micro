user	:=	$(shell whoami)

# GOBIN > GOPATH > INSTALLDIR
# Mac OS X
ifeq ($(shell uname),Darwin)
GOBIN	:=	$(shell echo ${GOBIN} | cut -d':' -f1)
GOPATH	:=	$(shell echo $(GOPATH) | cut -d':' -f1)
endif
BIN		:= 	""
# check GOBIN
ifneq ($(GOBIN),)
	BIN=$(GOBIN)
else
# check GOPATH
	ifneq ($(GOPATH),)
		BIN=$(GOPATH)/bin
	endif
endif

all:
	@cd cmd/micro && go build && cd - &> /dev/null
	@cd cmd/protoc-gen-go-http && go build && cd - &> /dev/null
	@cd cmd/protoc-gen-go-errors && go build && cd - &> /dev/null

.PHONY: install
install: all
ifeq ($(user),root)
#root, install for all user
	@cp ./cmd/micro/micro /usr/bin
	@cp ./cmd/protoc-gen-go-errors/protoc-gen-go-errors /usr/bin
	@cp ./cmd/protoc-gen-go-http/protoc-gen-go-http /usr/bin
else
#!root, install for current user
	$(shell cp ./cmd/micro/micro '$(BIN)';cp ./cmd/protoc-gen-go-errors/protoc-gen-go-errors '$(BIN)';cp ./cmd/protoc-gen-go-http/protoc-gen-go-http '$(BIN)';)
endif
	@which protoc-gen-go &> /dev/null || go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.30.0
	@which protoc-gen-go-grpc &> /dev/null || go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0
	@which protoc-gen-validate  &> /dev/null || go install github.com/envoyproxy/protoc-gen-validate@v1.0.0
	@echo "install finished"

.PHONY: uninstall
uninstall:
	$(shell for i in `which -a micro | grep -v '/usr/bin/micro' 2>/dev/null | sort | uniq`; do read -p "Press to remove $${i} (y/n): " REPLY; if [ $${REPLY} = "y" ]; then rm -f $${i}; fi; done)
	$(shell for i in `which -a protoc-gen-go-http | grep -v '/usr/bin/protoc-gen-go-http' 2>/dev/null | sort | uniq`; do read -p "Press to remove $${i} (y/n): " REPLY; if [ $${REPLY} = "y" ]; then rm -f $${i}; fi; done)
	$(shell for i in `which -a protoc-gen-go-errors | grep -v '/usr/bin/protoc-gen-go-errors' 2>/dev/null | sort | uniq`; do read -p "Press to remove $${i} (y/n): " REPLY; if [ $${REPLY} = "y" ]; then rm -f $${i}; fi; done)
	$(shell for i in `which -a protoc-gen-go | grep -v '/usr/bin/protoc-gen-go' 2>/dev/null | sort | uniq`; do read -p "Press to remove $${i} (y/n): " REPLY; if [ $${REPLY} = "y" ]; then rm -f $${i}; fi; done)
	$(shell for i in `which -a protoc-gen-go-grpc | grep -v '/usr/bin/protoc-gen-go-grpc' 2>/dev/null | sort | uniq`; do read -p "Press to remove $${i} (y/n): " REPLY; if [ $${REPLY} = "y" ]; then rm -f $${i}; fi; done)
	$(shell for i in `which -a protoc-gen-validate | grep -v '/usr/bin/protoc-gen-validate' 2>/dev/null | sort | uniq`; do read -p "Press to remove $${i} (y/n): " REPLY; if [ $${REPLY} = "y" ]; then rm -f $${i}; fi; done)
	@echo "uninstall finished"

.PHONY: proto
proto:
	@protoc --proto_path=./third_party --proto_path=. --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. --go-http_out=paths=source_relative:. api/metadata/metadata.proto
	@protoc --proto_path=. --go_out=paths=source_relative:. cmd/protoc-gen-go-errors/errors/errors.proto
	@protoc --proto_path=. --go_out=paths=source_relative:. errors/errors.proto
	@echo "generate proto finished"

.PHONY: test
test:
	@go test -v -race ./...
	@echo "go test finished"