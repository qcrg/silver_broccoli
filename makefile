BUILD_DIR ?= build
GEN_DIR ?= gen
PROG_NAME ?= silver_broccoli

default: api/rpc.capnp.go dev

build:
	-@mkdir build
	go build -o ${BUILD_DIR}/${PROG_NAME} cmd/main.go

build_test:
	-@mkdir build
	go build -o ${BUILD_DIR}/${PROG_NAME}_test cmd/test/main.go

run:
	${BUILD_DIR}/${PROG_NAME}

dev:
	go run cmd/main.go

test:
	go run cmd/test/main.go

auto_test:
	go test ./...

clean:
	${RM} -rf ${BUILD_DIR}

api/rpc.capnp.go: schemes/rpc.capnp
	@PATH="${GOPATH}/bin:${PATH}" \
  	capnp compile \
      -I `go list -m -f '{{.Dir}}' capnproto.org/go/capnp/v3`/std \
      --src-prefix=schemes -ogo:api $<

.PHONY: clean build run dev test compile_capnp
