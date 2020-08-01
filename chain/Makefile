#VERSION := $(shell echo $(shell git describe --tags) | sed 's/^v//')
#COMMIT := $(shell git log -1 --format='%H')
VERSION := 0.1.0
COMMIT := ''
CHAIN_NAME := BC
CHAIN_NAME_LO := $(shell echo $(CHAIN_NAME) | tr '[A-Z]' '[a-z]')

ldflags = -X github.com/hashrs/blockchain/framework/chain-app/version.Name='$(CHAIN_NAME)Chain' \
	-X github.com/hashrs/blockchain/framework/chain-app/version.ServerName=node \
	-X github.com/hashrs/blockchain/framework/chain-app/version.ClientName=cli \
	-X github.com/hashrs/blockchain/framework/chain-app/version.Version=$(VERSION) \
	-X github.com/hashrs/blockchain/framework/chain-app/version.Commit=$(COMMIT) \
	-X github.com/hashrs/blockchain/ChainName=$(CHAIN_NAME)
	
BUILD_FLAGS := -ldflags '$(ldflags)'

all: install

TARGET_PATH =  $(shell pwd)/../../../pkg
TARGET_OS = darwin linux windows

define INSTALL_GO_LIB
	GO111MODULE=off CGO_ENABLED=0 GOOS=$(1) GOARCH=amd64 go build -o ../bin/$(CHAIN_NAME_LO)-node-$(1) -pkgdir $(TARGET_PATH)/$(1)_amd64 ./cmd/node
	GO111MODULE=off CGO_ENABLED=0 GOOS=$(1) GOARCH=amd64 go build -o ../bin/$(CHAIN_NAME_LO)-cli-$(1) -pkgdir $(TARGET_PATH)/$(1)_amd64 ./cmd/cli

endef	

build_a:
	#go mod vendor
	$(foreach e, $(TARGET_OS), $(call INSTALL_GO_LIB,$(e)))

install: go.sum
	go install -mod=readonly $(BUILD_FLAGS) ./cmd/node
	go install -mod=readonly $(BUILD_FLAGS) ./cmd/cli

build: go.sum
	go build $(BUILD_FLAGS) -o ../bin/$(CHAIN_NAME_LO)-node ./cmd/node
	go build $(BUILD_FLAGS) -o ../bin/$(CHAIN_NAME_LO)-cli ./cmd/cli

go.sum: go.mod
		@echo "--> Ensure dependencies have not been modified"
		go mod verify
