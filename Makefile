.PHONY: proto-gen

proto-gen:
	 protoc --go_out=. ./proto/bundle.proto