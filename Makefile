PROJECT=spa

DEFAULT_OUT=bin/$(PROJECT)
ifdef GOOS
DEFAULT_OUT:=$(DEFAULT_OUT).$(GOOS)
endif
ifdef GOARCH
DEFAULT_OUT:=$(DEFAULT_OUT).$(GOARCH)
endif
ifeq ($(GOOS),windows)
DEFAULT_OUT:=$(DEFAULT_OUT).exe
endif
OUT?=$(DEFAULT_OUT)

.PHONY: build
build:
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $(OUT) .

.PHONY: build-linux
build-linux:
	$(MAKE) build GOOS=linux GOARCH=amd64

.PHONY: lint
lint:
	golangci-lint run
