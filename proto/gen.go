package proto

// Generating our proto files
//go:generate sh -c "protoc --proto_path=. *.proto --go_out=plugins=grpc:../internal/pb"
