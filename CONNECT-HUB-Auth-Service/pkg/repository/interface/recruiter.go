package interfaces

import req "ConnetHub_auth/pkg/utils/reqAndResponse"

type RecruiterRepository interface {
	RecruiterSignup(data req.RecruiterSignUp) (req.RecruiterDetailsResponse, error)
	RecruiterLogin(data req.RecruiterLogin) (req.RecruiterDetailsResponse, error)
	CheckRecruiterExistsByEmail(email string) (bool, error)
	CheckRecruiterBlockByEmail(email string) (bool, error)
}
