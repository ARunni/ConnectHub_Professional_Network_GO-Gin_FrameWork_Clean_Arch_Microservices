package handler

import (
	interfaces "connectHub_gateway/pkg/client/auth/interface"
	"connectHub_gateway/pkg/utils/models"
	"connectHub_gateway/pkg/utils/response"
	"net/http"
	"strconv"

	msg "github.com/ARunni/Error_Message"
	"github.com/gin-gonic/gin"
)

type JobSeekerHandler struct {
	GRPC_Client interfaces.JobSeekerAuthClient
}

func NewJobSeekerAuthHandler(grpc_client interfaces.JobSeekerAuthClient) *JobSeekerHandler {
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

func (jh *JobSeekerHandler) JobSeekerGetProfile(c *gin.Context) {

	userIdstring, _ := c.Get("id")
	userId, strErr := userIdstring.(int)
	if !strErr {
		errResp := response.ClientResponse(http.StatusBadRequest, msg.MsgFormatErr, nil, strErr)
		c.JSON(http.StatusBadRequest, errResp)
		return
	}
	jobseeker, err := jh.GRPC_Client.JobSeekerGetProfile(userId)
	if err != nil {
		errREsp := response.ClientResponse(http.StatusBadRequest, msg.MsgGettingDataErr, nil, err.Error())
		c.JSON(http.StatusBadRequest, errREsp)
		return
	}
	successResp := response.ClientResponse(http.StatusOK, msg.MsgGetSucces, jobseeker, nil)
	c.JSON(http.StatusOK, successResp)

}

func (jh *JobSeekerHandler) JobSeekerEditProfile(c *gin.Context) {

	userIdstring, _ := c.Get("id")
	userId, strErr := userIdstring.(int)
	if !strErr {
		errResp := response.ClientResponse(http.StatusBadRequest, msg.MsgFormatErr, nil, strErr)
		c.JSON(http.StatusBadRequest, errResp)
		return
	}

	var jobseekerData models.JobSeekerProfile
	jobseekerData.ID = uint(userId)

	if err := c.ShouldBindJSON(&jobseekerData); err != nil {
		errResp := response.ClientResponse(http.StatusBadRequest, "Incorrect Format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errResp)
		return
	}

	jobseeker, err := jh.GRPC_Client.JobSeekerEditProfile(jobseekerData)
	if err != nil {
		errREsp := response.ClientResponse(http.StatusBadRequest, msg.MsgGettingDataErr, nil, err.Error())
		c.JSON(http.StatusBadRequest, errREsp)
		return
	}
	successResp := response.ClientResponse(http.StatusOK, msg.MsgGetSucces, jobseeker, nil)
	c.JSON(http.StatusOK, successResp)

}

func (jh *JobSeekerHandler) GetAllPolicies(c *gin.Context) {

	data, err := jh.GRPC_Client.GetAllPolicies()
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, msg.MsgGettingDataErr, nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, msg.MsgGetSucces, data, nil)
	c.JSON(http.StatusOK, successRes)

}

func (jh *JobSeekerHandler) GetOnePolicy(c *gin.Context) {

	idStr := c.Query("policy_id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, msg.MsgIdGetErr, nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	data, err := jh.GRPC_Client.GetOnePolicy(id)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, msg.MsgGettingDataErr, nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, msg.MsgGetSucces, data, nil)
	c.JSON(http.StatusOK, successRes)
}
