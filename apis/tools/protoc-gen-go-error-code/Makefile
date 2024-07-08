.PHONY: build
build:
	go build -o protoc-gen-go-error-code ./ && sudo cp ./protoc-gen-go-error-code /usr/local/bin/protoc-gen-go-error-code \
	&& sudo cp ./protoc-gen-go-error-code /usr/local/bin/protoc-gen-go-error-code

api:
	protoc --proto_path=. --proto_path=../../../apis/third_party \
	   --go_out=paths=source_relative:. \
	   --go-error-code_out=paths=source_relative:. \
	   errors.proto