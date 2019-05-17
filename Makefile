.EXPORT_ALL_VARIABLES:

.PHONY: all
all: build

GOPATH  := $(shell pwd):$(HOME)/go

all: build

build:
        go build
