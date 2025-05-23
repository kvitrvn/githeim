SHELL := /bin/bash

GIT_VERSION ?= $(shell git describe --tags --abbrev=0 2>/dev/null || echo "0.0.0-dev")
GIT_COMMIT ?= $(shell git describe --abbrev=8 --dirty --always --long)
LDFLAGS ?= -w -s \
	-X 'github.com/kvitrvn/githeim/internal/build.Version=$(GIT_VERSION)' \
	-X 'github.com/kvitrvn/githeim/internal/build.Commit=$(GIT_COMMIT)'

GCFLAGS ?= -trimpath=$(PWD)
ASMFLAGS ?= -trimpath=$(PWD) \

watch: tools/modd/bin/modd
	tools/modd/bin/modd

build:
	CGO_ENABLED=0 \
		go build \
			-ldflags "$(LDFLAGS)" \
			-gcflags "$(GCFLAGS)" \
			-asmflags "$(ASMFLAGS)" \
			-o ./bin/githeim .

tools/modd/bin/modd:
	mkdir -p tools/modd/bin
	GOBIN=$(PWD)/tools/modd/bin go install github.com/cortesi/modd/cmd/modd@latest