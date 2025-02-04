branch = main

.PHONY: pb clean deploy

pb:
	mkdir -p pb
	protoc --go_out=pb --go-grpc_out=pb protos/*.proto && \
	protoc --dart_out=grpc:mobile/lib protos/*.proto