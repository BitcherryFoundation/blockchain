GOTOOLS = \
	github.com/golangci/golangci-lint/cmd/golangci-lint
GOTOOLS_CHECK = golangci-lint

all: check_tools test

########################################
###  Build static lib .a

TARGET_PATH =  $(shell pwd)/../../pkg
TARGET_OS = darwin linux windows

define INSTALL_GO_LIB
	GO111MODULE=off CGO_ENABLED=0 GOOS=$(1) GOARCH=amd64 go install -i -pkgdir $(TARGET_PATH)/$(1)_amd64 ./...

endef	

build_a:
	#go mod vendor
	$(foreach e, $(TARGET_OS), $(call INSTALL_GO_LIB,$(e)))