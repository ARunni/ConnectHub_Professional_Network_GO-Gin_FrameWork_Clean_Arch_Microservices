package handler

import (
	logging "connectHub_gateway/Logging"
	interfaces "connectHub_gateway/pkg/client/job/interface"
	"connectHub_gateway/pkg/utils/models"
	"connectHub_gateway/pkg/utils/response"
	"errors"
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

	logrusLogger, logrusLogFile := logging.InitLogrusLogger("./Logging/connectHub_gateway.log")
	defer logrusLogFile.Close()

	recruiterID, ok := c.Get("id")
	if !ok {
		logrusLogger.Error("Failed to Get Data: ", errors.New("error in getting data"))
		errResp := response.ClientResponse(http.StatusBadRequest, msg.MsdGetIdErr, nil, errors.New(msg.ErrGetData))
		c.JSON(http.StatusBadRequest, errResp)
		return
	}
	recruiterId := recruiterID.(int)
	var data models.JobOpening
	data.EmployerID = recruiterId

	if !ok {
		logrusLogger.Error("Failed to Get Data: ", errors.New("convertion of datatype is failed"))
		errResp := response.ClientResponse(http.StatusBadRequest, msg.MsdGetIdErr, nil, errors.New(msg.ErrDatatypeConversion))
		c.JSON(http.StatusBadRequest, errResp)
		return
	}
	if err := c.ShouldBindJSON(&data); err != nil {
		logrusLogger.Error("Failed to Get Data: ", err)
		errResp := response.ClientResponse(http.StatusBadRequest, "Incorrect Format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errResp)
		return
	}

	jobData, err := jh.GRPC_Client.PostJob(data)
	if err != nil {
		logrusLogger.Error("Failed to Post Job: ", err)
		errResp := response.ClientResponse(http.StatusInternalServerError, "Cannot Post Job", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errResp)
		return
	}
	logrusLogger.Info("Jobseeker Signup Successful")
	successResp := response.ClientResponse(http.StatusOK, "Created job", jobData, nil)
	c.JSON(http.StatusOK, successResp)

}

///

func (jh *RecruiterJobHandler) GetAllJobs(c *gin.Context) {

	logrusLogger, logrusLogFile := logging.InitLogrusLogger("./Logging/connectHub_gateway.log")
	defer logrusLogFile.Close()

	userIdstring, _ := c.Get("id")
	userId, strErr := userIdstring.(int)
	int32Value := int32(userId)
	if !strErr {
		logrusLogger.Error("Failed to Get Data: ", errors.New("Details is not in correct format"))
		errResp := response.ClientResponse(http.StatusBadRequest, msg.MsgFormatErr, nil, strErr)
		c.JSON(http.StatusBadRequest, errResp)
		return
	}

	jobs, err := jh.GRPC_Client.GetAllJobs(int32Value)
	if err != nil {
		logrusLogger.Error("Failed to Get All Jobs: ", err)
		errs := response.ClientResponse(http.StatusInternalServerError, "Failed to fetch jobs", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}
	logrusLogger.Info("Jobs retrieved Successful")
	response := response.ClientResponse(http.StatusOK, "Jobs retrieved successfully", jobs, nil)
	c.JSON(http.StatusOK, response)
}

func (jh *RecruiterJobHandler) GetOneJob(c *gin.Context) {

	logrusLogger, logrusLogFile := logging.InitLogrusLogger("./Logging/connectHub_gateway.log")
	defer logrusLogFile.Close()

	idStr := c.Query("job_id")

	userIdstring, _ := c.Get("id")
	userId, strErr := userIdstring.(int)
	int32Value := int32(userId)
	if !strErr {
		logrusLogger.Error("Failed to Get Data: ", errors.New("Details is not in correct format"))
		errResp := response.ClientResponse(http.StatusBadRequest, msg.MsgFormatErr, nil, strErr)
		c.JSON(http.StatusBadRequest, errResp)
		return
	}

	jobID, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		logrusLogger.Error("Failed to Get Data: ", err)
		errs := response.ClientResponse(http.StatusBadRequest, "Invalid job ID", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	jobs, err := jh.GRPC_Client.GetOneJob(int32Value, int32(jobID))
	if err != nil {
		logrusLogger.Error("Failed to Get One Job: ", err)
		errs := response.ClientResponse(http.StatusInternalServerError, "Failed to fetch jobs", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}

	logrusLogger.Info("Jobs retrieved Successful")
	response := response.ClientResponse(http.StatusOK, "Jobs retrieved successfully", jobs, nil)
	c.JSON(http.StatusOK, response)
}

func (jh *RecruiterJobHandler) DeleteAJob(c *gin.Context) {

	logrusLogger, logrusLogFile := logging.InitLogrusLogger("./Logging/connectHub_gateway.log")
	defer logrusLogFile.Close()

	idStr := c.Query("job_id")

	userIdstring, _ := c.Get("id")
	userId, strErr := userIdstring.(int)
	int32Value := int32(userId)

	if !strErr {
		logrusLogger.Error("Failed to Get Data: ", errors.New("Details not in correct format"))
		errResp := response.ClientResponse(http.StatusBadRequest, msg.MsgFormatErr, nil, strErr)
		c.JSON(http.StatusBadRequest, errResp)
		return
	}

	jobID, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		logrusLogger.Error("Invalid job ID: ", err)
		errs := response.ClientResponse(http.StatusBadRequest, "Invalid job ID", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	err = jh.GRPC_Client.DeleteAJob(int32Value, int32(jobID))

	if err != nil {
		logrusLogger.Error("Failed to Delete A Job: ", err)
		errs := response.ClientResponse(http.StatusInternalServerError, "Failed to delete job", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}
	logrusLogger.Info("Job Deleted successfully")
	response := response.ClientResponse(http.StatusOK, "Job Deleted successfully", nil, nil)
	c.JSON(http.StatusOK, response)
}

func (jh *RecruiterJobHandler) UpdateAJob(c *gin.Context) {

	logrusLogger, logrusLogFile := logging.InitLogrusLogger("./Logging/connectHub_gateway.log")
	defer logrusLogFile.Close()

	idStr := c.Query("job_id")
	jobID, err := strconv.Atoi(idStr)
	if err != nil {
		logrusLogger.Error("Failed to Get Data: ", err)
		errs := response.ClientResponse(http.StatusBadRequest, "Invalid job ID", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	userIdstring, _ := c.Get("id")
	userId, strErr := userIdstring.(int)
	int32Value := int32(userId)
	if !strErr {
		logrusLogger.Error("Failed to Get Data: ", errors.New("Details not in correct format"))
		errResp := response.ClientResponse(http.StatusBadRequest, msg.MsgFormatErr, nil, strErr)
		c.JSON(http.StatusBadRequest, errResp)
		return
	}

	var jobOpening models.JobOpening
	if err := c.ShouldBindJSON(&jobOpening); err != nil {
		logrusLogger.Error("Failed to Get Data: ", err)
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	UpdateJobOpening, err := jh.GRPC_Client.UpdateAJob(int32Value, int32(jobID), jobOpening)

	if err != nil {
		logrusLogger.Error("Failed to Update A Job: ", err)
		errs := response.ClientResponse(http.StatusInternalServerError, "Failed to update job", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}
	logrusLogger.Info("Job updated successfully")
	response := response.ClientResponse(http.StatusOK, "Job updated successfully", UpdateJobOpening, nil)
	c.JSON(http.StatusOK, response)
}

func (jh *RecruiterJobHandler) GetJobAppliedCandidates(c *gin.Context) {

	logrusLogger, logrusLogFile := logging.InitLogrusLogger("./Logging/connectHub_gateway.log")
	defer logrusLogFile.Close()

	userIdAny, ok := c.Get("id")
	if !ok {
		logrusLogger.Error("Failed to Get Data: ", errors.New("getting user id failed"))
		errs := response.ClientResponse(http.StatusBadRequest, "getting user id failed", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	userId, ok := userIdAny.(int)
	if !ok {
		logrusLogger.Error("Failed to Get Data: ", errors.New("converting user id failed"))
		errs := response.ClientResponse(http.StatusBadRequest, "converting user id failed", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	job, err := jh.GRPC_Client.GetJobAppliedCandidates(userId)
	if err != nil {
		logrusLogger.Error("Failed to Get Job Applied Candidates: ", err)
		errs := response.ClientResponse(http.StatusInternalServerError, "Failed to Getting Applied Jobs", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}
	logrusLogger.Info("Get Job Applied Candidates Successful")
	response := response.ClientResponse(http.StatusOK, "Getting Applied Jobs  successfully", job, nil)
	c.JSON(http.StatusOK, response)
}

func (jh *RecruiterJobHandler) ScheduleInterview(c *gin.Context) {

	logrusLogger, logrusLogFile := logging.InitLogrusLogger("./Logging/connectHub_gateway.log")
	defer logrusLogFile.Close()

	userIdAny, ok := c.Get("id")
	if !ok {
		logrusLogger.Error("Failed to Get Data: ", errors.New("getting user id failed"))
		errs := response.ClientResponse(http.StatusBadRequest, "getting user id failed", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	userId, ok := userIdAny.(int)
	if !ok {
		logrusLogger.Error("Failed to Get Data: ", errors.New("converting user id failed"))
		errs := response.ClientResponse(http.StatusBadRequest, "converting user id failed", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	var jobSchedule models.ScheduleReq
	if err := c.ShouldBindJSON(&jobSchedule); err != nil {
		logrusLogger.Error("Failed to Get Data: ", err)
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	jobSchedule.RecruiterID = uint(userId)

	job, err := jh.GRPC_Client.ScheduleInterview(jobSchedule)
	if err != nil {
		logrusLogger.Error("Failed to Schedule Interview : ", err)
		errs := response.ClientResponse(http.StatusInternalServerError, "Failed to interview schedule", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}
	logrusLogger.Info("Schedule Interview Successful")
	response := response.ClientResponse(http.StatusOK, "interview scheduled successfully", job, nil)
	c.JSON(http.StatusOK, response)
}
