server_root := ./server

# Makefile target for creating feature structure (run directly in your terminal)
# e.g. make feature ...feature_names (space separated e.g. make feature auth users)
feature:
	@$(foreach feature,$(filter-out $@,$(MAKECMDGOALS)), $(MAKE) create_feature_structure FEATURE_NAME=$(feature);)

# Function to create feature structure for each feature
create_feature_structure:
	@$(eval feature_root=$(server_root)/features/$(FEATURE_NAME))
	@echo "Creating feature structure in $(feature_root)"

	# Define directories to be created
	@$(eval directories := \
		"$(feature_root)/business_logic/app/repository/models" \
		"$(feature_root)/business_logic/services" \
		"$(feature_root)/di" \
		"$(feature_root)/pkg" \
	)

	# Create directories
	@$(foreach dir,$(directories),mkdir -p "$(dir)";)

	# Define file paths and contents
	@$(eval files := \
		"$(feature_root)/business_logic/app/use_case.go|package app" \
		"$(feature_root)/business_logic/app/data_source.go|package app" \
		"$(feature_root)/business_logic/app/repository/noop.go|package repository" \
		"$(feature_root)/business_logic/services/grpc.go|package service" \
		"$(feature_root)/di/injector.go|package di" \
		"$(feature_root)/pkg/repository.go|package pkg" \
		"$(feature_root)/pkg/data_source.go|package pkg" \
	)

	# Create files with content using a shell loop
	@for file in $(files); do \
		path=$$(echo "$$file" | cut -d'|' -f1); \
		content=$$(echo "$$file" | cut -d'|' -f2); \
		echo "Creating file $$path with content: $$content"; \
		echo "$$content" > "$$path"; \
	done

# Trick to allow passing arguments to a target
%:
	@:

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

.PHONY: gen-protos install-gateway-deps build-binary build-docker start-sql-server start-docker-services apply-k8s feature
