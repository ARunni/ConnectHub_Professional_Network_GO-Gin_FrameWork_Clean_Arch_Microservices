run:
	go run cmd/main.go

proto:
	protoc --go_out=. --go-grpc_out=. ./pkg/pb/auth/admin/admin.proto
	protoc --go_out=. --go-grpc_out=. ./pkg/pb/auth/jobseeker/jobseeker.proto
	protoc --go_out=. --go-grpc_out=. ./pkg/pb/auth/recruiter/recruiter.proto

all:
	protoc --go_out=. --go-grpc_out=. ./pkg/pb/auth/admin/admin.proto
	protoc --go_out=. --go-grpc_out=. ./pkg/pb/auth/jobseeker/jobseeker.proto
	protoc --go_out=. --go-grpc_out=. ./pkg/pb/auth/recruiter/recruiter.proto
	go run cmd/main.go
