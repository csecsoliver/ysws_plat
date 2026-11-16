gen: gen-ruby gen-go

gen-ruby:
	cd backbackend-rb && bundle exec grpc_tools_ruby_protoc -I ../proto \
	  --ruby_out=protobufschema --grpc_out=protobufschema ../proto/greeter.proto

gen-go:
	protoc -I proto \
	  --go_out=frontbackend-go/protobufschema --go_opt=paths=source_relative \
	  --go-grpc_out=frontbackend-go/protobufschema --go-grpc_opt=paths=source_relative \
	  proto/greeter.proto
