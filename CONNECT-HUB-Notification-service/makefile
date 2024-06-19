run:
	go run cmd/main.go	

proto:
	protoc --go_out=. --go-grpc_out=. ./pkg/pb/auth/auth.proto
	protoc --go_out=. --go-grpc_out=. ./pkg/pb/notification/notification.proto
	
deps:
	go mod tidy

