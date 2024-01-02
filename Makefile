build-binary:
	@echo "building binary..." && \
	cd server && \
	rm -rf ./bin && \
	mkdir -p ./bin && \
	go build -o ./bin/sampler-go-grpc-http-gateway-server ./cmd/main.go

gen-protos:
	@echo "running generator for proto files..." && \
	chmod +x ./scripts/gen-protos.sh && ./scripts/gen-protos.sh

install-gateway-deps:
	cd server && \
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
        github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
        google.golang.org/protobuf/cmd/protoc-gen-go \
        google.golang.org/grpc/cmd/protoc-gen-go-grpc

build-docker:
	@echo "building docker image..." && \
	cd server && \
	docker buildx build --platform linux/amd64 -t eganowdevteam/sampler-go-grpc-http-gateway-server:latest .

.PHONY: gen-protos install-gateway-deps build-binary build-docker
