package handler

import (
	interfaces "github.com/ARunni/connectHub_gateway/pkg/client/auth/interface"
	"github.com/ARunni/connectHub_gateway/pkg/utils/models"
	"github.com/ARunni/connectHub_gateway/pkg/utils/response"
	"errors"
	"net/http"
	"os"
	"strconv"

	logging "github.com/ARunni/connectHub_gateway/Logging"

	msg "github.com/ARunni/Error_Message"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type JobSeekerHandler struct {
	GRPC_Client interfaces.JobSeekerAuthClient
	Logger      *logrus.Logger
	LogFile     *os.File
}

func NewJobSeekerAuthHandler(grpc_client interfaces.JobSeekerAuthClient) *JobSeekerHandler {

	logger, logFile := logging.InitLogrusLogger("./Logging/connectHub_gateway.log")
	return &JobSeekerHandler{
		GRPC_Client: grpc_client,
		Logger:      logger,
		LogFile:     logFile,
	}
}

// JobSeekerSignup handles the signup operation for a job seeker.
// @Summary Job seeker signup
// @Description Register a new job seeker
// @Tags Job Seeker
// @Accept json
// @Produce json
// @Param body body models.JobSeekerSignUp true "Job seeker data for signup"
// @Success 200 {object} response.Response "Job seeker signup successful"
// @Failure 400 {object} response.Response "Incorrect request format or missing required fields"
// @Failure 500 {object} response.Response "Internal server error: failed to signup job seeker"
// @Router /jobseeker/signup [post]
func (jh *JobSeekerHandler) JobSeekerSignup(c *gin.Context) {

	var jobseekerData models.JobSeekerSignUp

	if err := c.ShouldBindJSON(&jobseekerData); err != nil {
		jh.Logger.Error("Failed to Get Data: ", err)
		errResp := response.ClientResponse(http.StatusBadRequest, "Incorrect Format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errResp)
		return
	}

	jobseeker, err := jh.GRPC_Client.JobSeekerSignup(jobseekerData)

	if err != nil {
		jh.Logger.Error("Failed to Signup: ", err)
		errResp := response.ClientResponse(http.StatusInternalServerError, "Signup failed Jobseeker", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errResp)
		return
	}

	jh.Logger.Info("Jobseeker Signup Successful")

	successResp := response.ClientResponse(http.StatusOK, "Jobseeker Signup Successfully", jobseeker, nil)
	c.JSON(http.StatusOK, successResp)

}

// JobSeekerLogin handles the login operation for a job seeker.
// @Summary Job seeker login
// @Description Authenticate a job seeker
// @Tags Job Seeker
// @Accept json
// @Produce json
// @Param body body models.JobSeekerLogin true "Job seeker credentials for login"
// @Success 200 {object} response.Response "Job seeker login successful"
// @Failure 400 {object} response.Response "Incorrect request format or missing required fields"
// @Failure 500 {object} response.Response "Internal server error: failed to authenticate job seeker"
// @Router /jobseeker/login [post]
func (jh *JobSeekerHandler) JobSeekerLogin(c *gin.Context) {

	var jobseekerData models.JobSeekerLogin

	if err := c.ShouldBindJSON(&jobseekerData); err != nil {
		jh.Logger.Error("Failed to Get Data: ", err)
		errResp := response.ClientResponse(http.StatusBadRequest, "Incorrect Format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errResp)
		return
	}

	jobseeker, err := jh.GRPC_Client.JobSeekerLogin(jobseekerData)

	if err != nil {
		jh.Logger.Error("Failed to Signin: ", err)
		errResp := response.ClientResponse(http.StatusInternalServerError, "Cannot authenticate Jobseeker", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errResp)
		return
	}

	jh.Logger.Info("Jobseeker Signin Successful")

	successResp := response.ClientResponse(http.StatusOK, "Jobseeker Authenticated Successfully", jobseeker, nil)
	c.JSON(http.StatusOK, successResp)

}

// JobSeekerGetProfile retrieves the profile of a job seeker.
// @Summary Get job seeker profile
// @Description Retrieve the profile of a job seeker
// @Tags Job Seeker
// @Accept json
// @Produce json
// @Security BearerTokenAuth
// @Success 200 {object} response.Response "Successfully retrieved job seeker profile"
// @Failure 400 {object} response.Response "Failed to retrieve job seeker profile"
// @Router /jobseeker/profile [get]
func (jh *JobSeekerHandler) JobSeekerGetProfile(c *gin.Context) {

	userIdstring, _ := c.Get("id")
	userId, strErr := userIdstring.(int)
	if !strErr {
		jh.Logger.Error("Failed to Get Data: ", errors.New("getting id failed"))
		errResp := response.ClientResponse(http.StatusBadRequest, msg.MsgFormatErr, nil, strErr)
		c.JSON(http.StatusBadRequest, errResp)
		return
	}

	jobseeker, err := jh.GRPC_Client.JobSeekerGetProfile(userId)

	if err != nil {
		jh.Logger.Error("Failed to Get Profile: ", err)
		errREsp := response.ClientResponse(http.StatusBadRequest, msg.MsgGettingDataErr, nil, err.Error())
		c.JSON(http.StatusBadRequest, errREsp)
		return
	}

	jh.Logger.Info("Getting profile Successful")

	successResp := response.ClientResponse(http.StatusOK, msg.MsgGetSucces, jobseeker, nil)
	c.JSON(http.StatusOK, successResp)

}

// JobSeekerEditProfile handles the profile editing operation for a job seeker.
// @Summary Edit job seeker profile
// @Description Edit the profile of a job seeker
// @Tags Job Seeker
// @Accept json
// @Produce json
// @Security BearerTokenAuth
// @Param body body models.JobSeekerProfile true "Job seeker profile data for editing"
// @Success 200 {object} response.Response "Job seeker profile edited successfully"
// @Failure 400 {object} response.Response "Incorrect request format or missing required fields"
// @Failure 500 {object} response.Response "Internal server error: failed to edit job seeker profile"
// @Router /jobseeker/profile [put]
func (jh *JobSeekerHandler) JobSeekerEditProfile(c *gin.Context) {

	userIdstring, _ := c.Get("id")
	userId, strErr := userIdstring.(int)

	if !strErr {
		jh.Logger.Error("Failed to Get Data: ", errors.New("getting id failed"))
		errResp := response.ClientResponse(http.StatusBadRequest, msg.MsgFormatErr, nil, strErr)
		c.JSON(http.StatusBadRequest, errResp)
		return
	}

	var jobseekerData models.JobSeekerProfile
	jobseekerData.ID = uint(userId)

	if err := c.ShouldBindJSON(&jobseekerData); err != nil {
		jh.Logger.Error("Failed to Get Data: ", err)
		errResp := response.ClientResponse(http.StatusBadRequest, "Incorrect Format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errResp)
		return
	}

	jobseeker, err := jh.GRPC_Client.JobSeekerEditProfile(jobseekerData)

	if err != nil {
		jh.Logger.Error("Failed to edit profile: ", err)
		errREsp := response.ClientResponse(http.StatusBadRequest, msg.MsgGettingDataErr, nil, err.Error())
		c.JSON(http.StatusBadRequest, errREsp)
		return
	}
	jh.Logger.Info("Jobseeker edit profile Successful")

	successResp := response.ClientResponse(http.StatusOK, msg.MsgGetSucces, jobseeker, nil)
	c.JSON(http.StatusOK, successResp)

}

// GetAllPolicies retrieves all policies applicable to job seekers.
// @Summary Get all policies
// @Description Retrieve all policies applicable to job seekers
// @Tags Policy
// @Accept json
// @Produce json
// @Security BearerTokenAuth
// @Success 200 {object} response.Response "Successfully retrieved all policies"
// @Failure 400 {object} response.Response "Failed to retrieve policies"
// @Router /jobseeker/policies [get]
func (jh *JobSeekerHandler) GetAllPolicies(c *gin.Context) {

	data, err := jh.GRPC_Client.GetAllPolicies()

	if err != nil {
		jh.Logger.Error("Failed to get all policies: ", err)
		errorRes := response.ClientResponse(http.StatusBadRequest, msg.MsgGettingDataErr, nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	jh.Logger.Info("Jobseeker get all policies Successful")

	successRes := response.ClientResponse(http.StatusOK, msg.MsgGetSucces, data, nil)
	c.JSON(http.StatusOK, successRes)

}

// GetOnePolicy retrieves details of a specific policy based on its ID.
// @Summary Get one policy
// @Description Retrieve details of a specific policy based on its ID
// @Tags Policy
// @Accept json
// @Produce json
// @Security BearerTokenAuth
// @Param policy_id query int true "Policy ID to retrieve"
// @Success 200 {object} response.Response "Successfully retrieved the policy details"
// @Failure 400 {object} response.Response "Failed to retrieve policy details"
// @Router /jobseeker/policies/{policy_id} [get]
func (jh *JobSeekerHandler) GetOnePolicy(c *gin.Context) {

	idStr := c.Query("policy_id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		jh.Logger.Error("Failed to Get Data: ", err)
		errorRes := response.ClientResponse(http.StatusBadRequest, msg.MsgIdGetErr, nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	data, err := jh.GRPC_Client.GetOnePolicy(id)

	if err != nil {
		jh.Logger.Error("Failed to get one policy: ", err)
		errorRes := response.ClientResponse(http.StatusBadRequest, msg.MsgGettingDataErr, nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	jh.Logger.Info("Jobseeker getting one policy Successful")

	successRes := response.ClientResponse(http.StatusOK, msg.MsgGetSucces, data, nil)
	c.JSON(http.StatusOK, successRes)
}
