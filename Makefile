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
	docker buildx build --platform linux/amd64 -f server/devops/docker/Dockerfile -t eganowdevteam/sampler-go-grpc-http-gateway-server:latest .

start-sql-server:
	@echo "starting sql server..." && \
	docker run -d --name mssql-server -e 'ACCEPT_EULA=Y' -e 'SA_PASSWORD=Password123' -p 1433:1433 mcr.microsoft.com/mssql/server:2022-latest

start-docker-services:
	@echo "starting docker services..." && \
	docker-compose -f server/devops/docker/compose.yaml up -d

apply-k8s:
	@echo "applying k8s resources..." && \
	kubectl apply -f server/devops/k8s -R

.PHONY: gen-protos install-gateway-deps build-binary build-docker start-sql-server start-docker-services apply-k8s
