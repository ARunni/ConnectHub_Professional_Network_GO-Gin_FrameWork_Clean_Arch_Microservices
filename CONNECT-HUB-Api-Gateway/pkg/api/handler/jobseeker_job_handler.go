package handler

import (
	interfaces "connectHub_gateway/pkg/client/job/interface"
	"connectHub_gateway/pkg/utils/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type JobseekerJobHandler struct {
	GRPC_Client interfaces.JobseekerJobClient
}

func NewJobseekerJobHandler(grpc_client interfaces.JobseekerJobClient) *JobseekerJobHandler {
	return &JobseekerJobHandler{
		GRPC_Client: grpc_client,
	}
}

func (jh *JobseekerJobHandler) JobSeekerGetAllJobs(c *gin.Context) {
	keyword := c.Query("Keyword")

	if keyword == "" {
		errs := response.ClientResponse(http.StatusBadRequest, "Keyword parameter is required", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	jobs, err := jh.GRPC_Client.JobSeekerGetAllJobs(keyword)
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "Failed to fetch jobs", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}

	if len(jobs) == 0 {
		errMsg := "No jobs found matching your query"
		errs := response.ClientResponse(http.StatusOK, errMsg, nil, nil)
		c.JSON(http.StatusOK, errs)
		return
	}

	response := response.ClientResponse(http.StatusOK, "Jobs retrieved successfully", jobs, nil)
	c.JSON(http.StatusOK, response)
}
