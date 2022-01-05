ifndef GOROOT
	export GOROOT=$(realpath ../../go)
	export GOPATH=$(realpath ../)
	export PATH := $(GOROOT)/bin:$(GOPATH)/bin:$(PATH)
endif

run:
	go run server.go

curl:
	curl \
	-H "Content-Type: application/json" \
	-d '{ "query": "mutation  { updateTodoStatus(todoId: [\"todos_01F2QQ3XZDZ0ND8EDZKP15N6ST\",\"todos_01F2QQ40N9GPC3CA55KVTDV17W\"],status: cleared)}" }' \
	http://localhost:8080/query

generate:
	go run github.com/99designs/gqlgen generate --verbose

skeleton:
	go run github.com/99designs/gqlgen init --verbose

install:
	go get github.com/99designs/gqlgen@v0.14.0
	@#go get github.com/99designs/gqlgen@latest

modinit:
	go mod init github.com/siongui/go-graphql-server-client-example/server

modtidy:
	#go list -m all
	go mod tidy
