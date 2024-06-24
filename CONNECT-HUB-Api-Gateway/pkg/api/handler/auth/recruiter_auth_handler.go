package handler

import (
	logging "github.com/ARunni/connectHub_gateway/Logging"
	interfaces "github.com/ARunni/connectHub_gateway/pkg/client/auth/interface"
	"github.com/ARunni/connectHub_gateway/pkg/utils/models"
	"github.com/ARunni/connectHub_gateway/pkg/utils/response"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"

	msg "github.com/ARunni/Error_Message"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type RecruiterHandler struct {
	GRPC_Client interfaces.RecruiterAuthClient
	Logger      *logrus.Logger
	LogFile     *os.File
}

func NewRecruiterAuthHandler(grpc_client interfaces.RecruiterAuthClient) *RecruiterHandler {
	logger, logFile := logging.InitLogrusLogger("./Logging/connectHub_gateway.log")
	return &RecruiterHandler{
		GRPC_Client: grpc_client,
		Logger:      logger,
		LogFile:     logFile,
	}
}

// RecruiterSignup handles the signup operation for a recruiter.
// @Summary Recruiter signup
// @Description Register a new recruiter
// @Tags Recruiter Authentication Management
// @Accept json
// @Produce json
// @Param body body models.RecruiterSignUp true "Recruiter signup data"
// @Success 200 {object} response.Response "Recruiter signup successful"
// @Failure 400 {object} response.Response "Incorrect request format or missing required fields"
// @Failure 500 {object} response.Response "Internal server error: failed to signup recruiter"
// @Router /recruiter/signup [post]
func (jh *RecruiterHandler) RecruiterSignup(c *gin.Context) {

	var recruiterData models.RecruiterSignUp

	if err := c.ShouldBindJSON(&recruiterData); err != nil {
		jh.Logger.Error("Failed to Get Data: ", err)
		errResp := response.ClientResponse(http.StatusBadRequest, "Incorrect Format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errResp)
		return
	}

	recruiter, err := jh.GRPC_Client.RecruiterSignup(recruiterData)

	if err != nil {
		jh.Logger.Error("Failed to Recruiter Signup: ", err)
		errResp := response.ClientResponse(http.StatusInternalServerError, "Signup failed Recruiter", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errResp)
		return
	}

	jh.Logger.Info("Recruiter Signup Successful")

	successResp := response.ClientResponse(http.StatusOK, "Recruiter Signup Successfully", recruiter, nil)
	c.JSON(http.StatusOK, successResp)

}

// RecruiterLogin handles the login operation for a recruiter.
// @Summary Recruiter login
// @Description Authenticate a recruiter
// @Tags Recruiter Authentication Management
// @Accept json
// @Produce json
// @Param body body models.RecruiterLogin true "Recruiter credentials for login"
// @Success 200 {object} response.Response "Recruiter login successful"
// @Failure 400 {object} response.Response "Incorrect request format or missing required fields"
// @Failure 500 {object} response.Response "Internal server error: failed to authenticate recruiter"
// @Router /recruiter/login [post]
func (jh *RecruiterHandler) RecruiterLogin(c *gin.Context) {

	var recruiterData models.RecruiterLogin

	if err := c.ShouldBindJSON(&recruiterData); err != nil {
		jh.Logger.Error("Failed to Get Data: ", err)
		errResp := response.ClientResponse(http.StatusBadRequest, "Incorrect Format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errResp)
		return
	}

	recruiter, err := jh.GRPC_Client.RecruiterLogin(recruiterData)

	if err != nil {
		jh.Logger.Error("Failed to Recruiter Login: ", err)
		errResp := response.ClientResponse(http.StatusInternalServerError, "Cannot authenticate Recruiter", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errResp)
		return
	}

	jh.Logger.Info("Recruiter Login Successful")
	successResp := response.ClientResponse(http.StatusOK, "Recruiter Authenticated Successfully", recruiter, nil)
	c.JSON(http.StatusOK, successResp)

}

// RecruiterGetProfile handles the endpoint for retrieving a recruiter's profile.
// @Summary Get Recruiter Profile
// @Description Retrieves the profile of the logged-in recruiter.
// @Tags Recruiter Profile Management
// @Accept json
// @Produce json
// @Security BearerTokenAuth
// @Success 200 {object} response.Response "Profile retrieved successfully"
// @Failure 400 {object} response.Response "Failed to retrieve profile"
// @Router /recruiter/profile [get]
func (jh *RecruiterHandler) RecruiterGetProfile(c *gin.Context) {

	userIdstring, ok := c.Get("id")
	fmt.Println("status ", ok)
	if !ok {
		err := errors.New("error in getting id")
		jh.Logger.Error("Failed to Get Data: ", err)
		errResp := response.ClientResponse(http.StatusBadRequest, msg.MsgIdGetErr, nil, err.Error())
		c.JSON(http.StatusBadRequest, errResp)
		return
	}
	userId, strErr := userIdstring.(int)
	fmt.Println("recruiter id ", userId)
	fmt.Println("recruiter id ", userIdstring)
	if !strErr {
		jh.Logger.Error("Failed to Get Data: ")
		errResp := response.ClientResponse(http.StatusBadRequest, msg.MsgFormatErr, nil, strErr)
		c.JSON(http.StatusBadRequest, errResp)
		return
	}
	recruiter, err := jh.GRPC_Client.RecruiterGetProfile(userId)

	if err != nil {
		jh.Logger.Error("Failed to Recruiter Get Profile: ", err)
		errREsp := response.ClientResponse(http.StatusBadRequest, msg.MsgGettingDataErr, nil, err.Error())
		c.JSON(http.StatusBadRequest, errREsp)
		return
	}

	jh.Logger.Info("Recruiter Get Profile Successful")

	successResp := response.ClientResponse(http.StatusOK, msg.MsgGetSucces, recruiter, nil)
	c.JSON(http.StatusOK, successResp)

}

// RecruiterEditProfile handles the endpoint for editing a recruiter's profile.
// @Summary Edit Recruiter Profile
// @Description Allows a recruiter to edit their profile information.
// @Tags Recruiter Profile Management
// @Accept json
// @Produce json
// @Security BearerTokenAuth
// @Param id path int true "Recruiter ID"
// @Param recruiterData body models.RecruiterProfileReq true "Recruiter Profile Data"
// @Success 200 {object} response.Response "Recruiter profile updated successfully"
// @Failure 400 {object} response.Response "Failed to update profile"
// @Router /recruiter/profile [patch]
func (jh *RecruiterHandler) RecruiterEditProfile(c *gin.Context) {

	userIdstring, _ := c.Get("id")
	userId, strErr := userIdstring.(int)
	if !strErr {
		jh.Logger.Error("Failed to Get Data: ")
		errResp := response.ClientResponse(http.StatusBadRequest, msg.MsgFormatErr, nil, strErr)
		c.JSON(http.StatusBadRequest, errResp)
		return
	}

	var recruiterData models.RecruiterProfile
	recruiterData.ID = uint(userId)

	if err := c.ShouldBindJSON(&recruiterData); err != nil {
		jh.Logger.Error("Failed to Get Data: ", err)
		errResp := response.ClientResponse(http.StatusBadRequest, "Incorrect Format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errResp)
		return
	}

	recruiter, err := jh.GRPC_Client.RecruiterEditProfile(recruiterData)

	if err != nil {
		jh.Logger.Error("Failed to Recruiter Edit Profile: ", err)
		errREsp := response.ClientResponse(http.StatusBadRequest, msg.MsgGettingDataErr, nil, err.Error())
		c.JSON(http.StatusBadRequest, errREsp)
		return
	}
	jh.Logger.Info("Recruiter Edit Profile Successful")
	successResp := response.ClientResponse(http.StatusOK, msg.MsgGetSucces, recruiter, nil)
	c.JSON(http.StatusOK, successResp)

}

// GetAllPolicies handles the endpoint for retrieving all policies.
// @Summary Get All Policies
// @Description Retrieves all policies from the system.
// @Tags Recruiter Policy Management
// @Accept json
// @Produce json
// @Security BearerTokenAuth
// @Success 200 {object} response.Response "All policies retrieved successfully"
// @Failure 400 {object} response.Response "Failed to retrieve policies"
// @Router /recruiter/policies [get]
func (jh *RecruiterHandler) GetAllPolicies(c *gin.Context) {

	data, err := jh.GRPC_Client.GetAllPolicies()
	if err != nil {
		jh.Logger.Error("Failed to Get All Policies: ", err)

		errorRes := response.ClientResponse(http.StatusBadRequest, msg.MsgGettingDataErr, nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	jh.Logger.Info("Get All Policies Successful")

	successRes := response.ClientResponse(http.StatusOK, msg.MsgGetSucces, data, nil)
	c.JSON(http.StatusOK, successRes)

}

// GetOnePolicy handles the endpoint for retrieving a single policy by ID.
// @Summary Get One Policy
// @Description Retrieves a single policy based on the provided policy ID.
// @Tags Recruiter Policy Management
// @Accept json
// @Produce json
// @Security BearerTokenAuth
// @Param policy_id query integer true "Policy ID"
// @Success 200 {object} response.Response "Policy retrieved successfully"
// @Failure 400 {object} response.Response "Failed to retrieve policy or incorrect data format"
// @Router /recruiter/policy [get]
func (jh *RecruiterHandler) GetOnePolicy(c *gin.Context) {

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
		jh.Logger.Error("Failed to Get One Policy: ", err)
		errorRes := response.ClientResponse(http.StatusBadRequest, msg.MsgGettingDataErr, nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	jh.Logger.Info("Get One Policy Successful")

	successRes := response.ClientResponse(http.StatusOK, msg.MsgGetSucces, data, nil)
	c.JSON(http.StatusOK, successRes)

}
