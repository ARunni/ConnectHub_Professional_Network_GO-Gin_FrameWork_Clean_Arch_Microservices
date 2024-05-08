package handler

import (
	interfaces "connectHub_gateway/pkg/client/interface"
	"connectHub_gateway/pkg/utils/models"
	"connectHub_gateway/pkg/utils/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type JobSeekerHandler struct {
	GRPC_Client interfaces.JobSeekerClient
}

func NewJobSeekerHandler(grpc_client interfaces.JobSeekerClient) *JobSeekerHandler {
	return &JobSeekerHandler{
		GRPC_Client: grpc_client,
	}
}

func (jh *JobSeekerHandler) JobSeekerSignup(c *gin.Context) {

	var jobseekerData models.JobSeekerSignUp

	if err := c.ShouldBindJSON(&jobseekerData); err != nil {
		errResp := response.ClientResponse(http.StatusBadRequest, "Incorrect Format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errResp)
		return
	}

	jobseeker, err := jh.GRPC_Client.JobSeekerSignup(jobseekerData)

	if err != nil {
		errResp := response.ClientResponse(http.StatusInternalServerError, "Signup failed Jobseeker", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errResp)
		return
	}

	successResp := response.ClientResponse(http.StatusOK, "Jobseeker Signup Successfully", jobseeker, nil)
	c.JSON(http.StatusOK, successResp)

}

func (jh *JobSeekerHandler) JobSeekerLogin(c *gin.Context) {

	var jobseekerData models.JobSeekerLogin

	if err := c.ShouldBindJSON(&jobseekerData); err != nil {
		errResp := response.ClientResponse(http.StatusBadRequest, "Incorrect Format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errResp)
		return
	}

	jobseeker, err := jh.GRPC_Client.JobSeekerLogin(jobseekerData)

	if err != nil {
		errResp := response.ClientResponse(http.StatusInternalServerError, "Cannot authenticate Jobseeker", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errResp)
		return
	}

	successResp := response.ClientResponse(http.StatusOK, "Jobseeker Authenticated Successfully", jobseeker, nil)
	c.JSON(http.StatusOK, successResp)

}
