ifndef GOROOT
	export GOROOT=$(realpath ../../go)
	export GOPATH=$(realpath ../)
	export PATH := $(GOROOT)/bin:$(GOPATH)/bin:$(PATH)
endif

skeleton:
	go run github.com/99designs/gqlgen init

install:
	go get github.com/99designs/gqlgen@latest

modinit:
	go mod init github.com/siongui/go-graphql-server-client-example/server

modtidy:
	#go list -m all
	go mod tidy
