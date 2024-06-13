package handler

import (
	"errors"
	"net/http"
	"os"
	"strconv"

	logging "github.com/ARunni/connectHub_gateway/Logging"
	interfaces "github.com/ARunni/connectHub_gateway/pkg/client/job/interface"
	"github.com/ARunni/connectHub_gateway/pkg/utils/models"
	"github.com/ARunni/connectHub_gateway/pkg/utils/response"

	msg "github.com/ARunni/Error_Message"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type RecruiterJobHandler struct {
	GRPC_Client interfaces.RecruiterJobClient
	Logger      *logrus.Logger
	LogFile     *os.File
}

func NewRecruiterJobHandler(grpc_client interfaces.RecruiterJobClient) *RecruiterJobHandler {
	logger, logFile := logging.InitLogrusLogger("./Logging/connectHub_gateway.log")
	return &RecruiterJobHandler{
		GRPC_Client: grpc_client,
		Logger:      logger,
		LogFile:     logFile,
	}
}

// PostJob creates a new job posting by a recruiter.
// @Summary Create job posting
// @Description Create a new job posting by a recruiter
// @Tags Recruiter
// @Accept json
// @Produce json
// @Security BearerTokenAuth
// @Param body body models.JobOpening true "Job opening data for posting"
// @Success 200 {object} response.Response "Job posted successfully"
// @Failure 400 {object} response.Response "Failed to post job: missing or incorrect parameters"
// @Failure 500 {object} response.Response "Internal server error: failed to post job"
// @Router /recruiter/jobs [post]
func (jh *RecruiterJobHandler) PostJob(c *gin.Context) {

	recruiterID, ok := c.Get("id")
	if !ok {
		jh.Logger.Error("Failed to Get Data: ", errors.New("error in getting data"))
		errResp := response.ClientResponse(http.StatusBadRequest, msg.MsdGetIdErr, nil, errors.New(msg.ErrGetData))
		c.JSON(http.StatusBadRequest, errResp)
		return
	}
	recruiterId := recruiterID.(int)
	var data models.JobOpening
	data.EmployerID = recruiterId

	if !ok {
		jh.Logger.Error("Failed to Get Data: ", errors.New("convertion of datatype is failed"))
		errResp := response.ClientResponse(http.StatusBadRequest, msg.MsdGetIdErr, nil, errors.New(msg.ErrDatatypeConversion))
		c.JSON(http.StatusBadRequest, errResp)
		return
	}
	if err := c.ShouldBindJSON(&data); err != nil {
		jh.Logger.Error("Failed to Get Data: ", err)
		errResp := response.ClientResponse(http.StatusBadRequest, "Incorrect Format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errResp)
		return
	}

	jobData, err := jh.GRPC_Client.PostJob(data)
	if err != nil {
		jh.Logger.Error("Failed to Post Job: ", err)
		errResp := response.ClientResponse(http.StatusInternalServerError, "Cannot Post Job", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errResp)
		return
	}
	jh.Logger.Info("Jobseeker Signup Successful")
	successResp := response.ClientResponse(http.StatusOK, "Created job", jobData, nil)
	c.JSON(http.StatusOK, successResp)

}

// GetAllJobs retrieves all jobs posted by a recruiter.
// @Summary Get all jobs
// @Description Retrieve all jobs posted by a recruiter
// @Tags Recruiter
// @Accept json
// @Produce json
// @Security BearerTokenAuth
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} response.Response "Successfully retrieved jobs"
// @Failure 400 {object} response.Response "Failed to retrieve jobs: missing or incorrect parameters"
// @Failure 500 {object} response.Response "Internal server error: failed to retrieve jobs"
// @Router /recruiter/jobs [get]
func (jh *RecruiterJobHandler) GetAllJobs(c *gin.Context) {

	userIdstring, _ := c.Get("id")
	userId, strErr := userIdstring.(int)
	int32Value := int32(userId)
	if !strErr {
		jh.Logger.Error("Failed to Get Data: ", errors.New("details is not in correct format"))
		errResp := response.ClientResponse(http.StatusBadRequest, msg.MsgFormatErr, nil, strErr)
		c.JSON(http.StatusBadRequest, errResp)
		return
	}

	jobs, err := jh.GRPC_Client.GetAllJobs(int32Value)
	if err != nil {
		jh.Logger.Error("Failed to Get All Jobs: ", err)
		errs := response.ClientResponse(http.StatusInternalServerError, "Failed to fetch jobs", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}
	jh.Logger.Info("Jobs retrieved Successful")
	response := response.ClientResponse(http.StatusOK, "Jobs retrieved successfully", jobs, nil)
	c.JSON(http.StatusOK, response)
}

// GetOneJob retrieves the details of a single job posted by a recruiter.
// @Summary Get one job
// @Description Retrieve the details of a single job posted by a recruiter
// @Tags Recruiter
// @Accept json
// @Produce json
// @Security BearerTokenAuth
// @Param job_id query int true "ID of the job to retrieve"
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} response.Response "Successfully retrieved job"
// @Failure 400 {object} response.Response "Failed to retrieve job: missing or incorrect parameters"
// @Failure 500 {object} response.Response "Internal server error: failed to retrieve job"
// @Router /recruiter/job [get]
func (jh *RecruiterJobHandler) GetOneJob(c *gin.Context) {

	idStr := c.Query("job_id")

	userIdstring, _ := c.Get("id")
	userId, strErr := userIdstring.(int)
	int32Value := int32(userId)
	if !strErr {
		jh.Logger.Error("Failed to Get Data: ", errors.New("details is not in correct format"))
		errResp := response.ClientResponse(http.StatusBadRequest, msg.MsgFormatErr, nil, strErr)
		c.JSON(http.StatusBadRequest, errResp)
		return
	}

	jobID, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		jh.Logger.Error("Failed to Get Data: ", err)
		errs := response.ClientResponse(http.StatusBadRequest, "Invalid job ID", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	jobs, err := jh.GRPC_Client.GetOneJob(int32Value, int32(jobID))
	if err != nil {
		jh.Logger.Error("Failed to Get One Job: ", err)
		errs := response.ClientResponse(http.StatusInternalServerError, "Failed to fetch jobs", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}

	jh.Logger.Info("Jobs retrieved Successful")
	response := response.ClientResponse(http.StatusOK, "Jobs retrieved successfully", jobs, nil)
	c.JSON(http.StatusOK, response)
}

// DeleteAJob deletes a job posted by a recruiter.
// @Summary Delete a job
// @Description Delete a job posted by a recruiter
// @Tags Recruiter
// @Accept json
// @Produce json
// @Security BearerTokenAuth
// @Param job_id query int true "ID of the job to delete"
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} response.Response "Job successfully deleted"
// @Failure 400 {object} response.Response "Failed to delete job: missing or incorrect parameters"
// @Failure 500 {object} response.Response "Internal server error: failed to delete job"
// @Router /recruiter/job [delete]
func (jh *RecruiterJobHandler) DeleteAJob(c *gin.Context) {

	idStr := c.Query("job_id")

	userIdstring, _ := c.Get("id")
	userId, strErr := userIdstring.(int)
	int32Value := int32(userId)

	if !strErr {
		jh.Logger.Error("Failed to Get Data: ", errors.New("details not in correct format"))
		errResp := response.ClientResponse(http.StatusBadRequest, msg.MsgFormatErr, nil, strErr)
		c.JSON(http.StatusBadRequest, errResp)
		return
	}

	jobID, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		jh.Logger.Error("Invalid job ID: ", err)
		errs := response.ClientResponse(http.StatusBadRequest, "Invalid job ID", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	err = jh.GRPC_Client.DeleteAJob(int32Value, int32(jobID))

	if err != nil {
		jh.Logger.Error("Failed to Delete A Job: ", err)
		errs := response.ClientResponse(http.StatusInternalServerError, "Failed to delete job", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}
	jh.Logger.Info("Job Deleted successfully")
	response := response.ClientResponse(http.StatusOK, "Job Deleted successfully", nil, nil)
	c.JSON(http.StatusOK, response)
}

// UpdateAJob updates a job posted by a recruiter.
// @Summary Update a job
// @Description Update a job posted by a recruiter
// @Tags Recruiter
// @Accept json
// @Produce json
// @Security BearerTokenAuth
// @Param job_id query int true "ID of the job to update"
// @Param Authorization header string true "Bearer token"
// @Param body body models.JobOpening true "Updated job details"
// @Success 200 {object} response.Response "Job successfully updated"
// @Failure 400 {object} response.Response "Failed to update job: missing or incorrect parameters"
// @Failure 500 {object} response.Response "Internal server error: failed to update job"
// @Router /recruiter/job [put]
func (jh *RecruiterJobHandler) UpdateAJob(c *gin.Context) {

	idStr := c.Query("job_id")
	jobID, err := strconv.Atoi(idStr)
	if err != nil {
		jh.Logger.Error("Failed to Get Data: ", err)
		errs := response.ClientResponse(http.StatusBadRequest, "Invalid job ID", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	userIdstring, _ := c.Get("id")
	userId, strErr := userIdstring.(int)
	int32Value := int32(userId)
	if !strErr {
		jh.Logger.Error("Failed to Get Data: ", errors.New("details not in correct format"))
		errResp := response.ClientResponse(http.StatusBadRequest, msg.MsgFormatErr, nil, strErr)
		c.JSON(http.StatusBadRequest, errResp)
		return
	}

	var jobOpening models.JobOpening
	if err := c.ShouldBindJSON(&jobOpening); err != nil {
		jh.Logger.Error("Failed to Get Data: ", err)
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	UpdateJobOpening, err := jh.GRPC_Client.UpdateAJob(int32Value, int32(jobID), jobOpening)

	if err != nil {
		jh.Logger.Error("Failed to Update A Job: ", err)
		errs := response.ClientResponse(http.StatusInternalServerError, "Failed to update job", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}
	jh.Logger.Info("Job updated successfully")
	response := response.ClientResponse(http.StatusOK, "Job updated successfully", UpdateJobOpening, nil)
	c.JSON(http.StatusOK, response)
}

// GetJobAppliedCandidates retrieves the candidates who applied for a job posted by a recruiter.
// @Summary Get applied candidates for a job
// @Description Retrieves the candidates who applied for a job posted by a recruiter
// @Tags Recruiter
// @Accept json
// @Produce json
// @Security BearerTokenAuth
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} response.Response "Applied candidates retrieved successfully"
// @Failure 400 {object} response.Response "Failed to get applied candidates: missing or incorrect parameters"
// @Failure 500 {object} response.Response "Internal server error: failed to get applied candidates"
// @Router /recruiter/job/applied-candidates [get]
func (jh *RecruiterJobHandler) GetJobAppliedCandidates(c *gin.Context) {

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

	job, err := jh.GRPC_Client.GetJobAppliedCandidates(userId)
	if err != nil {
		jh.Logger.Error("Failed to Get Job Applied Candidates: ", err)
		errs := response.ClientResponse(http.StatusInternalServerError, "Failed to Getting Applied Jobs", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}
	jh.Logger.Info("Get Job Applied Candidates Successful")
	response := response.ClientResponse(http.StatusOK, "Getting Applied Jobs  successfully", job, nil)
	c.JSON(http.StatusOK, response)
}

// ScheduleInterview schedules an interview by a recruiter for a job applicant.
// @Summary Schedule interview
// @Description Schedule an interview by a recruiter for a job applicant
// @Tags Recruiter
// @Accept json
// @Produce json
// @Security BearerTokenAuth
// @Param Authorization header string true "Bearer token"
// @Param jobSchedule body models.ScheduleReq true "Schedule details"
// @Success 200 {object} response.Response "Interview scheduled successfully"
// @Failure 400 {object} response.Response "Failed to schedule interview: missing or incorrect parameters"
// @Failure 500 {object} response.Response "Internal server error: failed to schedule interview"
// @Router /recruiter/job/schedule-interview [post]
func (jh *RecruiterJobHandler) ScheduleInterview(c *gin.Context) {

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
	var jobSchedule models.ScheduleReq
	if err := c.ShouldBindJSON(&jobSchedule); err != nil {
		jh.Logger.Error("Failed to Get Data: ", err)
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	jobSchedule.RecruiterID = uint(userId)

	job, err := jh.GRPC_Client.ScheduleInterview(jobSchedule)
	if err != nil {
		jh.Logger.Error("Failed to Schedule Interview : ", err)
		errs := response.ClientResponse(http.StatusInternalServerError, "Failed to interview schedule", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}
	jh.Logger.Info("Schedule Interview Successful")
	response := response.ClientResponse(http.StatusOK, "interview scheduled successfully", job, nil)
	c.JSON(http.StatusOK, response)
}
