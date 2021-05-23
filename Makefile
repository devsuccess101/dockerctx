VERSION ?= v0.0.0+unknown

all: linux

linux:
	go build -ldflags "-X main.Version=$(VERSION)" -o dist/dockerctx cmd/dockerctx/main.go
