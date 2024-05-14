package handler

import (
	interfaces "connectHub_gateway/pkg/client/job/interface"
	"connectHub_gateway/pkg/utils/models"
	"connectHub_gateway/pkg/utils/response"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	msg "github.com/ARunni/Error_Message"
	"github.com/gin-gonic/gin"
)

type RecruiterJobHandler struct {
	GRPC_Client interfaces.RecruiterJobClient
}

func NewRecruiterJobHandler(grpc_client interfaces.RecruiterJobClient) *RecruiterJobHandler {
	return &RecruiterJobHandler{
		GRPC_Client: grpc_client,
	}
}

func (jh *RecruiterJobHandler) PostJob(c *gin.Context) {
	recruiterID, ok := c.Get("id")
	if !ok {
		errResp := response.ClientResponse(http.StatusBadRequest, msg.MsdGetIdErr, nil, errors.New(msg.ErrGetData))
		c.JSON(http.StatusBadRequest, errResp)
		return
	}
	recruiterId := recruiterID.(int)
	var data models.JobOpening
	data.EmployerID = recruiterId

	if !ok {
		errResp := response.ClientResponse(http.StatusBadRequest, msg.MsdGetIdErr, nil, errors.New(msg.ErrDatatypeConversion))
		c.JSON(http.StatusBadRequest, errResp)
		return
	}
	if err := c.ShouldBindJSON(&data); err != nil {
		errResp := response.ClientResponse(http.StatusBadRequest, "Incorrect Format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errResp)
		return
	}

	jobData, err := jh.GRPC_Client.PostJob(data)
	if err != nil {
		errResp := response.ClientResponse(http.StatusInternalServerError, "Cannot Post Job", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errResp)
		return
	}
	successResp := response.ClientResponse(http.StatusOK, "Created job", jobData, nil)
	c.JSON(http.StatusOK, successResp)

}

///

func (jh *RecruiterJobHandler) GetAllJobs(c *gin.Context) {

	userIdstring, _ := c.Get("id")
	userId, strErr := userIdstring.(int)
	int32Value := int32(userId)
	if !strErr {
		errResp := response.ClientResponse(http.StatusBadRequest, msg.MsgFormatErr, nil, strErr)
		c.JSON(http.StatusBadRequest, errResp)
		return
	}

	jobs, err := jh.GRPC_Client.GetAllJobs(int32Value)
	if err != nil {
		// Handle error if any
		errs := response.ClientResponse(http.StatusInternalServerError, "Failed to fetch jobs", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}

	response := response.ClientResponse(http.StatusOK, "Jobs retrieved successfully", jobs, nil)
	c.JSON(http.StatusOK, response)
}

func (jh *RecruiterJobHandler) GetOneJob(c *gin.Context) {
	idStr := c.Query("job_id")

	userIdstring, _ := c.Get("id")
	userId, strErr := userIdstring.(int)
	int32Value := int32(userId)
	if !strErr {
		errResp := response.ClientResponse(http.StatusBadRequest, msg.MsgFormatErr, nil, strErr)
		c.JSON(http.StatusBadRequest, errResp)
		return
	}

	jobID, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Invalid job ID", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	jobs, err := jh.GRPC_Client.GetOneJob(int32Value, int32(jobID))
	if err != nil {
		fmt.Println("ahjghgkg", err)
		errs := response.ClientResponse(http.StatusInternalServerError, "Failed to fetch jobs", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}

	response := response.ClientResponse(http.StatusOK, "Jobs retrieved successfully", jobs, nil)
	c.JSON(http.StatusOK, response)
}

func (jh *RecruiterJobHandler) DeleteAJob(c *gin.Context) {
	idStr := c.Query("job_id")

	userIdstring, _ := c.Get("id")
	userId, strErr := userIdstring.(int)
	int32Value := int32(userId)

	fmt.Println("userid string", userIdstring)
	fmt.Println("userid int", userId)
	fmt.Println("userid int32", int32Value)
	if !strErr {
		errResp := response.ClientResponse(http.StatusBadRequest, msg.MsgFormatErr, nil, strErr)
		c.JSON(http.StatusBadRequest, errResp)
		return
	}

	jobID, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Invalid job ID", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	err = jh.GRPC_Client.DeleteAJob(int32Value, int32(jobID))
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "Failed to delete job", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}

	response := response.ClientResponse(http.StatusOK, "Job Deleted successfully", nil, nil)
	c.JSON(http.StatusOK, response)
}

func (jh *RecruiterJobHandler) UpdateAJob(c *gin.Context) {

	idStr := c.Query("job_id")
	jobID, err := strconv.Atoi(idStr)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Invalid job ID", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	userIdstring, _ := c.Get("id")
	userId, strErr := userIdstring.(int)
	int32Value := int32(userId)
	if !strErr {
		errResp := response.ClientResponse(http.StatusBadRequest, msg.MsgFormatErr, nil, strErr)
		c.JSON(http.StatusBadRequest, errResp)
		return
	}

	var jobOpening models.JobOpening
	if err := c.ShouldBindJSON(&jobOpening); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	UpdateJobOpening, err := jh.GRPC_Client.UpdateAJob(int32Value, int32(jobID), jobOpening)
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "Failed to update job", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}

	response := response.ClientResponse(http.StatusOK, "Job updated successfully", UpdateJobOpening, nil)
	c.JSON(http.StatusOK, response)
}
