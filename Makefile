generate_grpc_code:
	protoc \
	--go_out=allocator \
	--go_opt=paths=source_relative \
	--go-grpc_out=allocator \
	--go-grpc_opt=paths=source_relative \
	allocator.proto