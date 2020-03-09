.PHONY:proto
proto:
	protoc -I api/ api.proto --go_out=plugins=grpc:api