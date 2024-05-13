package handler

import (
	interfaces "connectHub_gateway/pkg/client/interface"
	"connectHub_gateway/pkg/utils/models"
	"connectHub_gateway/pkg/utils/response"
	"errors"
	"net/http"

	msg "github.com/ARunni/Error_Message"
	"github.com/gin-gonic/gin"
)

type JobHandler struct {
	GRPC_Client interfaces.JobClient
}

func NewJobHandler(grpc_client interfaces.JobClient) *JobHandler {
	return &JobHandler{
		GRPC_Client: grpc_client,
	}
}

func (jh *JobHandler) PostJob(c *gin.Context) {
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
