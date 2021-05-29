protoc --proto_path=protos --go_out=./.. models.proto

protoc --proto_path=protos --go_out=./.. --validate_out="lang=go:./../"  models.proto

protoc --proto_path=protos --plugin=protoc-gen-go --go-grpc_out=./../ service.proto