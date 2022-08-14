SERVER_OUTPUT_DIR := server/messenger
CLIENT_OUTPUT_DIR := client/src/messenger

compile:
	protoc api/v1/*.proto \
	--go_out=${SERVER_OUTPUT_DIR} --go_opt=paths=source_relative \
	--go-grpc_out=${SERVER_OUTPUT_DIR} --go-grpc_opt=paths=source_relative \