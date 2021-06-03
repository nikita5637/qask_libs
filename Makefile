PROTOGEN_FLAGS=-I=./ --go_out=. --go_opt=paths=source_relative --go-grpc_out=require_unimplemented_servers=false,paths=source_relative:.
PROTOC_BIN=protoc
generate:
	$(PROTOC_BIN) $(PROTOGEN_FLAGS) pb/commonpb/common.proto
	$(PROTOC_BIN) $(PROTOGEN_FLAGS) pb/actionpb/action.proto

clean:
	rm pb/actionpb/*.go pb/commonpb/*.go

.DEFAULT_GOAL: generate