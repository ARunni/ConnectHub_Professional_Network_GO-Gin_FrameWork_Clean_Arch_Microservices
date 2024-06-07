run:
	go run cmd/main.go

dep:
	go mod tidy

proto:
	protoc --go_out=. --go-grpc_out=. ./pkg/pb/job/recruiter/job.proto
	protoc --go_out=. --go-grpc_out=. ./pkg/pb/job/jobseeker/job.proto
	protoc --go_out=. --go-grpc_out=. ./pkg/pb/auth/auth.proto

