wire:
	protoc -I . \
	--go-wire-compiler_out protoc \
 	./protoc/*/*.proto


compile-proto:
	protoc -I . \
        --go_out . --go_opt paths=source_relative \
        --go-grpc_out . --go-grpc_opt paths=source_relative --go-grpc_opt=require_unimplemented_servers=false \
        protoc/*/*.proto

grpc_proxy:
	protoc -I . \
	--grpc-gateway_out . --grpc-gateway_opt=logtostderr=true \
	--grpc-gateway_opt=paths=source_relative ./protoc/*/*.proto

proto_swagger:
	protoc -I . \
	--openapiv2_out . \
	--openapiv2_opt logtostderr=true \
	./protoc/*/*.proto

install-grpc:
	go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go install \
      github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
      github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
      google.golang.org/protobuf/cmd/protoc-gen-go \
      google.golang.org/grpc/cmd/protoc-gen-go-grpc