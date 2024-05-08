package interfaces

import req "ConnetHub_auth/pkg/utils/reqAndResponse"

type AdminUseCase interface {
	AdminLogin(adminDetails req.AdminLogin) (req.TokenAdmin, error)
}