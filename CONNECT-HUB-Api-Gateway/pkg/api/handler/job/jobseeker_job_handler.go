package handler

import (
	interfaces "connectHub_gateway/pkg/client/job/interface"
	"connectHub_gateway/pkg/utils/response"
	"net/http"
	"strconv"

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

func (jh *JobseekerJobHandler) JobSeekerGetJobByID(c *gin.Context) {
	jobIdString := c.Query("job_id")

	if jobIdString == "" {
		errs := response.ClientResponse(http.StatusBadRequest, "job_id parameter is required", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	jobId, err := strconv.Atoi(jobIdString)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "job_id conversion failed", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	job, err := jh.GRPC_Client.JobSeekerGetJobByID(jobId)
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "Failed to fetch job", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}

	response := response.ClientResponse(http.StatusOK, "Job retrieved successfully", job, nil)
	c.JSON(http.StatusOK, response)
}

func (jh *JobseekerJobHandler) JobSeekerApplyJob(c *gin.Context) {
	jobIdString := c.Query("job_id")
	userIdAny, ok := c.Get("id")
	if !ok {
		errs := response.ClientResponse(http.StatusBadRequest, "getting user id failed", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	userId, ok := userIdAny.(int)
	if !ok {
		errs := response.ClientResponse(http.StatusBadRequest, "converting user id failed", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	if jobIdString == "" {
		errs := response.ClientResponse(http.StatusBadRequest, "job_id parameter is required", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	// fmt.Println("hhahhglakg", jobIdString)
	jobId, err := strconv.Atoi(jobIdString)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "job_id conversion failed", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	job, err := jh.GRPC_Client.JobSeekerApplyJob(jobId, userId)
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "Failed to apply job", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}

	response := response.ClientResponse(http.StatusOK, "Job applied successfully", job, nil)
	c.JSON(http.StatusOK, response)
}
