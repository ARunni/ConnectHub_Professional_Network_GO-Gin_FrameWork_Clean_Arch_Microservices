package handler

import (
	interfaces "connectHub_gateway/pkg/client/interface"
	"connectHub_gateway/pkg/utils/models"
	"connectHub_gateway/pkg/utils/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RecruiterHandler struct {
	GRPC_Client interfaces.RecruiterClient
}

func NewRecruiterHandler(grpc_client interfaces.RecruiterClient) *RecruiterHandler {
	return &RecruiterHandler{
		GRPC_Client: grpc_client,
	}
}

func (jh *RecruiterHandler) RecruiterSignup(c *gin.Context) {

	var recruiterData models.RecruiterSignUp

	if err := c.ShouldBindJSON(&recruiterData); err != nil {
		errResp := response.ClientResponse(http.StatusBadRequest, "Incorrect Format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errResp)
		return
	}

	recruiter, err := jh.GRPC_Client.RecruiterSignup(recruiterData)

	if err != nil {
		errResp := response.ClientResponse(http.StatusInternalServerError, "Cannot authenticate Admin", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errResp)
		return
	}

	successResp := response.ClientResponse(http.StatusOK, "Admin Authenticated Successfully", recruiter, nil)
	c.JSON(http.StatusOK, successResp)

}
func (jh *RecruiterHandler) RecruiterLogin(c *gin.Context) {

	var recruiterData models.RecruiterLogin

	if err := c.ShouldBindJSON(&recruiterData); err != nil {
		errResp := response.ClientResponse(http.StatusBadRequest, "Incorrect Format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errResp)
		return
	}

	recruiter, err := jh.GRPC_Client.RecruiterLogin(recruiterData)

	if err != nil {
		errResp := response.ClientResponse(http.StatusInternalServerError, "Cannot authenticate Admin", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errResp)
		return
	}

	successResp := response.ClientResponse(http.StatusOK, "Admin Authenticated Successfully", recruiter, nil)
	c.JSON(http.StatusOK, successResp)

}
