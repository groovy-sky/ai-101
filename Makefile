#/bin/bash

BIN_DIR := tools/bin
SRC_DIR := tools/src

PLATFORMS := \
	linux/amd64 \
	linux/arm64 \
	darwin/amd64 \
	darwin/arm64 \
	windows/amd64

TOOLS := $(notdir $(wildcard $(SRC_DIR)/*))

.PHONY: build clean

build:
	@mkdir -p $(BIN_DIR)
	@for tool in $(TOOLS); do \
		for platform in $(PLATFORMS); do \
			GOOS=$${platform%/*}; \
			GOARCH=$${platform#*/}; \
			EXT=""; \
			if [ "$$GOOS" = "windows" ]; then EXT=".exe"; fi; \
			OUT="$(BIN_DIR)/$$tool-$$GOOS-$$GOARCH$$EXT"; \
			echo "Building $$OUT"; \
			(cd $(SRC_DIR)/$$tool && GOOS=$$GOOS GOARCH=$$GOARCH CGO_ENABLED=0 go build -o "$$(pwd)/../../bin/$$tool-$$GOOS-$$GOARCH$$EXT" .); \
		done; \
	done

clean:
	rm -rf $(BIN_DIR)