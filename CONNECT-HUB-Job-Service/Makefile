run:
	go run cmd/main.go

proto:
	protoc --go_out=. --go-grpc_out=. ./pkg/pb/job/job.proto
	# protoc --go_out=. --go-grpc_out=. ./pkg/pb/job/jobseeker.proto
	# protoc --go_out=. --go-grpc_out=. ./pkg/pb/job/recruiter.proto

all:
	protoc --go_out=. --go-grpc_out=. ./pkg/pb/job/job.proto
	# protoc --go_out=. --go-grpc_out=. ./pkg/pb/job/jobseeker.proto
	# protoc --go_out=. --go-grpc_out=. ./pkg/pb/job/recruiter.proto
	go run cmd/main.go