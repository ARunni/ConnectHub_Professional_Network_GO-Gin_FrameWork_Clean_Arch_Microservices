package interfaces

import req "ConnetHub_auth/pkg/utils/reqAndResponse"

type AdminUseCase interface {
	LoginHandler(adminDetails req.AdminLogin) (req.TokenAdmin, error)
}