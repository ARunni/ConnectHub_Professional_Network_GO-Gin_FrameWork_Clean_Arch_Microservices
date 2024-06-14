package client

import (
	"context"
	"fmt"
	"os"

	logging "github.com/ARunni/connectHub_gateway/Logging"
	interfaces "github.com/ARunni/connectHub_gateway/pkg/client/auth/interface"
	"github.com/ARunni/connectHub_gateway/pkg/config"
	"github.com/ARunni/connectHub_gateway/pkg/helper"
	pb "github.com/ARunni/connectHub_gateway/pkg/pb/auth/admin"
	"github.com/ARunni/connectHub_gateway/pkg/utils/models"

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
	ac.Logger.Info("Adminlogin on client started")
	admin, err := ac.Client.AdminLogin(context.Background(), &pb.AdminLoginInRequest{
		Email:    admindata.Email,
		Password: admindata.Password,
	})
	if err != nil {
		ac.Logger.Error("Error in AdminLogin", err)
		return models.TokenAdmin{}, err
	}
	ac.Logger.Info("Adminlogin on client successfull")
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
	ac.Logger.Info("GetRecruiters on client started")
	recruiter, err := ac.Client.GetRecruiters(context.Background(), &pb.GetRecruiterRequest{
		Page: int32(page),
	})
	if err != nil {
		ac.Logger.Error("Error in GetRecruiters", err)
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
	ac.Logger.Info("GetRecruiters on client successfull")
	return resp, nil

}

func (ac *adminClient) GetJobseekers(page int) ([]models.JobseekerDetailsAtAdmin, error) {
ac.Logger.Info("GetJobseekers on client started")
	jobSeeker, err := ac.Client.GetJobseekers(context.Background(), &pb.GetJobseekerRequest{
		Page: int32(page),
	})
	if err != nil {
		ac.Logger.Error("Error in GetJobseekers", err)
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
	ac.Logger.Info("GetJobseekers on client successfull")
	return resp, nil
}

func (ac *adminClient) BlockRecruiter(id int) (models.BlockRes, error) {
	ac.Logger.Info("BlockRecruiter on client started")
	recruterBlock, err := ac.Client.BlockRecruiter(context.Background(), &pb.BlockRecruiterRequest{
		GetRecruiterId: int64(id),
	})
	if err != nil {
		ac.Logger.Error("Error in BlockRecruiter", err)
		return models.BlockRes{}, err
	}
	ac.Logger.Info("BlockRecruiter on client successfull")
	return models.BlockRes{
		Status: recruterBlock.Status,
	}, nil
}

func (ac *adminClient) BlockJobseeker(id int) (models.BlockRes, error) {
	ac.Logger.Info("BlockJobseeker on client started")
	jobseekerBlock, err := ac.Client.BlockJobseeker(context.Background(), &pb.BlockJobseekerRequest{
		JobseekerId: int64(id),
	})
	if err != nil {
		ac.Logger.Error("Error in BlockJobseeker", err)
		return models.BlockRes{}, err
	}
	ac.Logger.Info("BlockJobseeker on client successfull")
	return models.BlockRes{
		Status: jobseekerBlock.Status,
	}, nil
}

func (ac *adminClient) UnBlockRecruiter(id int) (models.BlockRes, error) {
	ac.Logger.Info("UnBlockRecruiter on client started")
	recruiterUnBlock, err := ac.Client.UnBlockRecruiter(context.Background(), &pb.UnBlockRecruiterRequest{
		GetRecruiterId: int64(id),
	})
	if err != nil {
		ac.Logger.Error("Error in UnBlockRecruiter", err)
		return models.BlockRes{}, err
	}
	ac.Logger.Info("UnBlockRecruiter on client successfull")
	return models.BlockRes{
		Status: recruiterUnBlock.Status,
	}, nil

}

func (ac *adminClient) UnBlockJobseeker(id int) (models.BlockRes, error) {
	ac.Logger.Info("UnBlockJobseeker on client started")
	jobseekerUnBlock, err := ac.Client.UnBlockJobseeker(context.Background(), &pb.UnBlockJobseekerRequest{
		JobseekerId: int64(id),
	})
	if err != nil {
		ac.Logger.Error("Error in UnBlockJobseeker", err)
		return models.BlockRes{}, err
	}
	ac.Logger.Info("UnBlockJobseeker on client successfull")
	return models.BlockRes{
		Status: jobseekerUnBlock.Status,
	}, nil
}

func (ac *adminClient) GetJobseekerDetails(id int) (models.JobseekerDetailsAtAdmin, error) {
	ac.Logger.Info("GetJobseekerDetails on client started")
	jobseekerData, err := ac.Client.GetJobseekerDetails(context.Background(), &pb.GetJobseekerDetailsRequest{
		JobseekerId: int64(id),
	})
	if err != nil {
		ac.Logger.Error("Error in GetJobseekerDetails", err)
		return models.JobseekerDetailsAtAdmin{}, err
	}
	ac.Logger.Info("GetJobseekerDetails on client successfull")
	return models.JobseekerDetailsAtAdmin{
		Id:      int(jobseekerData.Id),
		Name:    jobseekerData.Firstname,
		Email:   jobseekerData.Email,
		Phone:   jobseekerData.PhoneNumber,
		Blocked: jobseekerData.Blocked,
	}, nil
}

func (ac *adminClient) GetRecruiterDetails(id int) (models.RecruiterDetailsAtAdmin, error) {
	ac.Logger.Info("GetRecruiterDetails on client started")
	recruiterdata, err := ac.Client.GetRecruiterDetails(context.Background(), &pb.GetRecruiterDetailsRequest{
		RecruiterId: int64(id),
	})
	if err != nil {
		ac.Logger.Error("Error in GetRecruiterDetails", err)
		return models.RecruiterDetailsAtAdmin{}, err
	}
	ac.Logger.Info("GetRecruiterDetails on client successfull")
	return models.RecruiterDetailsAtAdmin{
		Id:          int(recruiterdata.Id),
		CompanyName: recruiterdata.CompanyName,
		Email:       recruiterdata.Email,
		Phone:       recruiterdata.PhoneNumber,
		Blocked:     recruiterdata.Blocked,
	}, nil
}

func (ac *adminClient) CreatePolicy(data models.CreatePolicyReq) (models.CreatePolicyRes, error) {
	ac.Logger.Info("CreatePolicy on client started")
	recruiterdata, err := ac.Client.CreatePolicy(context.Background(), &pb.CreatePolicyRequest{
		Title:   data.Title,
		Content: data.Content,
	})
	if err != nil {
		ac.Logger.Error("Error in CreatePolicy", err)
		return models.CreatePolicyRes{}, err
	}
	ac.Logger.Info("CreatePolicy on client successfull")
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
	ac.Logger.Info("UpdatePolicy on client started")
	recruiterdata, err := ac.Client.UpdatePolicy(context.Background(), &pb.UpdatePolicyRequest{
		Id:      int64(data.Id),
		Title:   data.Title,
		Content: data.Content,
	})
	if err != nil {
		ac.Logger.Error("Error in UpdatePolicy", err)
		return models.CreatePolicyRes{}, err
	}
	ac.Logger.Info("UpdatePolicy on client successfull")
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
	ac.Logger.Info("DeletePolicy on client started")
	recruiterdata, err := ac.Client.DeletePolicy(context.Background(), &pb.DeletePolicyRequest{
		Id: int64(policy_id),
	})
	if err != nil {
		ac.Logger.Error("Error in DeletePolicy", err)
		return false, err
	}
	ac.Logger.Info("DeletePolicy on client successfull")
	return recruiterdata.Deleted, nil

}

func (ac *adminClient) GetAllPolicies() (models.GetAllPolicyRes, error) {
	ac.Logger.Info("GetAllPolicies on client started")
	recruiterdata, err := ac.Client.GetAllPolicies(context.Background(), &pb.GetAllPoliciesRequest{})
	if err != nil {
		ac.Logger.Error("Error in GetAllPolicies", err)
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
	ac.Logger.Info("GetAllPolicies on client successfull")
	return models.GetAllPolicyRes{
		Policies: policies,
	}, nil
}

func (ac *adminClient) GetOnePolicy(policy_id int) (models.CreatePolicyRes, error) {
	ac.Logger.Info("GetOnePolicy on client started")
	recruiterdata, err := ac.Client.GetOnePolicy(context.Background(), &pb.GetOnePolicyRequest{Id: int64(policy_id)})
	if err != nil {
		ac.Logger.Error("Error in GetOnePolicy", err)
		return models.CreatePolicyRes{}, err
	}
	ac.Logger.Info("GetOnePolicy on client successfull")

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
