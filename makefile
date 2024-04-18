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
	podman kube down deployment/dev/redpanda.yml

.PHONY: dev/metrics/up
## dev/metrics/up: Start metrics development server
dev/metrics/up:
	podman kube play deployment/dev/metrics.yml

.PHONY: dev/metrics/down
## dev/metrics/up: Stop metrics development server
dev/metrics/down:
	podman kube down deployment/dev/metrics.yml

.PHONY: dev/generate/proto
## dev/generate/proto: Generate packages from proto files using the buf cli tool
dev/generate/proto:
	buf generate proto

.PHONY: dev/app/run
## dev/app/run: Run the application
dev/app/run:
	set OTEL_EXPORTER_OTLP_INSECURE=true
	go run . serve -c config.yml