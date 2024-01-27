gen:
	protoc --proto_path=./proto ./proto/*.proto --plugin=$(GOPATH)/bin/protoc-gen-go --go_out=./pb
	protoc --proto_path=./proto ./proto/*.proto --plugin=$(GOPATH)/bin/protoc-gen-go-grpc --go-grpc_out=./pb

clean:
	rm pb/*.go

server:
	clear && go run cmd/server/main.go -port 8080

client:
	clear && go run cmd/client/main.go -address 0.0.0.0:8080

test:
	go test -cover -race ./...
