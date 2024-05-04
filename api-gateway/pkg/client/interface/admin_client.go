package interfaces

import "connectHub_gateway/pkg/utils/models"

type AdminClient interface {
	AdminLogin(admindata models.AdminLogin) (models.TokenAdmin, error)
}
