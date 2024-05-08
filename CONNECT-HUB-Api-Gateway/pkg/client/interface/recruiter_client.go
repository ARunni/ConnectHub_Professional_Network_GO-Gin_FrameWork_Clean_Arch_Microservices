package interfaces

import "connectHub_gateway/pkg/utils/models"

type RecruiterClient interface {
	RecruiterSignup(recruiterData models.RecruiterSignUp) (models.TokenRecruiter, error)
	RecruiterLogin(recruiterData models.RecruiterLogin) (models.TokenRecruiter, error)
}
