package handler

import (
	"net/http"
	"os"
	"strconv"

	logging "github.com/ARunni/connectHub_gateway/Logging"
	interfaces "github.com/ARunni/connectHub_gateway/pkg/client/auth/interface"
	"github.com/ARunni/connectHub_gateway/pkg/utils/models"
	"github.com/ARunni/connectHub_gateway/pkg/utils/response"

	msg "github.com/ARunni/Error_Message"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type AdminHandler struct {
	GRPC_Client interfaces.AdminAuthClient
	Logger      *logrus.Logger
	LogFile     *os.File
}

func NewAdminAuthHandler(grpc_client interfaces.AdminAuthClient) *AdminHandler {
	logger, logFile := logging.InitLogrusLogger("./Logging/connectHub_gateway.log")
	return &AdminHandler{
		GRPC_Client: grpc_client,
		Logger:      logger,
		LogFile:     logFile,
	}
}

// LoginHandler handles the login operation for an admin.
// @Summary Admin login
// @Description Authenticate an admin and get access token
// @Tags Admin Authentication Management
// @Accept json
// @Produce json
// @Param body body models.AdminLogin true "Admin credentials for login"
// @Success 200 {object} response.Response "Admin login successful"
// @Failure 400 {object} response.Response "Invalid request or constraints not satisfied"
// @Failure 401 {object} response.Response "Unauthorized: cannot authenticate user"
// @Router /admin/login [post]
func (ah *AdminHandler) AdminLogin(c *gin.Context) {

	var adminData models.AdminLogin

	if err := c.ShouldBindJSON(&adminData); err != nil {
		ah.Logger.Error("Failed to Get Data: ", err)
		errResp := response.ClientResponse(http.StatusBadRequest, msg.ErrFormat, nil, err.Error())
		c.JSON(http.StatusBadRequest, errResp)
		return
	}
	admin, err := ah.GRPC_Client.AdminLogin(adminData)
	if err != nil {
		ah.Logger.Error("Failed to login admin: ", err)
		errResp := response.ClientResponse(http.StatusInternalServerError, "Cannot authenticate Admin", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errResp)
		return
	}
	ah.Logger.Info("Admin signin Successful")
	successResp := response.ClientResponse(http.StatusOK, "Admin Authenticated Successfully", admin, nil)
	c.JSON(http.StatusOK, successResp)

}

// GetJobseekers handles fetching a paginated list of jobseekers.
// @Summary Get jobseekers
// @Description Retrieve a paginated list of jobseekers
// @Tags Admin Jobseeker Management
// @Accept json
// @Produce json
// @Security BearerTokenAuth
// @Param page query int true "Page number for pagination"
// @Success 200 {object} response.Response "Successfully retrieved jobseekers"
// @Failure 400 {object} response.Response "Invalid page number or constraints not satisfied"
// @Failure 500 {object} response.Response "Failed to retrieve jobseekers due to internal error"
// @Router /admin/jobseeker/all [get]
func (ah *AdminHandler) GetJobseekers(c *gin.Context) {

	pageStr := c.Query("page")
	page, err := strconv.Atoi(pageStr)

	if err != nil {
		ah.Logger.Error("Failed to Get Data: ", err)
		errorRes := response.ClientResponse(http.StatusBadRequest, msg.MsgPageNumFormatErr, nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	jobseeker, err := ah.GRPC_Client.GetJobseekers(page)
	if err != nil {
		ah.Logger.Error("Failed to Get Jobseekers: ", err)
		errorRes := response.ClientResponse(http.StatusBadRequest, msg.MsgGettingDataErr, nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	ah.Logger.Info("Get Jobseekers Successful")
	successRes := response.ClientResponse(http.StatusOK, msg.MsgGetSucces, jobseeker, nil)
	c.JSON(http.StatusOK, successRes)
}

// GetRecruiters handles fetching a paginated list of recruiters.
// @Summary Get recruiters
// @Description Retrieve a paginated list of recruiters
// @Tags Admin Recruiter Management
// @Accept json
// @Produce json
// @Security BearerTokenAuth
// @Param page query int true "Page number for pagination"
// @Success 200 {object} response.Response "Successfully retrieved recruiters"
// @Failure 400 {object} response.Response "Invalid page number or constraints not satisfied"
// @Failure 500 {object} response.Response "Failed to retrieve recruiters due to internal error"
// @Router /admin/recruiter/all [get]
func (ah *AdminHandler) GetRecruiters(c *gin.Context) {

	pageStr := c.Query("page")
	page, err := strconv.Atoi(pageStr)

	if err != nil {
		ah.Logger.Error("Failed to Get Data: ", err)
		errorRes := response.ClientResponse(http.StatusBadRequest, msg.MsgPageNumFormatErr, nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	recruiter, err := ah.GRPC_Client.GetRecruiters(page)
	if err != nil {
		ah.Logger.Error("Failed to Get Recruiters: ", err)
		errorRes := response.ClientResponse(http.StatusBadRequest, msg.MsgGettingDataErr, nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	ah.Logger.Info("Get Recruiters Successful")
	successRes := response.ClientResponse(http.StatusOK, msg.MsgGetSucces, recruiter, nil)
	c.JSON(http.StatusOK, successRes)
}

// BlockRecruiter blocks a recruiter by ID.
// @Summary Block a recruiter
// @Description Block a recruiter based on the provided ID
// @Tags Admin Recruiter Management
// @Accept json
// @Produce json
// @Security BearerTokenAuth
// @Param id query int true "Recruiter ID to block"
// @Success 200 {object} response.Response "Recruiter blocked successfully"
// @Failure 400 {object} response.Response "Invalid recruiter ID or failed to block recruiter"
// @Router /admin/recruiter/block [patch]
func (ah *AdminHandler) BlockRecruiter(c *gin.Context) {

	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		ah.Logger.Error("Failed to Get Data: ", err)
		errorRes := response.ClientResponse(http.StatusBadRequest, msg.MsgPageNumFormatErr, nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	blockRecruiter, err := ah.GRPC_Client.BlockRecruiter(id)

	if err != nil {
		ah.Logger.Error("Failed to Block Recruiter: ", err)
		errorRes := response.ClientResponse(http.StatusBadRequest, msg.MsgGettingDataErr, nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	ah.Logger.Info("Block Recruiter Successful")

	successRes := response.ClientResponse(http.StatusOK, msg.ErrUserBlockTrue, blockRecruiter, nil)
	c.JSON(http.StatusOK, successRes)
}

// BlockJobseeker blocks a jobseeker by ID.
// @Summary Block a jobseeker
// @Description Block a jobseeker based on the provided ID
// @Tags Admin Jobseeker Management
// @Accept json
// @Produce json
// @Security BearerTokenAuth
// @Param id query int true "Jobseeker ID to block"
// @Success 200 {object} response.Response "Jobseeker blocked successfully"
// @Failure 400 {object} response.Response "Invalid jobseeker ID or failed to block jobseeker"
// @Router /admin/jobseeker/block [patch]
func (ah *AdminHandler) BlockJobseeker(c *gin.Context) {

	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		ah.Logger.Error("Failed to Get Data: ", err)
		errorRes := response.ClientResponse(http.StatusBadRequest, msg.MsgPageNumFormatErr, nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	blockJobseeker, err := ah.GRPC_Client.BlockJobseeker(id)

	if err != nil {
		ah.Logger.Error("Failed to BlockJobseeker: ", err)
		errorRes := response.ClientResponse(http.StatusBadRequest, msg.MsgGettingDataErr, nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	ah.Logger.Info("BlockJobseeker is Successful")

	successRes := response.ClientResponse(http.StatusOK, msg.ErrUserBlockTrue, blockJobseeker, nil)
	c.JSON(http.StatusOK, successRes)
}

// UnBlockJobseeker unblocks a jobseeker by ID.
// @Summary Unblock a jobseeker
// @Description Unblock a jobseeker based on the provided ID
// @Tags Admin Jobseeker Management
// @Accept json
// @Produce json
// @Security BearerTokenAuth
// @Param id query int true "Jobseeker ID to unblock"
// @Success 200 {object} response.Response "Jobseeker unblocked successfully"
// @Failure 400 {object} response.Response "Invalid jobseeker ID or failed to unblock jobseeker"
// @Router /admin/jobseeker/unblock [patch]
func (ah *AdminHandler) UnBlockJobseeker(c *gin.Context) {

	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		ah.Logger.Error("Failed to Get Data: ", err)
		errorRes := response.ClientResponse(http.StatusBadRequest, msg.MsgPageNumFormatErr, nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	unBlockjobseeker, err := ah.GRPC_Client.UnBlockJobseeker(id)

	if err != nil {
		ah.Logger.Error("Failed to UnBlockJobseeker: ", err)
		errorRes := response.ClientResponse(http.StatusBadRequest, msg.MsgGettingDataErr, nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	ah.Logger.Info("UnBlock Jobseeker Successful")

	successRes := response.ClientResponse(http.StatusOK, msg.MsgGetSucces, unBlockjobseeker, nil)
	c.JSON(http.StatusOK, successRes)
}

// UnBlockRecruiter unblocks a recruiter by ID.
// @Summary Unblock a recruiter
// @Description Unblock a recruiter based on the provided ID
// @Tags Admin Recruiter Management
// @Accept json
// @Produce json
// @Security BearerTokenAuth
// @Param id query int true "Recruiter ID to unblock"
// @Success 200 {object} response.Response "Recruiter unblocked successfully"
// @Failure 400 {object} response.Response "Invalid recruiter ID or failed to unblock recruiter"
// @Router /admin/recruiter/unblock [patch]
func (ah *AdminHandler) UnBlockRecruiter(c *gin.Context) {

	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		ah.Logger.Error("Failed to Get Data: ", err)
		errorRes := response.ClientResponse(http.StatusBadRequest, msg.MsgPageNumFormatErr, nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	unBlockRecruiter, err := ah.GRPC_Client.UnBlockRecruiter(id)

	if err != nil {
		ah.Logger.Error("Failed to UnBlock Recruiter: ", err)
		errorRes := response.ClientResponse(http.StatusBadRequest, msg.MsgGettingDataErr, nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	ah.Logger.Info("UnBlock Recruiter Successful")

	successRes := response.ClientResponse(http.StatusOK, msg.MsgGetSucces, unBlockRecruiter, nil)
	c.JSON(http.StatusOK, successRes)
}

// GetJobseekerDetails retrieves the details of a jobseeker by ID.
// @Summary Get jobseeker details
// @Description Retrieve the details of a jobseeker based on the provided ID
// @Tags Admin Jobseeker Management
// @Accept json
// @Produce json
// @Security BearerTokenAuth
// @Param id query int true "Jobseeker ID to get details"
// @Success 200 {object} response.Response "Jobseeker details retrieved successfully"
// @Failure 400 {object} response.Response "Invalid jobseeker ID or failed to retrieve jobseeker details"
// @Router /admin/jobseeker [get]
func (ah *AdminHandler) GetJobseekerDetails(c *gin.Context) {

	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		ah.Logger.Error("Failed to Get Data: ", err)
		errorRes := response.ClientResponse(http.StatusBadRequest, msg.MsgPageNumFormatErr, nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	unBlockRecruiter, err := ah.GRPC_Client.GetJobseekerDetails(id)

	if err != nil {
		ah.Logger.Error("Failed to Get Jobseeker Details: ", err)
		errorRes := response.ClientResponse(http.StatusBadRequest, msg.MsgGettingDataErr, nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	ah.Logger.Info("Get Jobseeker Details Successful")

	successRes := response.ClientResponse(http.StatusOK, msg.MsgGetSucces, unBlockRecruiter, nil)
	c.JSON(http.StatusOK, successRes)
}

// GetRecruiterDetails retrieves the details of a recruiter by ID.
// @Summary Get recruiter details
// @Description Retrieve the details of a recruiter based on the provided ID
// @Tags Admin Recruiter Management
// @Accept json
// @Produce json
// @Security BearerTokenAuth
// @Param id query int true "Recruiter ID to get details"
// @Success 200 {object} response.Response "Recruiter details retrieved successfully"
// @Failure 400 {object} response.Response "Invalid recruiter ID or failed to retrieve recruiter details"
// @Router /admin/recruiter [get]
func (ah *AdminHandler) GetRecruiterDetails(c *gin.Context) {

	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		ah.Logger.Error("Failed to Get Data: ", err)
		errorRes := response.ClientResponse(http.StatusBadRequest, msg.MsgPageNumFormatErr, nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	unBlockRecruiter, err := ah.GRPC_Client.GetRecruiterDetails(id)

	if err != nil {
		ah.Logger.Error("Failed to Get Recruiter Details: ", err)
		errorRes := response.ClientResponse(http.StatusBadRequest, msg.MsgGettingDataErr, nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	ah.Logger.Info("Get Recruiter Details Successful")

	successRes := response.ClientResponse(http.StatusOK, msg.MsgGetSucces, unBlockRecruiter, nil)
	c.JSON(http.StatusOK, successRes)
}

// policies

// CreatePolicy creates a new policy.
// @Summary Create a new policy
// @Description Create a new policy based on the provided data
// @Tags Admin Policy Management
// @Accept json
// @Produce json
// @Security BearerTokenAuth
// @Param policy body models.CreatePolicyReq true "Policy data"
// @Success 200 {object} response.Response "Policy created successfully"
// @Failure 400 {object} response.Response "Invalid policy data or failed to create policy"
// @Router /admin/policy [post]
func (ah *AdminHandler) CreatePolicy(c *gin.Context) {

	var policyData models.CreatePolicyReq

	if err := c.ShouldBindJSON(&policyData); err != nil {
		ah.Logger.Error("Failed to Get Data: ", err)
		errResp := response.ClientResponse(http.StatusBadRequest, msg.ErrFormat, nil, err.Error())
		c.JSON(http.StatusBadRequest, errResp)
		return
	}

	data, err := ah.GRPC_Client.CreatePolicy(policyData)

	if err != nil {
		ah.Logger.Error("Failed to Create Policy: ", err)
		errorRes := response.ClientResponse(http.StatusBadRequest, msg.MsgGettingDataErr, nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	ah.Logger.Info("Create Policy Successful")

	successRes := response.ClientResponse(http.StatusOK, msg.MsgSuccess, data, nil)
	c.JSON(http.StatusOK, successRes)
}

// UpdatePolicy updates an existing policy.
// @Summary Update an existing policy
// @Description Update an existing policy based on the provided data
// @Tags Admin Policy Management
// @Accept json
// @Produce json
// @Security BearerTokenAuth
// @Param policy body models.UpdatePolicyReq true "Updated policy data"
// @Success 200 {object} response.Response "Policy updated successfully"
// @Failure 400 {object} response.Response "Invalid policy data or failed to update policy"
// @Router /admin/policy [put]
func (ah *AdminHandler) UpdatePolicy(c *gin.Context) {

	var policyData models.UpdatePolicyReq

	if err := c.ShouldBindJSON(&policyData); err != nil {
		ah.Logger.Error("Failed to Get Data: ", err)
		errResp := response.ClientResponse(http.StatusBadRequest, msg.ErrFormat, nil, err.Error())
		c.JSON(http.StatusBadRequest, errResp)
		return
	}

	data, err := ah.GRPC_Client.UpdatePolicy(policyData)

	if err != nil {
		ah.Logger.Error("Failed to Update Policy: ", err)
		errorRes := response.ClientResponse(http.StatusBadRequest, msg.MsgGettingDataErr, nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	ah.Logger.Info("Update Policy Successful")

	successRes := response.ClientResponse(http.StatusOK, msg.MsgSuccess, data, nil)
	c.JSON(http.StatusOK, successRes)
}

// DeletePolicy deletes a policy by ID.
// @Summary Delete a policy
// @Description Delete a policy based on the provided ID
// @Tags Admin Policy Management
// @Accept json
// @Produce json
// @Security BearerTokenAuth
// @Param policy_id query int true "Policy ID to delete"
// @Success 200 {object} response.Response "Policy deleted successfully"
// @Failure 400 {object} response.Response "Invalid policy ID or failed to delete policy"
// @Router /admin/policy [delete]
func (ah *AdminHandler) DeletePolicy(c *gin.Context) {

	idStr := c.Query("policy_id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		ah.Logger.Error("Failed to Get Data: ", err)
		errorRes := response.ClientResponse(http.StatusBadRequest, msg.MsgIdGetErr, nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	data, err := ah.GRPC_Client.DeletePolicy(id)

	if err != nil {
		ah.Logger.Error("Failed to Delete Policy: ", err)
		errorRes := response.ClientResponse(http.StatusBadRequest, "Deleting operation failed", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	ah.Logger.Info("Delete Policy Successful")

	successRes := response.ClientResponse(http.StatusOK, msg.MsgSuccess, data, nil)
	c.JSON(http.StatusOK, successRes)
}

// GetAllPolicies retrieves all policies.
// @Summary Retrieve all policies
// @Description Retrieves all existing policies
// @Tags Admin Policy Management
// @Accept json
// @Produce json
// @Security BearerTokenAuth
// @Success 200 {object} response.Response "Policies retrieved successfully"
// @Failure 400 {object} response.Response "Failed to retrieve policies"
// @Router /admin/policies [get]
func (ah *AdminHandler) GetAllPolicies(c *gin.Context) {

	data, err := ah.GRPC_Client.GetAllPolicies()

	if err != nil {
		ah.Logger.Error("Failed to Get Data: ", err)
		errorRes := response.ClientResponse(http.StatusBadRequest, msg.MsgGettingDataErr, nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	ah.Logger.Info("Get All Policies Successful")

	successRes := response.ClientResponse(http.StatusOK, msg.MsgGetSucces, data, nil)
	c.JSON(http.StatusOK, successRes)
}

// GetOnePolicy retrieves a policy by ID.
// @Summary Retrieve a policy
// @Description Retrieves a policy based on the provided ID
// @Tags Admin Policy Management
// @Accept json
// @Produce json
// @Security BearerTokenAuth
// @Param policy_id query int true "Policy ID to retrieve"
// @Success 200 {object} response.Response "Policy retrieved successfully"
// @Failure 400 {object} response.Response "Invalid policy ID or failed to retrieve policy"
// @Router /admin/policy [get]
func (ah *AdminHandler) GetOnePolicy(c *gin.Context) {

	idStr := c.Query("policy_id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		ah.Logger.Error("Failed to Get Data: ", err)
		errorRes := response.ClientResponse(http.StatusBadRequest, msg.MsgIdGetErr, nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	data, err := ah.GRPC_Client.GetOnePolicy(id)
	if err != nil {
		ah.Logger.Error("Failed to Get One Policy: ", err)
		errorRes := response.ClientResponse(http.StatusBadRequest, msg.MsgGettingDataErr, nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	ah.Logger.Info("Get One Policy Successful")
	successRes := response.ClientResponse(http.StatusOK, msg.MsgGetSucces, data, nil)
	c.JSON(http.StatusOK, successRes)
}
