.PHONY: build proto-gen

proto-gen:
	 protoc --go_out=. ./proto/meta.proto

build:
	go build -o build/bundler ./cmd/bundler/main.go