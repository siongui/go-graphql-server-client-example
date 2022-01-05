ifndef GOROOT
	export GOROOT=$(realpath ../../go)
	export GOPATH=$(realpath ../)
	export PATH := $(GOROOT)/bin:$(GOPATH)/bin:$(PATH)
endif

run:
	go run client.go

fmt:
	go fmt *.go

modinit:
	go mod init github.com/siongui/go-graphql-server-client-example/client

modtidy:
	#go list -m all
	go mod tidy
