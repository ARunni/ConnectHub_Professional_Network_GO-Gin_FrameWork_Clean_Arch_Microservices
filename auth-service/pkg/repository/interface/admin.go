package interfaces

import req "ConnetHub_auth/pkg/utils/reqAndResponse"

type AdminRepository interface {
	AdminLogin(admin req.AdminLogin) (req.AdminDetailsResponse, error)
	CheckAdminExistsByEmail(email string) (bool, error)
}
