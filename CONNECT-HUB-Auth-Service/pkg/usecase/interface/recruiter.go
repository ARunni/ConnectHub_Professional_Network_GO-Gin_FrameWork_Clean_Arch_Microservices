package interfaces

import req "ConnetHub_auth/pkg/utils/reqAndResponse"

type RecruiterUseCase interface {
	RecruiterSignup(recruiterdata req.RecruiterSignUp) (req.TokenRecruiter, error)
	RecruiterLogin(recruiterDetails req.RecruiterLogin) (req.TokenRecruiter, error)
}
