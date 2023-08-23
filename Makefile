gen:
	protoc --proto_path=./proto ./proto/*.proto --plugin=$(GOPATH)/bin/protoc-gen-go --go_out=./pb
	protoc --proto_path=./proto ./proto/*.proto --plugin=$(GOPATH)/bin/protoc-gen-go-grpc --go-grpc_out=./pb

clean:
	rm pb/*.go

run:
	go run main.go

test:
	go test -cover -race ./...
