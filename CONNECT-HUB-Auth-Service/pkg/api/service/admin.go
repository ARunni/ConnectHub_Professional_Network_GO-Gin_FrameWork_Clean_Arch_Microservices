package service

import (
	"context"
	"os"

	logging "github.com/ARunni/ConnetHub_auth/Logging"
	pb "github.com/ARunni/ConnetHub_auth/pkg/pb/auth/admin"
	interfaces "github.com/ARunni/ConnetHub_auth/pkg/usecase/interface"
	req "github.com/ARunni/ConnetHub_auth/pkg/utils/reqAndResponse"

	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type AdminServer struct {
	adminUseCase interfaces.AdminUseCase
	pb.UnimplementedAdminServer
	Logger  *logrus.Logger
	LogFile *os.File
}

func NewAdminServer(useCase interfaces.AdminUseCase) pb.AdminServer {
	logger, logFile := logging.InitLogrusLogger("./Logging/connectHub_Auth.log")
	return &AdminServer{
		adminUseCase: useCase,
		Logger:       logger,
		LogFile:      logFile,
	}
}

func (ad *AdminServer) AdminLogin(ctx context.Context, Req *pb.AdminLoginInRequest) (*pb.AdminLoginResponse, error) {
	ad.Logger.Info("AdminLogin at AdminServer started")
	adminLogin := req.AdminLogin{
		Email:    Req.Email,
		Password: Req.Password,
	}
	ad.Logger.Info("AdminLogin at usecase started")
	admin, err := ad.adminUseCase.AdminLogin(adminLogin)
	if err != nil {
		ad.Logger.Error("Error in AdminLogin at usecase: ", err)
		return &pb.AdminLoginResponse{}, err
	}
	adminDetails := &pb.AdminDetails{
		Id:        uint64(admin.Admin.ID),
		Firstname: admin.Admin.Firstname,
		Lastname:  admin.Admin.Lastname,
		Email:     admin.Admin.Email,
	}
	ad.Logger.Info("AdminLogin at usecase success")
	ad.Logger.Info("AdminLogin at AdminServer success")
	return &pb.AdminLoginResponse{
		Status:       200,
		AdminDetails: adminDetails,
		Token:        admin.Token,
	}, nil
}

func (as *AdminServer) GetJobseekers(ctx context.Context, Req *pb.GetJobseekerRequest) (*pb.GetJobseekerResponse, error) {
	as.Logger.Info("GetJobseekers at AdminServer started")
	page := Req.Page
	as.Logger.Info("GetJobseekers at adminUseCase started")
	GetJobseeker, err := as.adminUseCase.GetJobseekers(int(page))
	if err != nil {
		as.Logger.Error("Error in GetJobseekers at usecase: ", err)
		return &pb.GetJobseekerResponse{}, err
	}
	var jobseekerDetails []*pb.JobSeekerDetails
	for _, data := range GetJobseeker {
		jobseekerDetails = append(jobseekerDetails, &pb.JobSeekerDetails{
			Id:          uint64(data.Id),
			Email:       data.Email,
			Firstname:   data.Name,
			PhoneNumber: data.Phone,
			Blocked:     data.Blocked,
		})
	}
	as.Logger.Info("GetJobseekers at adminUseCase success")
	as.Logger.Info("GetJobseekers at AdminServer success")
	return &pb.GetJobseekerResponse{
		Jobseeker: jobseekerDetails,
	}, nil

}

func (as *AdminServer) GetRecruiters(ctx context.Context, Req *pb.GetRecruiterRequest) (*pb.GetRecruitersResponse, error) {
	as.Logger.Info("GetRecruiters at AdminServer started")

	page := Req.Page
	as.Logger.Info("GetRecruiters at adminUseCase started")

	GetRecruiters, err := as.adminUseCase.GetRecruiters(int(page))

	if err != nil {
		as.Logger.Error("Error in GetRecruiters at usecase: ", err)
		return &pb.GetRecruitersResponse{}, err
	}

	var recruitersDetails []*pb.RecruiterDetails
	for _, data := range GetRecruiters {
		recruitersDetails = append(recruitersDetails, &pb.RecruiterDetails{
			Id:          uint64(data.Id),
			Email:       data.Contact_mail,
			CompanyName: data.CompanyName,
			PhoneNumber: data.Phone,
			Blocked:     data.Blocked,
		})
	}
	as.Logger.Info("GetRecruiters at adminUseCase success")
	as.Logger.Info("GetRecruiters at AdminServer success")
	return &pb.GetRecruitersResponse{
		Recruiter: recruitersDetails,
	}, nil

}

func (as *AdminServer) BlockRecruiter(ctx context.Context, Req *pb.BlockRecruiterRequest) (*pb.BlockRecruiterResponse, error) {
	as.Logger.Info("BlockRecruiter at AdminServer started")

	recruiterId := Req.GetRecruiterId
	as.Logger.Info("BlockRecruiter at adminUseCase started")

	result, err := as.adminUseCase.BlockRecruiter(int(recruiterId))
	if err != nil {
		as.Logger.Error("Error in BlockRecruiter at usecase: ", err)
		return &pb.BlockRecruiterResponse{}, err
	}
	as.Logger.Info("BlockRecruiter at adminUseCase success")
	as.Logger.Info("BlockRecruiter at AdminServer success")
	return &pb.BlockRecruiterResponse{
		Status: result.Status,
	}, nil
}

func (as *AdminServer) UnBlockRecruiter(ctx context.Context, Req *pb.UnBlockRecruiterRequest) (*pb.UnBlockRecruiterResponse, error) {
	as.Logger.Info("UnBlockRecruiter at AdminServer started")
	recruiterId := Req.GetRecruiterId
	as.Logger.Info("UnBlockRecruiter at adminUseCase started")

	result, err := as.adminUseCase.UnBlockRecruiter(int(recruiterId))
	if err != nil {
		as.Logger.Error("Error in UnBlockRecruiter at usecase: ", err)
		return &pb.UnBlockRecruiterResponse{}, err
	}
	as.Logger.Info("UnBlockRecruiter at adminUseCase success")
	as.Logger.Info("UnBlockRecruiter at AdminServer success")
	return &pb.UnBlockRecruiterResponse{
		Status: result.Status,
	}, nil
}

func (as *AdminServer) UnBlockJobseeker(ctx context.Context, Req *pb.UnBlockJobseekerRequest) (*pb.UnBlockJobseekerResponse, error) {
	as.Logger.Info("UnBlockJobseeker at AdminServer started")
	jobseekerId := Req.JobseekerId
	as.Logger.Info("UnBlockJobseeker at adminUseCase started")

	result, err := as.adminUseCase.UnBlockJobseeker(int(jobseekerId))
	if err != nil {
		as.Logger.Error("Error in UnBlockJobseeker at usecase: ", err)
		return &pb.UnBlockJobseekerResponse{}, err
	}
	as.Logger.Info("UnBlockJobseeker at adminUseCase success")
	as.Logger.Info("UnBlockJobseeker at AdminServer success")
	return &pb.UnBlockJobseekerResponse{
		Status: result.Status,
	}, nil
}

func (as *AdminServer) BlockJobseeker(ctx context.Context, Req *pb.BlockJobseekerRequest) (*pb.BlockJobseekerResponse, error) {
	as.Logger.Info("BlockJobseeker at AdminServer started")
	jobseekerId := Req.JobseekerId
	as.Logger.Info("BlockJobseeker at adminUseCase started")

	result, err := as.adminUseCase.BlockJobseeker(int(jobseekerId))
	if err != nil {
		as.Logger.Error("Error in BlockJobseeker at usecase: ", err)
		return &pb.BlockJobseekerResponse{}, err
	}
	as.Logger.Info("BlockJobseeker at adminUseCase success")
	as.Logger.Info("BlockJobseeker at AdminServer success")
	return &pb.BlockJobseekerResponse{
		Status: result.Status,
	}, nil
}

func (as *AdminServer) GetJobseekerDetails(ctx context.Context, Req *pb.GetJobseekerDetailsRequest) (*pb.GetJobseekerDetailsResponse, error) {
	as.Logger.Info("GetJobseekerDetails at AdminServer started")
	jobseekerId := Req.JobseekerId
	as.Logger.Info("GetJobseekerDetails at adminUseCase started")

	result, err := as.adminUseCase.GetJobseekerDetails(int(jobseekerId))
	if err != nil {
		as.Logger.Error("Error in GetJobseekerDetails at usecase: ", err)
		return &pb.GetJobseekerDetailsResponse{}, err
	}
	as.Logger.Info("GetJobseekerDetails at adminUseCase success")
	as.Logger.Info("GetJobseekerDetails at AdminServer success")
	return &pb.GetJobseekerDetailsResponse{
		Id:          uint64(result.Id),
		Email:       result.Email,
		Firstname:   result.Name,
		PhoneNumber: result.Phone,
		Blocked:     result.Blocked,
	}, nil
}

func (as *AdminServer) GetRecruiterDetails(ctx context.Context, Req *pb.GetRecruiterDetailsRequest) (*pb.GetRecruiterDetailsResponse, error) {
	as.Logger.Info("GetRecruiterDetails at AdminServer started")
	recruiterId := Req.RecruiterId
	as.Logger.Info("GetRecruiterDetails at adminUseCase started")

	result, err := as.adminUseCase.GetRecruiterDetails(int(recruiterId))
	if err != nil {
		as.Logger.Error("Error in GetRecruiterDetails at usecase: ", err)
		return &pb.GetRecruiterDetailsResponse{}, err
	}
	as.Logger.Info("GetRecruiterDetails at adminUseCase success")
	as.Logger.Info("GetRecruiterDetails at AdminServer success")
	return &pb.GetRecruiterDetailsResponse{
		Id:          uint64(result.Id),
		Email:       result.Contact_mail,
		CompanyName: result.CompanyName,
		PhoneNumber: result.Phone,
		Blocked:     result.Blocked,
	}, nil
}

func (as *AdminServer) CreatePolicy(ctx context.Context, Req *pb.CreatePolicyRequest) (*pb.CreatePolicyResponse, error) {
	as.Logger.Info("CreatePolicy at AdminServer started")
	var policy = req.CreatePolicyReq{
		Title:   Req.Title,
		Content: Req.Content,
	}
	as.Logger.Info("CreatePolicy at adminUseCase started")

	result, err := as.adminUseCase.CreatePolicy(policy)
	if err != nil {
		as.Logger.Error("Error in CreatePolicy at usecase: ", err)
		return &pb.CreatePolicyResponse{}, err
	}
	as.Logger.Info("CreatePolicy at adminUseCase success")
	as.Logger.Info("CreatePolicy at AdminServer success")
	return &pb.CreatePolicyResponse{
		Policy: &pb.Policy{
			Id:        int64(result.Policies.ID),
			Title:     result.Policies.Title,
			Content:   result.Policies.Content,
			CreatedAt: timestamppb.New(result.Policies.CreatedAt),
			UpdatedAt: timestamppb.New(result.Policies.UpdatedAt),
		},
	}, nil
}

func (as *AdminServer) UpdatePolicy(ctx context.Context, Req *pb.UpdatePolicyRequest) (*pb.UpdatePolicyResponse, error) {
	as.Logger.Info("UpdatePolicy at AdminServer started")
	var policy = req.UpdatePolicyReq{
		Id:      int(Req.Id),
		Title:   Req.Title,
		Content: Req.Content,
	}
	as.Logger.Info("UpdatePolicy at adminUseCase started")

	result, err := as.adminUseCase.UpdatePolicy(policy)
	if err != nil {
		as.Logger.Error("Error in UpdatePolicy at usecase: ", err)
		return &pb.UpdatePolicyResponse{}, err
	}
	as.Logger.Info("UpdatePolicy at adminUseCase success")
	as.Logger.Info("UpdatePolicy at AdminServer success")
	return &pb.UpdatePolicyResponse{
		Policy: &pb.Policy{
			Id:        int64(result.Policies.ID),
			Title:     result.Policies.Title,
			Content:   result.Policies.Content,
			CreatedAt: timestamppb.New(result.Policies.CreatedAt),
			UpdatedAt: timestamppb.New(result.Policies.UpdatedAt),
		},
	}, nil
}

func (as *AdminServer) DeletePolicy(ctx context.Context, Req *pb.DeletePolicyRequest) (*pb.DeletePolicyResponse, error) {

	as.Logger.Info("DeletePolicy at AdminServer started")

	var policy_id = Req.Id
	as.Logger.Info("DeletePolicy at adminUseCase started")

	result, err := as.adminUseCase.DeletePolicy(int(policy_id))
	if err != nil {
		as.Logger.Error("Error in DeletePolicy at usecase: ", err)
		return &pb.DeletePolicyResponse{}, err
	}
	as.Logger.Info("DeletePolicy at adminUseCase success")
	as.Logger.Info("DeletePolicy at AdminServer success")

	return &pb.DeletePolicyResponse{
		Deleted: result,
	}, nil
}

func (as *AdminServer) GetAllPolicies(ctx context.Context, Req *pb.GetAllPoliciesRequest) (*pb.GetAllPoliciesResponse, error) {
	as.Logger.Info("DeletePolicy at AdminServer started")
	as.Logger.Info("GetAllPolicies at adminUseCase started")

	result, err := as.adminUseCase.GetAllPolicies()
	if err != nil {
		as.Logger.Error("Error in GetAllPolicies at usecase: ", err)
		return &pb.GetAllPoliciesResponse{}, err
	}
	as.Logger.Info("GetAllPolicies at adminUseCase success")
	var policies []*pb.Policy

	for _, p := range result.Policies {
		policy := &pb.Policy{
			Id:        int64(p.ID),
			Title:     p.Title,
			Content:   p.Content,
			CreatedAt: timestamppb.New(p.CreatedAt),
			UpdatedAt: timestamppb.New(p.UpdatedAt),
		}

		policies = append(policies, policy)
	}
	
	as.Logger.Info("GetAllPolicies at AdminServer success")

	return &pb.GetAllPoliciesResponse{
		Policies: policies,
	}, nil
}

func (as *AdminServer) GetOnePolicy(ctx context.Context, Req *pb.GetOnePolicyRequest) (*pb.GetOnePolicyResponse, error) {
	as.Logger.Info("GetOnePolicy at AdminServer started")

	policy_id := Req.Id
	as.Logger.Info("GetOnePolicy at adminUseCase started")
	result, err := as.adminUseCase.GetOnePolicy(int(policy_id))
	if err != nil {
		as.Logger.Error("Error in GetOnePolicy at usecase: ", err)
		return &pb.GetOnePolicyResponse{}, err
	}
	as.Logger.Info("GetOnePolicy at adminUseCase success")
	as.Logger.Info("GetOnePolicy at AdminServer success")

	return &pb.GetOnePolicyResponse{
		Policy: &pb.Policy{
			Id:        int64(result.Policies.ID),
			Title:     result.Policies.Title,
			Content:   result.Policies.Content,
			CreatedAt: timestamppb.New(result.Policies.CreatedAt),
			UpdatedAt: timestamppb.New(result.Policies.UpdatedAt),
		},
	}, nil
}
