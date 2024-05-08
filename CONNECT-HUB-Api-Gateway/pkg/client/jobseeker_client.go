package client

import (
	interfaces "connectHub_gateway/pkg/client/interface"
	"connectHub_gateway/pkg/config"
	pb "connectHub_gateway/pkg/pb/auth/jobseeker"
	"connectHub_gateway/pkg/utils/models"
	"context"
	"fmt"

	"google.golang.org/grpc"
)

type jobseekerClient struct {
	Client pb.JobseekerClient
}

func NewJobSeekerClient(cfg config.Config) interfaces.JobSeekerClient {
	fmt.Println("jobseeker client")
	grpcConnection, err := grpc.Dial(cfg.ConnetHubAuth, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not Connect to Auth jobseeker", err)
	}

	grpcClient := pb.NewJobseekerClient(grpcConnection)
	return &jobseekerClient{
		Client: grpcClient,
	}

}

func (jc *jobseekerClient) JobSeekerSignup(jobseekerData models.JobSeekerSignUp) (models.TokenJobSeeker, error) {
	jobseeker, err := jc.Client.JobSeekerSignup(context.Background(), &pb.JobSeekerSignupRequest{
		Firstname:   jobseekerData.FirstName,
		Lastname:    jobseekerData.LastName,
		Password:    jobseekerData.Password,
		Email:       jobseekerData.Email,
		PhoneNumber: jobseekerData.PhoneNumber,
		DateOfBirth: jobseekerData.DateOfBirth,
		Gender:      jobseekerData.Gender,
	})
	if err != nil {
		return models.TokenJobSeeker{}, err
	}
	return models.TokenJobSeeker{
		JobSeeker: models.JobSeekerDetailsResponse{
			ID:          uint(jobseeker.JobSeekerDetails.Id),
			Email:       jobseeker.JobSeekerDetails.Email,
			FirstName:   jobseeker.JobSeekerDetails.Firstname,
			LastName:    jobseeker.JobSeekerDetails.Lastname,
			PhoneNumber: jobseeker.JobSeekerDetails.PhoneNumber,
			DateOfBirth: jobseeker.JobSeekerDetails.DateOfBirth,
		},
		Token: jobseeker.Token,
	}, nil
}

func (jc *jobseekerClient) JobSeekerLogin(jobseekerData models.JobSeekerLogin) (models.TokenJobSeeker, error) {
	jobseeker, err := jc.Client.JobSeekerLogin(context.Background(), &pb.JobSeekerLoginRequest{
		Email:    jobseekerData.Email,
		Password: jobseekerData.Password,
	})
	if err != nil {
		return models.TokenJobSeeker{}, err
	}

	return models.TokenJobSeeker{
		JobSeeker: models.JobSeekerDetailsResponse{
			ID:          uint(jobseeker.JobSeekerDetails.Id),
			FirstName:   jobseeker.JobSeekerDetails.Firstname,
			LastName:    jobseeker.JobSeekerDetails.Lastname,
			Email:       jobseeker.JobSeekerDetails.Email,
			PhoneNumber: jobseeker.JobSeekerDetails.PhoneNumber,
		},
		Token: jobseeker.Token,
	}, nil

}
