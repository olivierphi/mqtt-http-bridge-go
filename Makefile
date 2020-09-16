.PHONY: install
install:
# Install third-party packages in a "vendor/" folder
	go mod vendor

.PHONY: start
start:
	go run -mod=vendor cmd/server/main.go

.PHONY: vet
vet: ## "Govet is concerned with correctness"...
	go vet pkg/*.go

.PHONY: lint
lint: ## ..."whereas golint is concerned with coding style"
# N.B.: Install "go lint" with `go get -u golang.org/x/lint/golint`
	golint pkg/

.PHONY: fmt
fmt:
	go fmt \
		github.com/DrBenton/mqtt-http-bridge-go/pkg \
		github.com/DrBenton/mqtt-http-bridge-go/cmd/server

bin/server: pkg/*.go
	go build -mod=vendor \
		-o bin/server \
		cmd/server/main.go
