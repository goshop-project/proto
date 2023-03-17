.PHONY: all clean generate fmt tidy
.PHONY: FORCE

GO ?= go
GOFMT ?= gofmt
GOFMT_FLAGS = -w -l -s
GOGENERATE_FLAGS = -v

GOPATH ?= $(shell $(GO) env GOPATH)
GOBIN ?= $(GOPATH)/bin

TOOLSDIR := $(CURDIR)/tools
TMPDIR ?= $(CURDIR)/.tmp

BUF_INSTALL_URL ?= github.com/bufbuild/buf/cmd/buf
GNOSTIC_INSTALL_URL ?= github.com/google/gnostic
GNOSTIC_GRPC_INSTALL_URL ?= github.com/googleapis/gnostic-grpc
PROTOC_GEN_OPENAPIv2 ?= github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
PROTOC_GEN_GO_GRPC_INSTALL_URL ?= google.golang.org/grpc/cmd/protoc-gen-go-grpc
PROTOC_GEN_GO_INSTALL_URL ?= google.golang.org/protobuf/cmd/protoc-gen-go
PROTOC_GEN_GRPC_GATEWAY_INSTALL_URL ?= github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
REVIVE_INSTALL_URL ?= github.com/mgechev/revive

BUF ?= buf

REVIVE ?= $(GOBIN)/revive
REVIVE_CONF ?= $(TOOLSDIR)/revive.toml
REVIVE_RUN_ARGS ?= -config $(REVIVE_CONF) -formatter friendly

GO_INSTALL_URLS = \
	$(BUF_INSTALL_URL) \
	$(GNOSTIC_INSTALL_URL) \
	$(GNOSTIC_GRPC_INSTALL_URL) \
	$(PROTOC_GEN_OPENAPIv2) \
	$(PROTOC_GEN_GO_GRPC_INSTALL_URL) \
	$(PROTOC_GEN_GO_INSTALL_URL) \
	$(PROTOC_GEN_GRPC_GATEWAY_INSTALL_URL) \
	$(REVIVE_INSTALL_URL) \

V = 0
Q = $(if $(filter 1,$V),,@)
M = $(shell if [ "$$(tput colors 2> /dev/null || echo 0)" -ge 8 ]; then printf "\033[34;1m▶\033[0m"; else printf "▶"; fi)

all: get generate tidy build

clean: ; $(info $(M) cleaning…)
	rm -rf $(TMPDIR)

$(TMPDIR)/index: $(TOOLSDIR)/gen_index.sh Makefile FORCE ; $(info $(M) generating index…)
	$Q mkdir -p $(@D)
	$Q $< > $@~
	$Q if cmp $@ $@~ 2> /dev/null >&2; then rm $@~; else mv $@~ $@; fi

$(TMPDIR)/gen.mk: $(TOOLSDIR)/gen_mk.sh $(TMPDIR)/index Makefile ; $(info $(M) generating subproject rules…)
	$Q mkdir -p $(@D)
	$Q $< $(TMPDIR)/index > $@~
	$Q if cmp $@ $@~ 2> /dev/null >&2; then rm $@~; else mv $@~ $@; fi

include $(TMPDIR)/gen.mk

fmt: ; $(info $(M) reformatting sources…)
	$Q find . -name '*.go' | xargs -r $(GOFMT) $(GOFMT_FLAGS)

tidy: fmt

generate: ; $(info $(M) running go:generate…)
	$Q git grep -l '^//go:generate' | sort -uV | xargs -r -n1 $(GO) generate $(GOGENERATE_FLAGS)

$(REVIVE):
	$Q $(GO) install -v $(REVIVE_INSTALL_URL)
