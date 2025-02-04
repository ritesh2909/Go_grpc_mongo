PROTO_SRC := pb/*.proto
PROTO_OUT := pb

PROTOC_GEN_GO := protoc-gen-go
PROTOC_GEN_GRPC := protoc-gen-go-grpc

generate:
	protoc --go_out=$(PROTO_OUT) --go-grpc_out=$(PROTO_OUT) $(PROTO_SRC)

clean:
	rm -rf $(PROTO_OUT)/user_crud

build:
	go build -o myapp main.go
