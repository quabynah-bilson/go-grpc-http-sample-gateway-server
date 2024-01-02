gen-protos:
	@echo "running generator for proto files..." && \
	chmod +x ./scripts/gen-protos.sh && ./scripts/gen-protos.sh

install-gateway-deps:
	cd server && \
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
        github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
        google.golang.org/protobuf/cmd/protoc-gen-go \
        google.golang.org/grpc/cmd/protoc-gen-go-grpc

.PHONY: gen-protos install-gateway-deps
