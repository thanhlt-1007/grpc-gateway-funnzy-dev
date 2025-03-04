compile:
	rm -rf pbs;
	mkdir pbs;
	rm -rf swaggers;
	mkdir swaggers;
	protoc \
		--proto_path=protos protos/*.proto \
		--go_out=pbs \
		--go_opt=paths=source_relative \
		--go-grpc_out=pbs \
		--go-grpc_opt=paths=source_relative \
		--go-grpc_opt=require_unimplemented_servers=false \
		--grpc-gateway_out=pbs \
		--grpc-gateway_opt=paths=source_relative \
		--openapiv2_out=swaggers;
