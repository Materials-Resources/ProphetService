## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'


# ==================================================================================== #
# Development
# ==================================================================================== #

.PHONY: dev/setup
## dev/setup: Setup the development environment
dev/setup: dev/tools dev/redpanda/up dev/metrics/up
	@echo "Install linter"
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.57.2

.PHONY: dev/clean
## dev/clean: Clean the development environment
dev/clean: dev/redpanda/down dev/metrics/down

.PHONY: dev/tools
## dev/tools: Check and Install the development tools
dev/tools:

	@echo "Checking for buf cli"
	@which buf || $(MAKE) dev/tools/install/buf

	@echo "Checking for podman"
	@which podman || $(MAKE) dev/tools/install/podman

.PHONY: dev/tools/install/podman
dev/tools/install/podman:
	@echo "Installing podman"
	 brew install podman

.PHONY: dev/tools/install/buf
dev/tools/install/buf:
	@echo "Installing buf cli"
	 brew install bufbuild/buf/buf

.PHONY: dev/redpanda/up
## dev/redpanda/up: Start redpanda development cluster
dev/redpanda/up:
	podman kube play deployment/dev/redpanda.yml

.PHONY: dev/redpanda/down
## dev/redpanda/down: Stop redpanda development cluster
dev/redpanda/down:
	podman play kube --down deployment/dev/redpanda.yml

.PHONY: dev/metrics/up
## dev/metrics/up: Start metrics development server
dev/metrics/up:
	podman play kube deployment/dev/metrics.yml

.PHONY: dev/metrics/down
## dev/metrics/up: Stop metrics development server
dev/metrics/down:
	podman play kube --down deployment/dev/metrics.yml




# ==================================================================================== #
# Quality Control
# ==================================================================================== #
.PHONY: qc
## qc: Run all quality control checks
qc: qc/proto qc/app

.PHONY: qc/proto
qc/proto:
	buf breaking proto --against 'https://github.com/Materials-Resources/s_prophet.git#branch=main,subdir=proto'
	buf lint proto

.PHONY: qc/app
## qc/app: Validate the application
qc/app:
	go fmt ./...
	go vet ./...
	golangci-lint run ./...


# ==================================================================================== #
# Code Generation
# ==================================================================================== #
.PHONY: gen/proto
## gen/proto: Generate packages from proto files using the buf cli tool
gen/proto: qc/proto
	buf generate proto


# ==================================================================================== #
# Application
# ==================================================================================== #
.PHONY: app/build
## app/build: Build the application
app/build: qc
	go build -o bin/s_prophet .

.PHONY: app/serve
## app/serve: Run the application
app/serve: qc
	set OTEL_EXPORTER_OTLP_INSECURE=true
	go run . serve -c config.yml