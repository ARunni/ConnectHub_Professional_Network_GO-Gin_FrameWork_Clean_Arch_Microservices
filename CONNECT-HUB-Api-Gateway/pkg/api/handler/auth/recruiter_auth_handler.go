package handler

import (
	logging "connectHub_gateway/Logging"
	interfaces "connectHub_gateway/pkg/client/auth/interface"
	"connectHub_gateway/pkg/utils/models"
	"connectHub_gateway/pkg/utils/response"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	msg "github.com/ARunni/Error_Message"
	"github.com/gin-gonic/gin"
)

type RecruiterHandler struct {
	GRPC_Client interfaces.RecruiterAuthClient
}

func NewRecruiterAuthHandler(grpc_client interfaces.RecruiterAuthClient) *RecruiterHandler {
	return &RecruiterHandler{
		GRPC_Client: grpc_client,
	}
}

// RecruiterSignup handles the signup operation for a recruiter.
// @Summary Recruiter signup
// @Description Register a new recruiter
// @Tags Recruiter
// @Accept json
// @Produce json
// @Param body body models.RecruiterSignUp true "Recruiter signup data"
// @Success 200 {object} response.Response "Recruiter signup successful"
// @Failure 400 {object} response.Response "Incorrect request format or missing required fields"
// @Failure 500 {object} response.Response "Internal server error: failed to signup recruiter"
// @Router /recruiter/signup [post]
func (jh *RecruiterHandler) RecruiterSignup(c *gin.Context) {

	logrusLogger, logrusLogFile := logging.InitLogrusLogger("./Logging/connectHub_gateway.log")
	defer logrusLogFile.Close()

	var recruiterData models.RecruiterSignUp

	if err := c.ShouldBindJSON(&recruiterData); err != nil {
		logrusLogger.Error("Failed to Get Data: ", err)
		errResp := response.ClientResponse(http.StatusBadRequest, "Incorrect Format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errResp)
		return
	}

	recruiter, err := jh.GRPC_Client.RecruiterSignup(recruiterData)

	if err != nil {
		logrusLogger.Error("Failed to Recruiter Signup: ", err)
		errResp := response.ClientResponse(http.StatusInternalServerError, "Signup failed Recruiter", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errResp)
		return
	}

	logrusLogger.Info("Recruiter Signup Successful")

	successResp := response.ClientResponse(http.StatusOK, "Recruiter Signup Successfully", recruiter, nil)
	c.JSON(http.StatusOK, successResp)

}

// RecruiterLogin handles the login operation for a recruiter.
// @Summary Recruiter login
// @Description Authenticate a recruiter
// @Tags Recruiter
// @Accept json
// @Produce json
// @Param body body models.RecruiterLogin true "Recruiter credentials for login"
// @Success 200 {object} response.Response "Recruiter login successful"
// @Failure 400 {object} response.Response "Incorrect request format or missing required fields"
// @Failure 500 {object} response.Response "Internal server error: failed to authenticate recruiter"
// @Router /recruiter/login [post]
func (jh *RecruiterHandler) RecruiterLogin(c *gin.Context) {

	logrusLogger, logrusLogFile := logging.InitLogrusLogger("./Logging/connectHub_gateway.log")
	defer logrusLogFile.Close()

	var recruiterData models.RecruiterLogin

	if err := c.ShouldBindJSON(&recruiterData); err != nil {
		logrusLogger.Error("Failed to Get Data: ", err)
		errResp := response.ClientResponse(http.StatusBadRequest, "Incorrect Format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errResp)
		return
	}

	recruiter, err := jh.GRPC_Client.RecruiterLogin(recruiterData)

	if err != nil {
		logrusLogger.Error("Failed to Recruiter Login: ", err)
		errResp := response.ClientResponse(http.StatusInternalServerError, "Cannot authenticate Recruiter", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errResp)
		return
	}

	logrusLogger.Info("Recruiter Login Successful")
	successResp := response.ClientResponse(http.StatusOK, "Recruiter Authenticated Successfully", recruiter, nil)
	c.JSON(http.StatusOK, successResp)

}

// RecruiterGetProfile retrieves the profile of a recruiter.
// @Summary Get recruiter profile
// @Description Retrieve the profile of a recruiter
// @Tags Recruiter
// @Accept json
// @Produce json
// @Security BearerTokenAuth
// @Success 200 {object} response.Response "Successfully retrieved recruiter profile"
// @Failure 400 {object} response.Response "Failed to retrieve recruiter profile"
// @Router /recruiter/profile [get]
func (jh *RecruiterHandler) RecruiterGetProfile(c *gin.Context) {

	logrusLogger, logrusLogFile := logging.InitLogrusLogger("./Logging/connectHub_gateway.log")
	defer logrusLogFile.Close()

	userIdstring, ok := c.Get("id")
	fmt.Println("status ", ok)
	if !ok {
		err := errors.New("error in getting id")
		logrusLogger.Error("Failed to Get Data: ", err)
		errResp := response.ClientResponse(http.StatusBadRequest, msg.MsgIdGetErr, nil, err.Error())
		c.JSON(http.StatusBadRequest, errResp)
		return
	}
	userId, strErr := userIdstring.(int)
	fmt.Println("recruiter id ", userId)
	fmt.Println("recruiter id ", userIdstring)
	if !strErr {
		logrusLogger.Error("Failed to Get Data: ")
		errResp := response.ClientResponse(http.StatusBadRequest, msg.MsgFormatErr, nil, strErr)
		c.JSON(http.StatusBadRequest, errResp)
		return
	}
	recruiter, err := jh.GRPC_Client.RecruiterGetProfile(userId)

	if err != nil {
		logrusLogger.Error("Failed to Recruiter Get Profile: ", err)
		errREsp := response.ClientResponse(http.StatusBadRequest, msg.MsgGettingDataErr, nil, err.Error())
		c.JSON(http.StatusBadRequest, errREsp)
		return
	}

	logrusLogger.Info("Recruiter Get Profile Successful")

	successResp := response.ClientResponse(http.StatusOK, msg.MsgGetSucces, recruiter, nil)
	c.JSON(http.StatusOK, successResp)

}

// RecruiterEditProfile handles the profile editing operation for a recruiter.
// @Summary Edit recruiter profile
// @Description Edit the profile of a recruiter
// @Tags Recruiter
// @Accept json
// @Produce json
// @Security BearerTokenAuth
// @Param body body models.RecruiterProfile true "Recruiter profile data for editing"
// @Success 200 {object} response.Response "Recruiter profile edited successfully"
// @Failure 400 {object} response.Response "Incorrect request format or missing required fields"
// @Failure 500 {object} response.Response "Internal server error: failed to edit recruiter profile"
// @Router /recruiter/profile [put]
func (jh *RecruiterHandler) RecruiterEditProfile(c *gin.Context) {

	logrusLogger, logrusLogFile := logging.InitLogrusLogger("./Logging/connectHub_gateway.log")
	defer logrusLogFile.Close()

	userIdstring, _ := c.Get("id")
	userId, strErr := userIdstring.(int)
	if !strErr {
		logrusLogger.Error("Failed to Get Data: ")
		errResp := response.ClientResponse(http.StatusBadRequest, msg.MsgFormatErr, nil, strErr)
		c.JSON(http.StatusBadRequest, errResp)
		return
	}

	var recruiterData models.RecruiterProfile
	recruiterData.ID = uint(userId)

	if err := c.ShouldBindJSON(&recruiterData); err != nil {
		logrusLogger.Error("Failed to Get Data: ", err)
		errResp := response.ClientResponse(http.StatusBadRequest, "Incorrect Format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errResp)
		return
	}

	recruiter, err := jh.GRPC_Client.RecruiterEditProfile(recruiterData)

	if err != nil {
		logrusLogger.Error("Failed to Recruiter Edit Profile: ", err)
		errREsp := response.ClientResponse(http.StatusBadRequest, msg.MsgGettingDataErr, nil, err.Error())
		c.JSON(http.StatusBadRequest, errREsp)
		return
	}
	logrusLogger.Info("Recruiter Edit Profile Successful")
	successResp := response.ClientResponse(http.StatusOK, msg.MsgGetSucces, recruiter, nil)
	c.JSON(http.StatusOK, successResp)

}

// GetAllPolicies retrieves all policies applicable to recruiters.
// @Summary Get all policies
// @Description Retrieve all policies applicable to recruiters
// @Tags Policy
// @Accept json
// @Produce json
// @Security BearerTokenAuth
// @Success 200 {object} response.Response "Successfully retrieved all policies"
// @Failure 400 {object} response.Response "Failed to retrieve policies"
// @Router /recruiter/policies [get]
func (jh *RecruiterHandler) GetAllPolicies(c *gin.Context) {

	logrusLogger, logrusLogFile := logging.InitLogrusLogger("./Logging/connectHub_gateway.log")
	defer logrusLogFile.Close()

	data, err := jh.GRPC_Client.GetAllPolicies()
	if err != nil {
		logrusLogger.Error("Failed to Get All Policies: ", err)

		errorRes := response.ClientResponse(http.StatusBadRequest, msg.MsgGettingDataErr, nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	logrusLogger.Info("Get All Policies Successful")

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
// @Router /recruiter/policies/{policy_id} [get]
func (jh *RecruiterHandler) GetOnePolicy(c *gin.Context) {

	logrusLogger, logrusLogFile := logging.InitLogrusLogger("./Logging/connectHub_gateway.log")
	defer logrusLogFile.Close()

	idStr := c.Query("policy_id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		logrusLogger.Error("Failed to Get Data: ", err)
		errorRes := response.ClientResponse(http.StatusBadRequest, msg.MsgIdGetErr, nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	data, err := jh.GRPC_Client.GetOnePolicy(id)

	if err != nil {
		logrusLogger.Error("Failed to Get One Policy: ", err)
		errorRes := response.ClientResponse(http.StatusBadRequest, msg.MsgGettingDataErr, nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	logrusLogger.Info("Get One Policy Successful")

	successRes := response.ClientResponse(http.StatusOK, msg.MsgGetSucces, data, nil)
	c.JSON(http.StatusOK, successRes)

}
