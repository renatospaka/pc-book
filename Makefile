gen:
	protoc --proto_path=./proto ./proto/*.proto --plugin=$(go env GOPATH)/bin/protoc-gen-go --go_out=./pb

clean:
	rm pb/*.go

run:
	go run main.go
