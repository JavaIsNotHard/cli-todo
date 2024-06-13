MAIN_PACKAGE_PATH := ./cmd/app
BINARY_NAME := todo

.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

.PHONY: build
build:
	@echo "Building the application..."
	go build -o /tmp/bin/${BINARY_NAME} ${MAIN_PACKAGE_PATH}

.PHONY: run
run: build
	@echo "Running the application..."
	/tmp/bin/${BINARY_NAME}
