package client

import (
	interfaces "connectHub_gateway/pkg/client/auth/interface"
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

func NewJobSeekerAuthClient(cfg config.Config) interfaces.JobSeekerAuthClient {
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
		Firstname:       jobseekerData.FirstName,
		Lastname:        jobseekerData.LastName,
		Password:        jobseekerData.Password,
		Email:           jobseekerData.Email,
		PhoneNumber:     jobseekerData.PhoneNumber,
		DateOfBirth:     jobseekerData.DateOfBirth,
		Gender:          jobseekerData.Gender,
		ConfirmPassword: jobseekerData.ConfirmPassword,
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

func (jc *jobseekerClient) JobSeekerGetProfile(id int) (models.JobSeekerProfile, error) {
	profile, err := jc.Client.JobSeekerGetProfile(context.Background(), &pb.GetProfileRequest{
		Id: int32(id),
	})
	if err != nil {
		return models.JobSeekerProfile{}, err
	}

	return models.JobSeekerProfile{
		ID:          uint(profile.Profile.Id),
		Email:       profile.Profile.Email,
		FirstName:   profile.Profile.FirstName,
		LastName:    profile.Profile.LastName,
		PhoneNumber: profile.Profile.Phone,
		DateOfBirth: profile.Profile.DateOfBirth,
		Gender:      profile.Profile.Gender,
	}, nil

}

func (jc *jobseekerClient) JobSeekerEditProfile(profile models.JobSeekerProfile) (models.JobSeekerProfile, error) {
	profileData, err := jc.Client.JobSeekerEditProfile(context.Background(), &pb.JobSeekerEditProfileRequest{
		Profile: &pb.JobSeekerProfile{
			Id:          uint64(profile.ID),
			FirstName:   profile.FirstName,
			Gender:      profile.Gender,
			Email:       profile.Email,
			DateOfBirth: profile.DateOfBirth,
			Phone:       profile.PhoneNumber,
			LastName:    profile.LastName,
		},
	})
	if err != nil {
		return models.JobSeekerProfile{}, err
	}

	return models.JobSeekerProfile{
		ID:          uint(profileData.Profile.Id),
		FirstName:   profile.FirstName,
		Gender:      profile.Gender,
		Email:       profile.Email,
		DateOfBirth: profile.DateOfBirth,
		PhoneNumber: profile.PhoneNumber,
		LastName:    profile.LastName,
	}, nil

}