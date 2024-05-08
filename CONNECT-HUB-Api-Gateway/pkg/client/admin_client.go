package client

import (
	interfaces "connectHub_gateway/pkg/client/interface"
	"connectHub_gateway/pkg/config"
	pb "connectHub_gateway/pkg/pb/auth/admin"
	"connectHub_gateway/pkg/utils/models"
	"context"
	"fmt"

	"google.golang.org/grpc"
)

type adminClient struct {
	Client pb.AdminClient
}

func NewAdminClient(cfg config.Config) interfaces.AdminClient {
	fmt.Println("admin client")
	grpcConnection, err := grpc.Dial(cfg.ConnetHubAuth, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not Connect to Auth", err)
	}
	grpcClient := pb.NewAdminClient(grpcConnection)
	return &adminClient{
		Client: grpcClient,
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
