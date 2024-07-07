protoc boulder.proto --swift_out=.
protoc boulder.proto --go_out=. 

protoc boulder_service.proto --grpc-swift_out=.
protoc boulder_service.proto --go-grpc_out=. 
