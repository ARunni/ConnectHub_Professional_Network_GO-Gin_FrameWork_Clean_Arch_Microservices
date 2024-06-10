package handler

import (
	logging "connectHub_gateway/Logging"
	interfaces "connectHub_gateway/pkg/client/job/interface"
	"connectHub_gateway/pkg/utils/models"
	"connectHub_gateway/pkg/utils/response"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type JobseekerJobHandler struct {
	GRPC_Client interfaces.JobseekerJobClient
	Logger      *logrus.Logger
	LogFile     *os.File
}

func NewJobseekerJobHandler(grpc_client interfaces.JobseekerJobClient) *JobseekerJobHandler {
	logger, logFile := logging.InitLogrusLogger("./Logging/connectHub_gateway.log")
	return &JobseekerJobHandler{
		GRPC_Client: grpc_client,
		Logger:      logger,
		LogFile:     logFile,
	}
}

// JobSeekerGetAllJobs retrieves all jobs matching the provided keyword.
// @Summary Get all jobs for job seeker
// @Description Retrieve all jobs matching the provided keyword for job seekers
// @Tags Job Seeker
// @Accept json
// @Produce json
// @Security BearerTokenAuth
// @Param Keyword query string true "Keyword to search for jobs"
// @Success 200 {object} response.Response "Successfully retrieved jobs"
// @Failure 400 {object} response.Response "Failed to retrieve jobs: keyword parameter is required"
// @Failure 500 {object} response.Response "Internal server error: failed to fetch jobs"
// @Router /jobseeker/jobs [get]
func (jh *JobseekerJobHandler) JobSeekerGetAllJobs(c *gin.Context) {

	keyword := c.Query("Keyword")

	if keyword == "" {

		jh.Logger.Error("Failed to Get Data: ", errors.New("keyword parameter is required"))

		errs := response.ClientResponse(http.StatusBadRequest, "Keyword parameter is required", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	jobs, err := jh.GRPC_Client.JobSeekerGetAllJobs(keyword)

	if err != nil {

		jh.Logger.Error("Failed to Job Seeker Get All Jobs: ", err)
		errs := response.ClientResponse(http.StatusInternalServerError, "Failed to fetch jobs", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}

	if len(jobs) == 0 {
		jh.Logger.Error("Failed to Job Seeker Get All Jobs: ", errors.New("no jobs found matching your query"))
		errMsg := "No jobs found matching your query"
		errs := response.ClientResponse(http.StatusOK, errMsg, nil, nil)
		c.JSON(http.StatusOK, errs)
		return
	}

	jh.Logger.Info("Jobs retrieved successfully")

	response := response.ClientResponse(http.StatusOK, "Jobs retrieved successfully", jobs, nil)
	c.JSON(http.StatusOK, response)
}

// JobSeekerGetJobByID retrieves the details of a job by its ID for job seekers.
// @Summary Get job by ID for job seeker
// @Description Retrieve the details of a job by its ID for job seekers
// @Tags Job Seeker
// @Accept json
// @Produce json
// @Security BearerTokenAuth
// @Param job_id query int true "ID of the job to retrieve"
// @Success 200 {object} response.Response "Successfully retrieved job"
// @Failure 400 {object} response.Response "Failed to retrieve job: job_id parameter is required or conversion failed"
// @Failure 500 {object} response.Response "Internal server error: failed to fetch job"
// @Router /jobseeker/job [get]
func (jh *JobseekerJobHandler) JobSeekerGetJobByID(c *gin.Context) {

	jobIdString := c.Query("job_id")

	if jobIdString == "" {

		jh.Logger.Error("Failed to Get Data: ", errors.New("job_id parameter is required"))
		errs := response.ClientResponse(http.StatusBadRequest, "job_id parameter is required", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	jobId, err := strconv.Atoi(jobIdString)
	if err != nil {
		jh.Logger.Error("job_id conversion failed: ", err)
		errs := response.ClientResponse(http.StatusBadRequest, "job_id conversion failed", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	job, err := jh.GRPC_Client.JobSeekerGetJobByID(jobId)

	if err != nil {
		jh.Logger.Error("Job Seeker Get Job By ID failed: ", err)
		errs := response.ClientResponse(http.StatusInternalServerError, "Failed to fetch job", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}

	jh.Logger.Info("Job retrieved successfully")

	response := response.ClientResponse(http.StatusOK, "Job retrieved successfully", job, nil)
	c.JSON(http.StatusOK, response)
}

// JobSeekerApplyJob handles the job application operation for a job seeker.
// @Summary Apply for a job
// @Description Apply for a job as a job seeker
// @Tags Job Seeker
// @Accept multipart/form-data
// @Produce json
// @Security BearerTokenAuth
// @Param job_id formData int true "ID of the job to apply for"
// @Param cover_letter formData string true "Cover letter for the job application"
// @Param resume formData file true "Resume for the job application"
// @Success 200 {object} response.Response "Job applied successfully"
// @Failure 400 {object} response.Response "Failed to apply for the job: missing or incorrect parameters"
// @Failure 500 {object} response.Response "Internal server error: failed to apply for the job"
// @Router /jobseeker/apply [post]
func (jh *JobseekerJobHandler) JobSeekerApplyJob(c *gin.Context) {

	userIdAny, ok := c.Get("id")
	if !ok {
		jh.Logger.Error("Failed to Get Data: ", errors.New("getting user id failed"))
		errs := response.ClientResponse(http.StatusBadRequest, "getting user id failed", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	userId, ok := userIdAny.(int)
	if !ok {
		jh.Logger.Error("Failed to Get Data: ", errors.New("converting user id failed"))
		errs := response.ClientResponse(http.StatusBadRequest, "converting user id failed", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	var data models.ApplyJobReq
	data.JobseekerID = uint(userId)
	data.CoverLetter = c.PostForm("cover_letter")
	JobIDstr := c.PostForm("job_id")
	JobID, err := strconv.Atoi(JobIDstr)
	if err != nil {
		jh.Logger.Error("Failed to Get Data: ", errors.New("getting coverletter Failed"))
		errResp := response.ClientResponse(http.StatusInternalServerError, "Getting coverletter Failed", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errResp)
		return
	}
	data.JobID = uint(JobID)

	file, err := c.FormFile("resume")
	if err != nil {
		jh.Logger.Error("Failed to Get Data: ", errors.New("getting resume Failed"))
		errResp := response.ClientResponse(http.StatusInternalServerError, "Getting resume Failed", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errResp)
		return
	}
	fileContent, err := file.Open()
	if err != nil {
		jh.Logger.Error("Failed to Get Data: ", errors.New("error opening the file"))
		errResp := response.ClientResponse(http.StatusInternalServerError, "Error opening the file", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errResp)
		return

	}
	defer fileContent.Close()
	resumeData, err := ioutil.ReadAll(fileContent)
	if err != nil {
		jh.Logger.Error("Failed to Get Data: ", errors.New("error reading the file"))
		errResp := response.ClientResponse(http.StatusInternalServerError, "Error reading the file", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errResp)
		return
	}
	data.Resume = resumeData

	if data.JobID <= 0 {
		jh.Logger.Error("Failed to Get Data: ", errors.New("job_id parameter is required"))
		errs := response.ClientResponse(http.StatusBadRequest, "job_id parameter is required", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	job, err := jh.GRPC_Client.JobSeekerApplyJob(data)

	if err != nil {
		jh.Logger.Error("Failed to apply job: ", err)
		errs := response.ClientResponse(http.StatusInternalServerError, "Failed to apply job", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}
	jh.Logger.Info("Job applied successfully")

	response := response.ClientResponse(http.StatusOK, "Job applied successfully", job, nil)
	c.JSON(http.StatusOK, response)
}

// GetAppliedJobs retrieves the jobs that a job seeker has applied for.
// @Summary Get applied jobs
// @Description Retrieve the jobs that a job seeker has applied for
// @Tags Job Seeker
// @Accept json
// @Produce json
// @Security BearerTokenAuth
// @Success 200 {object} response.Response "Successfully retrieved applied jobs"
// @Failure 400 {object} response.Response "Failed to retrieve applied jobs: missing or incorrect parameters"
// @Failure 500 {object} response.Response "Internal server error: failed to retrieve applied jobs"
// @Router /jobseeker/appliedjobs [get]
func (jh *JobseekerJobHandler) GetAppliedJobs(c *gin.Context) {

	userIdAny, ok := c.Get("id")
	if !ok {
		jh.Logger.Error("Failed to Get Data: ", errors.New("getting user id failed"))
		errs := response.ClientResponse(http.StatusBadRequest, "getting user id failed", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	userId, ok := userIdAny.(int)
	if !ok {
		jh.Logger.Error("Failed to Get Data: ", errors.New("converting user id failed"))
		errs := response.ClientResponse(http.StatusBadRequest, "converting user id failed", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	job, err := jh.GRPC_Client.GetAppliedJobs(userId)

	if err != nil {
		jh.Logger.Error("Failed to Get Applied Jobs: ", err)
		errs := response.ClientResponse(http.StatusInternalServerError, "Failed to Getting Applied Jobs", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}
	jh.Logger.Info("Getting Applied Jobs  successfully")
	response := response.ClientResponse(http.StatusOK, "Getting Applied Jobs  successfully", job, nil)
	c.JSON(http.StatusOK, response)
}
