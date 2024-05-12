package interfaces

import req "ConnetHub_auth/pkg/utils/reqAndResponse"

type RecruiterUseCase interface {
	RecruiterSignup(recruiterdata req.RecruiterSignUp) (req.TokenRecruiter, error)
	RecruiterLogin(recruiterDetails req.RecruiterLogin) (req.TokenRecruiter, error)
	RecruiterGetProfile(id int) (req.RecruiterProfile, error)
	RecruiterEditProfile(profile req.RecruiterProfile) (req.RecruiterProfile, error)
}
