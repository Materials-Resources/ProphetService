## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'
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

.PHONY: dev/validate
## dev/validate: Validate the application
dev/validate:
	go fmt ./...
	go vet ./...


.PHONY: proto/breaking
## proto/breaking: Check for breaking changes in the proto files
proto/breaking:
	buf breaking proto --against 'https://github.com/Materials-Resources/s_prophet.git#branch=main,ref=HEAD~0'

.PHONY: proto/validate
## proto/validate: Validate the proto files
proto/validate: proto/breaking
	buf lint proto

.PHONY: proto/check
## proto/check: Check the proto files
proto/check: proto/validate proto/breaking

.PHONY: proto/generate
## proto/generate: Generate packages from proto files using the buf cli tool
proto/generate: proto/check
	buf generate proto


.PHONY: dev/app/run
## dev/app/run: Run the application
dev/app/run:
	set OTEL_EXPORTER_OTLP_INSECURE=true
	go run . serve -c config.yml