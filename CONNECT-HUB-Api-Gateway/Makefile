run:
	go run cmd/main.go

dep:
	go mod tidy

proto:
	protoc --go_out=. --go-grpc_out=. ./pkg/pb/auth/admin/admin.proto
	protoc --go_out=. --go-grpc_out=. ./pkg/pb/auth/jobseeker/jobseeker.proto
	protoc --go_out=. --go-grpc_out=. ./pkg/pb/auth/recruiter/recruiter.proto
	protoc --go_out=. --go-grpc_out=. ./pkg/pb/job/recruiter/job.proto
	protoc --go_out=. --go-grpc_out=. ./pkg/pb/job/jobseeker/job.proto
	protoc --go_out=. --go-grpc_out=. ./pkg/pb/post/jobseeker/post.proto
	protoc --go_out=. --go-grpc_out=. ./pkg/pb/chat/chat.proto
