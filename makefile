# ==================================================================================== #
# HELPERS
# ==================================================================================== #

.PHONY: help
## help: print this help message
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'


.PHONY: confirm
confirm:
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]

.PHONY: no-dirty
no-dirty:
	git diff --exit-code


# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #

.PHONY: devops/up
## devops/up: Scaffold the development environment
devops/up:
	podman kube play devops/tracing.yaml
	podman kube play devops/redpanda.yaml

.PHONY: devops/down
## devops/down: Clean the development environment
devops/down:
	podman kube play --down devops/tracing.yaml
	podman kube play --down devops/redpanda.yaml

# ==================================================================================== #
# Quality Control
# ==================================================================================== #
.PHONY: lint
## lint: Run quality control checks
lint:
	go vet ./...
	go fmt ./...
	golangci-lint run

# ==================================================================================== #
# Application
# ==================================================================================== #

.PHONY: app/run
## app/run: Run the application
app/serve:
	set OTEL_EXPORTER_OTLP_INSECURE=true
	set OTEL_EXPORTER_OTLP_ENDPOINT=http://localhost:4317
	go run ./cmd --config config.yaml serve

.PHONY: app/setup
## app/setup: Setup the application
app/setup:
	set OTEL_EXPORTER_OTLP_INSECURE=true
	go run ./cmd --config config.yaml setup


# ==================================================================================== #
# Build
# ==================================================================================== #

.PHONY: build
## build: Build the application
build:
	CGO_ENABLED=0 go build -trimpath -o ./bin/server_$(GOOS)_$(GOARCH)$(EXTENSION) ./cmd
.PHONY: build/linux_amd64
build/linux_amd64:
	GOOS=linux GOARCH=amd64 $(MAKE) build
.PHONY: build/linux_arm64
build/linux_arm64:
	GOOS=linux GOARCH=arm64 $(MAKE) build
.PHONY: build/darwin_amd64
build/darwin_amd64:
	GOOS=darwin GOARCH=amd64 $(MAKE) build
.PHONY: build/darwin_arm64
build/darwin_arm64:
	GOOS=darwin GOARCH=arm64 $(MAKE) build