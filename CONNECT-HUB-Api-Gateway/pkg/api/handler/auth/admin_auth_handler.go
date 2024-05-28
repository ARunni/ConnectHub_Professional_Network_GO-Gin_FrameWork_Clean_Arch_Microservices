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

type AdminHandler struct {
	GRPC_Client interfaces.AdminAuthClient
}

func NewAdminAuthHandler(grpc_client interfaces.AdminAuthClient) *AdminHandler {
	return &AdminHandler{
		GRPC_Client: grpc_client,
	}
}

func (ah *AdminHandler) AdminLogin(c *gin.Context) {
	var adminData models.AdminLogin

	if err := c.ShouldBindJSON(&adminData); err != nil {
		errResp := response.ClientResponse(http.StatusBadRequest, msg.ErrFormat, nil, err.Error())
		c.JSON(http.StatusBadRequest, errResp)
		return
	}
	admin, err := ah.GRPC_Client.AdminLogin(adminData)
	if err != nil {
		errResp := response.ClientResponse(http.StatusInternalServerError, "Cannot authenticate Admin", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errResp)
		return
	}
	successResp := response.ClientResponse(http.StatusOK, "Admin Authenticated Successfully", admin, nil)
	c.JSON(http.StatusOK, successResp)

}

func (ah *AdminHandler) GetJobseekers(c *gin.Context) {
	pageStr := c.Query("page")
	page, err := strconv.Atoi(pageStr)

	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, msg.MsgPageNumFormatErr, nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	jobseeker, err := ah.GRPC_Client.GetJobseekers(page)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, msg.MsgGettingDataErr, nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, msg.MsgGetSucces, jobseeker, nil)
	c.JSON(http.StatusOK, successRes)
}

func (ah *AdminHandler) GetRecruiters(c *gin.Context) {
	pageStr := c.Query("page")
	page, err := strconv.Atoi(pageStr)

	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, msg.MsgPageNumFormatErr, nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	recruiter, err := ah.GRPC_Client.GetRecruiters(page)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, msg.MsgGettingDataErr, nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, msg.MsgGetSucces, recruiter, nil)
	c.JSON(http.StatusOK, successRes)
}

func (ah *AdminHandler) BlockRecruiter(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, msg.MsgPageNumFormatErr, nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	blockRecruiter, err := ah.GRPC_Client.BlockRecruiter(id)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, msg.MsgGettingDataErr, nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, msg.ErrUserBlockTrue, blockRecruiter, nil)
	c.JSON(http.StatusOK, successRes)
}

func (ah *AdminHandler) BlockJobseeker(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, msg.MsgPageNumFormatErr, nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	blockJobseeker, err := ah.GRPC_Client.BlockJobseeker(id)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, msg.MsgGettingDataErr, nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, msg.ErrUserBlockTrue, blockJobseeker, nil)
	c.JSON(http.StatusOK, successRes)
}

func (ah *AdminHandler) UnBlockJobseeker(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, msg.MsgPageNumFormatErr, nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	unBlockjobseeker, err := ah.GRPC_Client.UnBlockJobseeker(id)

	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, msg.MsgGettingDataErr, nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, msg.MsgGetSucces, unBlockjobseeker, nil)
	c.JSON(http.StatusOK, successRes)
}

func (ah *AdminHandler) UnBlockRecruiter(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, msg.MsgPageNumFormatErr, nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	unBlockRecruiter, err := ah.GRPC_Client.UnBlockRecruiter(id)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, msg.MsgGettingDataErr, nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, msg.MsgGetSucces, unBlockRecruiter, nil)
	c.JSON(http.StatusOK, successRes)
}

func (ah *AdminHandler) GetJobseekerDetails(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, msg.MsgPageNumFormatErr, nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	unBlockRecruiter, err := ah.GRPC_Client.GetJobseekerDetails(id)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, msg.MsgGettingDataErr, nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, msg.MsgGetSucces, unBlockRecruiter, nil)
	c.JSON(http.StatusOK, successRes)
}

func (ah *AdminHandler) GetRecruiterDetails(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, msg.MsgPageNumFormatErr, nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	unBlockRecruiter, err := ah.GRPC_Client.GetRecruiterDetails(id)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, msg.MsgGettingDataErr, nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, msg.MsgGetSucces, unBlockRecruiter, nil)
	c.JSON(http.StatusOK, successRes)
}

// policies

func (ah *AdminHandler) CreatePolicy(c *gin.Context) {
	var policyData models.CreatePolicyReq

	if err := c.ShouldBindJSON(&policyData); err != nil {
		errResp := response.ClientResponse(http.StatusBadRequest, msg.ErrFormat, nil, err.Error())
		c.JSON(http.StatusBadRequest, errResp)
		return
	}

	data, err := ah.GRPC_Client.CreatePolicy(policyData)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, msg.MsgGettingDataErr, nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, msg.MsgSuccess, data, nil)
	c.JSON(http.StatusOK, successRes)
}

func (ah *AdminHandler) UpdatePolicy(c *gin.Context) {

	var policyData models.UpdatePolicyReq

	if err := c.ShouldBindJSON(&policyData); err != nil {
		errResp := response.ClientResponse(http.StatusBadRequest, msg.ErrFormat, nil, err.Error())
		c.JSON(http.StatusBadRequest, errResp)
		return
	}

	data, err := ah.GRPC_Client.UpdatePolicy(policyData)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, msg.MsgGettingDataErr, nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, msg.MsgSuccess, data, nil)
	c.JSON(http.StatusOK, successRes)
}

func (ah *AdminHandler) DeletePolicy(c *gin.Context) {
	idStr := c.Query("policy_id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, msg.MsgIdGetErr, nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	data, err := ah.GRPC_Client.DeletePolicy(id)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "Deleting operation failed", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, msg.MsgSuccess, data, nil)
	c.JSON(http.StatusOK, successRes)
}

func (ah *AdminHandler) GetAllPolicies(c *gin.Context) {

	data, err := ah.GRPC_Client.GetAllPolicies()
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, msg.MsgGettingDataErr, nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, msg.MsgGetSucces, data, nil)
	c.JSON(http.StatusOK, successRes)
}

func (ah *AdminHandler) GetOnePolicy(c *gin.Context) {
	idStr := c.Query("policy_id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, msg.MsgIdGetErr, nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	data, err := ah.GRPC_Client.GetOnePolicy(id)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, msg.MsgGettingDataErr, nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, msg.MsgGetSucces, data, nil)
	c.JSON(http.StatusOK, successRes)
}
