run:
	go run cmd/main.go

deps:
	go mod tidy

proto:
	
	protoc --go_out=. --go-grpc_out=. ./pkg/pb/post/jobseeker/post.proto
	protoc --go_out=. --go-grpc_out=. ./pkg/pb/auth/auth.proto
	

