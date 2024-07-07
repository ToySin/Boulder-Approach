protoc boulder.proto --go_out=. --go_opt=paths=source_relative --swift_out=.

protoc boulder_service.proto --grpc-go_out=. --go-grpc_opt=paths=source_relative --grpc-swift_out=.
