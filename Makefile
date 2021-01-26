export GO111MODULE=on
export TF_LOG=DEBUG
SRC=$(shell find . -name '*.go')

.PHONY: all vet build test

all: build

build:
	go build -a -o terraform-provider-sonarcloud
	mkdir -p ~/.terraform.d/plugins/github.com/jkumar19/sonarcloud/0.1/linux_amd64/
	cp terraform-provider-sonarcloud ~/.terraform.d/plugins/github.com/jkumar19/sonarcloud/0.1/linux_amd64/
