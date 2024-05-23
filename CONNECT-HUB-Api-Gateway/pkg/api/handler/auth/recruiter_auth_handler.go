package handler

import (
	interfaces "connectHub_gateway/pkg/client/auth/interface"
	"connectHub_gateway/pkg/utils/models"
	"connectHub_gateway/pkg/utils/response"
	"errors"
	"fmt"
	"net/http"

	msg "github.com/ARunni/Error_Message"
	"github.com/gin-gonic/gin"
)

type RecruiterHandler struct {
	GRPC_Client interfaces.RecruiterAuthClient
}

func NewRecruiterAuthHandler(grpc_client interfaces.RecruiterAuthClient) *RecruiterHandler {
	return &RecruiterHandler{
		GRPC_Client: grpc_client,
	}
}

func (jh *RecruiterHandler) RecruiterSignup(c *gin.Context) {

	var recruiterData models.RecruiterSignUp

	if err := c.ShouldBindJSON(&recruiterData); err != nil {
		errResp := response.ClientResponse(http.StatusBadRequest, "Incorrect Format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errResp)
		return
	}

	recruiter, err := jh.GRPC_Client.RecruiterSignup(recruiterData)

	if err != nil {
		errResp := response.ClientResponse(http.StatusInternalServerError, "Signup failed Recruiter", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errResp)
		return
	}

	successResp := response.ClientResponse(http.StatusOK, "Recruiter Signup Successfully", recruiter, nil)
	c.JSON(http.StatusOK, successResp)

}
func (jh *RecruiterHandler) RecruiterLogin(c *gin.Context) {

	var recruiterData models.RecruiterLogin

	if err := c.ShouldBindJSON(&recruiterData); err != nil {
		errResp := response.ClientResponse(http.StatusBadRequest, "Incorrect Format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errResp)
		return
	}

	recruiter, err := jh.GRPC_Client.RecruiterLogin(recruiterData)

	if err != nil {
		errResp := response.ClientResponse(http.StatusInternalServerError, "Cannot authenticate Recruiter", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errResp)
		return
	}

	successResp := response.ClientResponse(http.StatusOK, "Recruiter Authenticated Successfully", recruiter, nil)
	c.JSON(http.StatusOK, successResp)

}

func (jh *RecruiterHandler) RecruiterGetProfile(c *gin.Context) {
	userIdstring, ok := c.Get("id")
	fmt.Println("status ", ok)
	if !ok {
		err := errors.New("error in getting id")
		errResp := response.ClientResponse(http.StatusBadRequest, msg.MsgIdGetErr, nil, err.Error())
		c.JSON(http.StatusBadRequest, errResp)
		return
	}
	userId, strErr := userIdstring.(int)
	fmt.Println("recruiter id ", userId)
	fmt.Println("recruiter id ", userIdstring)
	if !strErr {
		errResp := response.ClientResponse(http.StatusBadRequest, msg.MsgFormatErr, nil, strErr)
		c.JSON(http.StatusBadRequest, errResp)
		return
	}
	recruiter, err := jh.GRPC_Client.RecruiterGetProfile(userId)
	if err != nil {
		errREsp := response.ClientResponse(http.StatusBadRequest, msg.MsgGettingDataErr, nil, err.Error())
		c.JSON(http.StatusBadRequest, errREsp)
		return
	}
	successResp := response.ClientResponse(http.StatusOK, msg.MsgGetSucces, recruiter, nil)
	c.JSON(http.StatusOK, successResp)

}

func (jh *RecruiterHandler) RecruiterEditProfile(c *gin.Context) {

	userIdstring, _ := c.Get("id")
	userId, strErr := userIdstring.(int)
	if !strErr {
		errResp := response.ClientResponse(http.StatusBadRequest, msg.MsgFormatErr, nil, strErr)
		c.JSON(http.StatusBadRequest, errResp)
		return
	}

	var recruiterData models.RecruiterProfile
	recruiterData.ID = uint(userId)

	if err := c.ShouldBindJSON(&recruiterData); err != nil {
		errResp := response.ClientResponse(http.StatusBadRequest, "Incorrect Format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errResp)
		return
	}

	recruiter, err := jh.GRPC_Client.RecruiterEditProfile(recruiterData)
	if err != nil {
		errREsp := response.ClientResponse(http.StatusBadRequest, msg.MsgGettingDataErr, nil, err.Error())
		c.JSON(http.StatusBadRequest, errREsp)
		return
	}
	successResp := response.ClientResponse(http.StatusOK, msg.MsgGetSucces, recruiter, nil)
	c.JSON(http.StatusOK, successResp)

}
