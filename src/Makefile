BUILD_ENVPARAMS:=CGO_ENABLED=0

LOCAL_BIN:=$(CURDIR)/bin

ifndef BIN_DIR
BIN_DIR=./bin
endif

.PHONY: .run
.run:
	go run ./cmd/main.go

.PHONY: run
run: .run


.PHONY: .build
.build:
	$(BUILD_ENVPARAMS) go build -o "$(BIN_DIR)" ./cmd/main.go

PHONY: build
build: .build

PHONY: .bin-deps
.bin-deps:
	GOBIN=$(LOCAL_BIN) go install github.com/planetscale/vtprotobuf/cmd/protoc-gen-go-vtproto@v0.3.0 && \
  	GOBIN=$(LOCAL_BIN) go install github.com/mitchellh/gox@v1.0.1  && \
  	GOBIN=$(LOCAL_BIN) go install golang.org/x/tools/cmd/goimports@v0.1.9

PHONY: bin-deps
bin-deps: .bin-deps

#TODO: migration