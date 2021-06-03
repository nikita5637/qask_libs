PROTOC_BIN=protoc
generate:
	protoc --go_out=. --go-grpc_out=require_unimplemented_servers=false:. pb/commonpb/common.proto
	protoc -I=./ --go_out=. --go-grpc_out=require_unimplemented_servers=false:. pb/actionpb/action.proto

.DEFAULT_GOAL: generate