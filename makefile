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

.PHONY: devops/scaffold
## devops/scaffold: Scaffold the development environment
devops/scaffold:
	podman kube play devops/redpanda.yml


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
.PHONY: lint
## lint: Run quality control checks
lint:
	go vet ./...
	go fmt ./...
	golangci-lint run

# ==================================================================================== #
# Application
# ==================================================================================== #
.PHONY: app/build
## app/build: Build the application
app/build:
	go build -o bin/app .

.PHONY: app/run
## app/run: Run the application
app/run:
	set OTEL_EXPORTER_OTLP_INSECURE=true
	go run . serve -c config.yml