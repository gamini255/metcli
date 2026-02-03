.PHONY: fmt test build metcli run start

METCLI_ARGS ?=
BINDIR ?= bin
METCLI_BIN := $(BINDIR)/metcli
METCLI_DEPS := $(shell git ls-files '*.go' go.mod go.sum)

# Allow: `make metcli instagram home` (extra make goals become args).
ifneq (,$(filter $(firstword $(MAKECMDGOALS)),metcli run start))
  ifeq (,$(METCLI_ARGS))
    METCLI_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
  endif
  $(eval $(METCLI_ARGS):;@:)
endif

fmt:
	gofmt -w .

test:
	go test ./...

build:
	go build ./...

$(METCLI_BIN): $(METCLI_DEPS)
	@mkdir -p $(BINDIR)
	go build -o $(METCLI_BIN) ./cmd/metcli

metcli run start:
	@mkdir -p $(BINDIR)
	go build -o $(METCLI_BIN) ./cmd/metcli
	$(METCLI_BIN) $(filter-out --,$(METCLI_ARGS))
