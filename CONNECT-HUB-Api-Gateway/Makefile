run:
	go run cmd/main.go

deps:
	go mod tidy

proto:
	protoc --go_out=. --go-grpc_out=. ./pkg/pb/auth/admin/admin.proto
	protoc --go_out=. --go-grpc_out=. ./pkg/pb/auth/jobseeker/jobseeker.proto
	protoc --go_out=. --go-grpc_out=. ./pkg/pb/auth/recruiter/recruiter.proto
	protoc --go_out=. --go-grpc_out=. ./pkg/pb/job/recruiter/job.proto
	protoc --go_out=. --go-grpc_out=. ./pkg/pb/job/jobseeker/job.proto
	protoc --go_out=. --go-grpc_out=. ./pkg/pb/post/jobseeker/post.proto
	protoc --go_out=. --go-grpc_out=. ./pkg/pb/chat/chat.proto
	protoc --go_out=. --go-grpc_out=. ./pkg/pb/auth/auth/auth.proto
	protoc --go_out=. --go-grpc_out=. ./pkg/pb/notification/notification.proto
swag :
	swag init -g cmd/main.go -o ./cmd/docs