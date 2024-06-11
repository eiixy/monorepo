API_PROTO_FILES=$(shell find api -name *.proto)

gen:
	protoc --proto_path=./api/ \
		--proto_path=./third_party \
		--go_out=paths=source_relative:./api/ \
		--go-http_out=paths=source_relative:./api/ \
		--go-grpc_out=paths=source_relative:./api/ \
		--validate_out=paths=source_relative,lang=go:./api/ \
		--openapi_out=output_mode=source_relative:./api/ \
		$(API_PROTO_FILES)

