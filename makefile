dev/redpanda/up:
	podman kube play development/redpanda.yml
dev/redpanda/down:
	podman kube down development/redpanda.yml