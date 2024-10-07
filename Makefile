all:
	go build

.PHONY: spdx
spdx:
	./tools/spdx-ensure

.PHONY: lint
lint:
	golangci-lint run
