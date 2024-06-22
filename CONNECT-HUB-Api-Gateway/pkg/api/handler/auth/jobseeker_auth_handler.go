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


// JobSeekerSignup handles the endpoint for job seeker signup.
// @Summary Job Seeker Signup
// @Description Allows a job seeker to sign up by providing necessary information.
// @Tags Jobseeker Authentication Management
// @Accept json
// @Produce json
// @Param jobseekerData body models.JobSeekerSignUp true "Job Seeker Sign Up Data"
// @Success 200 {object} response.Response "Jobseeker signup successful"
// @Failure 400 {object} response.Response "Incorrect data format"
// @Failure 500 {object} response.Response "Internal server error"
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

// JobSeekerLogin handles the endpoint for job seeker login.
// @Summary Job Seeker Login
// @Description Allows a job seeker to log in by providing necessary credentials.
// @Tags Jobseeker Authentication Management
// @Accept json
// @Produce json
// @Param jobseekerData body models.JobSeekerLogin true "Job Seeker Login Data"
// @Success 200 {object} response.Response "Jobseeker authenticated successfully"
// @Failure 400 {object} response.Response "Incorrect data format"
// @Failure 500 {object} response.Response "Internal server error"
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

// JobSeekerGetProfile handles the endpoint for retrieving a job seeker's profile.
// @Summary Get Job Seeker Profile
// @Description Retrieves the profile of the logged-in job seeker.
// @Tags Jobseeker Profile Management
// @Accept json
// @Produce json
// @Security BearerTokenAuth
// @Success 200 {object} response.Response "Profile retrieved successfully"
// @Failure 400 {object} response.Response "Incorrect data format or failed to retrieve profile"
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

// JobSeekerEditProfile handles the endpoint for editing a job seeker's profile.
// @Summary Edit Job Seeker Profile
// @Description Allows a job seeker to edit their profile by providing necessary information.
// @Tags Jobseeker Profile Management
// @Accept json
// @Produce json
// @Security BearerTokenAuth
// @Param jobseekerData body models.JobSeekerProfile true "Job Seeker Profile Data"
// @Success 200 {object} response.Response "Profile edited successfully"
// @Failure 400 {object} response.Response "Incorrect data format or failed to edit profile"
// @Router /jobseeker/profile [patch]
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

// GetAllPolicies handles the endpoint for retrieving all policies.
// @Summary Get All Policies
// @Description Retrieves a list of all policies.
// @Tags Jobseeker Policies Management
// @Accept json
// @Produce json
// @Security BearerTokenAuth
// @Success 200 {object} response.Response "Policies retrieved successfully"
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

// GetOnePolicy handles the endpoint for retrieving a specific policy.
// @Summary Get One Policy
// @Description Retrieves a specific policy by its ID.
// @Tags Jobseeker Policies Management
// @Accept json
// @Produce json
// @Security BearerTokenAuth
// @Param policy_id query int true "Policy ID"
// @Success 200 {object} response.Response "Policy retrieved successfully"
// @Failure 400 {object} response.Response "Failed to retrieve policy"
// @Router /jobseeker/policy [get]
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
