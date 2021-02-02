# Project variables
PROJECT_NAME := plugin-sdk
PROJECT_REPO := github.com/terraform-docs/$(PROJECT_NAME)
DESCRIPTION  := Plugin SDK for building custom formatter for terraform-docs
LICENSE      := Apache-2.0

# Build variables
COVERAGE_OUT := coverage.out

# Go variables
GO_PACKAGE_FILES    := $(shell go list ./... | grep -v vendor/ | sed 's|$(PROJECT_REPO)/||')
GOIMPORTS_LOCAL_ARG := -local $(PROJECT_REPO)

# Binary versions
GOLANGCI_VERSION := v1.23.7

.PHONY: all
all: clean verify checkfmt lint test

.PHONY: clean
clean: ## Clean workspace
	@ $(MAKE) --no-print-directory log-$@
	rm -rf ./$(COVERAGE_OUT)

#########################
## Development targets ##
#########################
.PHONY: checkfmt
checkfmt: ## Check formatting of all go files
	@ $(MAKE) --no-print-directory log-$@
	@ if [ $$(goimports -l $(GOIMPORTS_LOCAL_ARG) $(GO_PACKAGE_FILES) | tee /dev/tty | grep -c ".go") = 0 ]; then \
		echo "OK" ; \
	else \
		exit 1 ; \
	fi

.PHONY: fmt
fmt: ## Format all go files
	@ $(MAKE) --no-print-directory log-$@
	goimports -w $(GOIMPORTS_LOCAL_ARG) $(GO_PACKAGE_FILES)

.PHONY: lint
lint: ## Run linter
	@ $(MAKE) --no-print-directory log-$@
	golangci-lint run ./...

.PHONY: staticcheck
staticcheck: ## Run staticcheck
	@ $(MAKE) --no-print-directory log-$@
	go run honnef.co/go/tools/cmd/staticcheck -- ./...

.PHONY: test
test: ## Run tests
	@ $(MAKE) --no-print-directory log-$@
	go test -coverprofile=$(COVERAGE_OUT) -covermode=atomic -v ./...

.PHONY: tidy
tidy: ## Tidy 'vendor' dependencies
	@ $(MAKE) --no-print-directory log-$@
	go mod tidy

.PHONY: verify
verify: ## Verify 'vendor' dependencies
	@ $(MAKE) --no-print-directory log-$@
	go mod verify

.PHONY: vendor
vendor: ## Download 'vendor' dependencies
	@ $(MAKE) --no-print-directory log-$@
	go mod vendor

####################
## Helper targets ##
####################
.PHONY: goimports
goimports:
ifeq (, $(shell which goimports))
	GO111MODULE=off go get -u golang.org/x/tools/cmd/goimports
endif

.PHONY: golangci
golangci:
ifeq (, $(shell which golangci-lint))
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s  -- -b $(shell go env GOPATH)/bin $(GOLANGCI_VERSION)
endif

.PHONY: tools
tools: ## Install required tools
	@ $(MAKE) --no-print-directory log-$@
	@ $(MAKE) --no-print-directory goimports golangci

########################################################################
## Self-Documenting Makefile Help                                     ##
## https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html ##
########################################################################
.PHONY: help
help:
	@ grep -h -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

log-%:
	@ grep -h -E '^$*:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m==> %s\033[0m\n", $$2}'
