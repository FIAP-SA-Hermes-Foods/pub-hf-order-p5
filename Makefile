build-proto:
	protoc \
	--go_out=order_pub_proto \
	--go_opt=paths=source_relative \
	--go-grpc_out=order_pub_proto \
	--go-grpc_opt=paths=source_relative \
	order_pub.proto
