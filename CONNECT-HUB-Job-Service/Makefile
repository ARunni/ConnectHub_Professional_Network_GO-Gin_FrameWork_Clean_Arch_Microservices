run:
	go run cmd/main.go

proto:
	protoc --go_out=. --go-grpc_out=. ./pkg/pb/job/recruiter/job.proto
	protoc --go_out=. --go-grpc_out=. ./pkg/pb/job/jobseeker/job.proto
	# protoc --go_out=. --go-grpc_out=. ./pkg/pb/job/admin/job.proto

all:
	protoc --go_out=. --go-grpc_out=. ./pkg/pb/job/recruiter/job.proto
	protoc --go_out=. --go-grpc_out=. ./pkg/pb/job/jobseeker/job.proto
	# protoc --go_out=. --go-grpc_out=. ./pkg/pb/job/admin/job.proto
	go run cmd/main.go

deps:
	go mod tidy