CONTAINER_ENGINE?=$(shell command -v podman 2>/dev/null || echo docker)

.PHONY: build-all
build-all:
	cd components/api-server && $(MAKE) image
	cd components/api-server && $(MAKE) image-controller

.PHONY: lint
lint:
	cd components/api-server && go fmt ./... && go vet ./...
	cd components/control-plane && go fmt ./... && go vet ./...

.PHONY: test-all
test-all:
	cd components/api-server && $(MAKE) test

.PHONY: kind-up
kind-up: build-all
	cd components/api-server && $(MAKE) kind-up

.PHONY: kind-down
kind-down:
	cd components/api-server && $(MAKE) kind-down

.PHONY: kind-rebuild
kind-rebuild:
	cd components/api-server && $(MAKE) kind-rebuild

.PHONY: kind-status
kind-status:
	cd components/api-server && $(MAKE) kind-status
