SRC_PATH:= ${PWD}

proto-gen:
	protoc --proto_path=proto proto/*.proto --go_out=. --validate_out=lang=go:. --go-grpc_out=:. --grpc-gateway_out=:. --openapiv2_out=:docs

build:
	@echo "building..."
	go build ./cmd/srv/.