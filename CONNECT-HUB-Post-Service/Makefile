run:
	go run cmd/main.go

dep:
	go mod tidy

proto:
	
	protoc --go_out=. --go-grpc_out=. ./pkg/pb/post/jobseeker/post.proto
	

