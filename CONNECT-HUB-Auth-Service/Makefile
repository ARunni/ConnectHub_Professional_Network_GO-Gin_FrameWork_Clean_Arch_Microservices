run:
	go run cmd/main.go

deps:
	go mod tidy

proto:
	protoc --go_out=. --go-grpc_out=. ./pkg/pb/auth/admin/admin.proto
	protoc --go_out=. --go-grpc_out=. ./pkg/pb/auth/jobseeker/jobseeker.proto
	protoc --go_out=. --go-grpc_out=. ./pkg/pb/auth/recruiter/recruiter.proto
	protoc --go_out=. --go-grpc_out=. ./pkg/pb/job/auth/auth.proto
	protoc --go_out=. --go-grpc_out=. ./pkg/pb/auth/auth/auth.proto

	