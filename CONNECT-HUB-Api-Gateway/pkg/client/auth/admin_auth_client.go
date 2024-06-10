package client

import (
	logging "connectHub_gateway/Logging"
	interfaces "connectHub_gateway/pkg/client/auth/interface"
	"connectHub_gateway/pkg/config"
	"connectHub_gateway/pkg/helper"
	pb "connectHub_gateway/pkg/pb/auth/admin"
	"connectHub_gateway/pkg/utils/models"
	"context"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type adminClient struct {
	Client  pb.AdminClient
	Logger  *logrus.Logger
	LogFile *os.File
}

func NewAdminAuthClient(cfg config.Config) interfaces.AdminAuthClient {

	logger, logFile := logging.InitLogrusLogger("./Logging/connectHub_gateway.log")

	grpcConnection, err := grpc.Dial(cfg.ConnetHubAuth, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not Connect to Auth", err)
	}
	grpcClient := pb.NewAdminClient(grpcConnection)
	return &adminClient{
		Client:  grpcClient,
		Logger:  logger,
		LogFile: logFile,
	}

}

func (ac *adminClient) AdminLogin(admindata models.AdminLogin) (models.TokenAdmin, error) {

	admin, err := ac.Client.AdminLogin(context.Background(), &pb.AdminLoginInRequest{
		Email:    admindata.Email,
		Password: admindata.Password,
	})
	if err != nil {
		return models.TokenAdmin{}, err
	}
	return models.TokenAdmin{
		Admin: models.AdminDetailsResponse{
			ID:        uint(admin.AdminDetails.Id),
			Email:     admin.AdminDetails.Email,
			Lastname:  admin.AdminDetails.Lastname,
			Firstname: admin.AdminDetails.Firstname,
		},
		Token: admin.Token,
	}, nil

}
func (ac *adminClient) GetRecruiters(page int) ([]models.RecruiterDetailsAtAdmin, error) {

	recruiter, err := ac.Client.GetRecruiters(context.Background(), &pb.GetRecruiterRequest{
		Page: int32(page),
	})
	if err != nil {
		return nil, err
	}
	var resp []models.RecruiterDetailsAtAdmin

	for _, data := range recruiter.Recruiter {
		resp = append(resp, models.RecruiterDetailsAtAdmin{
			Id:          int(data.Id),
			Email:       data.Email,
			CompanyName: data.CompanyName,
			Phone:       data.PhoneNumber,
			Blocked:     data.Blocked,
		})
	}
	return resp, nil

}

func (ac *adminClient) GetJobseekers(page int) ([]models.JobseekerDetailsAtAdmin, error) {

	jobSeeker, err := ac.Client.GetJobseekers(context.Background(), &pb.GetJobseekerRequest{
		Page: int32(page),
	})
	if err != nil {
		return nil, err
	}
	var resp []models.JobseekerDetailsAtAdmin

	for _, data := range jobSeeker.Jobseeker {
		resp = append(resp, models.JobseekerDetailsAtAdmin{
			Id:      int(data.Id),
			Email:   data.Email,
			Name:    data.Firstname,
			Phone:   data.PhoneNumber,
			Blocked: data.Blocked,
		})
	}
	return resp, nil
}

func (ac *adminClient) BlockRecruiter(id int) (models.BlockRes, error) {

	recruterBlock, err := ac.Client.BlockRecruiter(context.Background(), &pb.BlockRecruiterRequest{
		GetRecruiterId: int64(id),
	})
	if err != nil {
		return models.BlockRes{}, err
	}
	return models.BlockRes{
		Status: recruterBlock.Status,
	}, nil
}

func (ac *adminClient) BlockJobseeker(id int) (models.BlockRes, error) {

	jobseekerBlock, err := ac.Client.BlockJobseeker(context.Background(), &pb.BlockJobseekerRequest{
		JobseekerId: int64(id),
	})
	if err != nil {
		return models.BlockRes{}, err
	}
	return models.BlockRes{
		Status: jobseekerBlock.Status,
	}, nil
}

func (ac *adminClient) UnBlockRecruiter(id int) (models.BlockRes, error) {

	recruiterUnBlock, err := ac.Client.UnBlockRecruiter(context.Background(), &pb.UnBlockRecruiterRequest{
		GetRecruiterId: int64(id),
	})
	if err != nil {
		return models.BlockRes{}, err
	}
	return models.BlockRes{
		Status: recruiterUnBlock.Status,
	}, nil

}

func (ac *adminClient) UnBlockJobseeker(id int) (models.BlockRes, error) {

	jobseekerUnBlock, err := ac.Client.UnBlockJobseeker(context.Background(), &pb.UnBlockJobseekerRequest{
		JobseekerId: int64(id),
	})
	if err != nil {
		return models.BlockRes{}, err
	}
	return models.BlockRes{
		Status: jobseekerUnBlock.Status,
	}, nil
}

func (ac *adminClient) GetJobseekerDetails(id int) (models.JobseekerDetailsAtAdmin, error) {

	jobseekerData, err := ac.Client.GetJobseekerDetails(context.Background(), &pb.GetJobseekerDetailsRequest{
		JobseekerId: int64(id),
	})
	if err != nil {
		return models.JobseekerDetailsAtAdmin{}, err
	}
	return models.JobseekerDetailsAtAdmin{
		Id:      int(jobseekerData.Id),
		Name:    jobseekerData.Firstname,
		Email:   jobseekerData.Email,
		Phone:   jobseekerData.PhoneNumber,
		Blocked: jobseekerData.Blocked,
	}, nil
}

func (ac *adminClient) GetRecruiterDetails(id int) (models.RecruiterDetailsAtAdmin, error) {

	recruiterdata, err := ac.Client.GetRecruiterDetails(context.Background(), &pb.GetRecruiterDetailsRequest{
		RecruiterId: int64(id),
	})
	if err != nil {
		return models.RecruiterDetailsAtAdmin{}, err
	}
	return models.RecruiterDetailsAtAdmin{
		Id:          int(recruiterdata.Id),
		CompanyName: recruiterdata.CompanyName,
		Email:       recruiterdata.Email,
		Phone:       recruiterdata.PhoneNumber,
		Blocked:     recruiterdata.Blocked,
	}, nil
}

func (ac *adminClient) CreatePolicy(data models.CreatePolicyReq) (models.CreatePolicyRes, error) {

	recruiterdata, err := ac.Client.CreatePolicy(context.Background(), &pb.CreatePolicyRequest{
		Title:   data.Title,
		Content: data.Content,
	})
	if err != nil {
		return models.CreatePolicyRes{}, err
	}
	return models.CreatePolicyRes{
		Policies: models.Policy{
			ID:        uint(recruiterdata.Policy.Id),
			Title:     recruiterdata.Policy.Title,
			Content:   recruiterdata.Policy.Content,
			CreatedAt: helper.FromProtoTimestamp(recruiterdata.Policy.CreatedAt),
			UpdatedAt: helper.FromProtoTimestamp(recruiterdata.Policy.UpdatedAt),
		},
	}, nil
}

func (ac *adminClient) UpdatePolicy(data models.UpdatePolicyReq) (models.CreatePolicyRes, error) {

	recruiterdata, err := ac.Client.UpdatePolicy(context.Background(), &pb.UpdatePolicyRequest{
		Id:      int64(data.Id),
		Title:   data.Title,
		Content: data.Content,
	})
	if err != nil {
		return models.CreatePolicyRes{}, err
	}
	return models.CreatePolicyRes{
		Policies: models.Policy{
			ID:        uint(recruiterdata.Policy.Id),
			Title:     recruiterdata.Policy.Title,
			Content:   recruiterdata.Policy.Content,
			CreatedAt: helper.FromProtoTimestamp(recruiterdata.Policy.CreatedAt),
			UpdatedAt: helper.FromProtoTimestamp(recruiterdata.Policy.UpdatedAt),
		},
	}, nil
}

func (ac *adminClient) DeletePolicy(policy_id int) (bool, error) {

	recruiterdata, err := ac.Client.DeletePolicy(context.Background(), &pb.DeletePolicyRequest{
		Id: int64(policy_id),
	})
	if err != nil {
		return false, err
	}
	return recruiterdata.Deleted, nil

}

func (ac *adminClient) GetAllPolicies() (models.GetAllPolicyRes, error) {

	recruiterdata, err := ac.Client.GetAllPolicies(context.Background(), &pb.GetAllPoliciesRequest{})
	if err != nil {
		return models.GetAllPolicyRes{}, err
	}
	var policies []models.Policy
	for _, policy := range recruiterdata.Policies {
		policies = append(policies, models.Policy{
			ID:        uint(policy.Id),
			Title:     policy.Title,
			Content:   policy.Content,
			CreatedAt: helper.FromProtoTimestamp(policy.CreatedAt),
			UpdatedAt: helper.FromProtoTimestamp(policy.UpdatedAt),
		})

	}
	return models.GetAllPolicyRes{
		Policies: policies,
	}, nil
}

func (ac *adminClient) GetOnePolicy(policy_id int) (models.CreatePolicyRes, error) {

	recruiterdata, err := ac.Client.GetOnePolicy(context.Background(), &pb.GetOnePolicyRequest{Id: int64(policy_id)})
	if err != nil {
		return models.CreatePolicyRes{}, err
	}

	return models.CreatePolicyRes{
		Policies: models.Policy{
			ID:        uint(recruiterdata.Policy.Id),
			Title:     recruiterdata.Policy.Title,
			Content:   recruiterdata.Policy.Content,
			CreatedAt: helper.FromProtoTimestamp(recruiterdata.Policy.CreatedAt),
			UpdatedAt: helper.FromProtoTimestamp(recruiterdata.Policy.UpdatedAt),
		},
	}, nil
}
