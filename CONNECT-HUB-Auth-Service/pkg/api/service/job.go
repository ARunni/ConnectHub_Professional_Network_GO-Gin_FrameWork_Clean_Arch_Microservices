package service

import (
	pb "ConnetHub_auth/pkg/pb/job/auth"
	interfaces "ConnetHub_auth/pkg/usecase/interface"
	"context"
)

type JobAuthServer struct {
	recruiterUsecase interfaces.RecruiterUseCase
	pb.UnimplementedJobAuthServer
}

func NewJobauthServer(usecase interfaces.RecruiterUseCase) pb.JobAuthServer {
	return &JobAuthServer{recruiterUsecase: usecase}
}

func (js *JobAuthServer) GetDetailsById(ctx context.Context, Req *pb.GetDetailsByIdRequest) (*pb.GetDetailsByIdResponse, error) {
	email, name, err := js.recruiterUsecase.GetDetailsById(int(Req.Userid))
	if err != nil {
		return nil, err
	}
	return &pb.GetDetailsByIdResponse{
		Username: name,
		Email:    email,
	}, nil
}
