PROTOC_BIN=protoc
generate:
	protoc --go_out=. pb/commonpb/common.proto
	protoc -I=./ --go_out=. pb/actionpb/action.proto

.DEFAULT_GOAL: generate