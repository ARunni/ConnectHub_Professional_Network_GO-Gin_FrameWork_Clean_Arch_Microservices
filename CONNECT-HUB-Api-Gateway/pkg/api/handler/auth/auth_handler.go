package handler

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"

	logging "github.com/ARunni/connectHub_gateway/Logging"
	interfaces "github.com/ARunni/connectHub_gateway/pkg/client/auth/interface"
	"github.com/ARunni/connectHub_gateway/pkg/utils/response"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type AuthHandler struct {
	GRPC_Client interfaces.AuthClient
	Logger      *logrus.Logger
	LogFile     *os.File
}

func NewAuthHandler(grpc_client interfaces.AuthClient) *AuthHandler {

	logger, logFile := logging.InitLogrusLogger("./Logging/connectHub_gateway.log")
	return &AuthHandler{
		GRPC_Client: grpc_client,
		Logger:      logger,
		LogFile:     logFile,
	}
}

func (au *AuthHandler) VideoCallKey(c *gin.Context) {

	au.Logger.Info("Processing VideoCallKey")
	userID, _ := c.Get("id")
	userType, _ := c.Get("role")
	role, ok := userType.(string)
	fmt.Println("roleeeeeeeeeeee", role)
	if !ok {
		au.Logger.Error("Role is not a string", errors.New("invalid role type"))
		errs := response.ClientResponse(http.StatusBadRequest, "Invalid role type", nil, errors.New("invalid role type"))
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	if role != "recruiter" {
		au.Logger.Error("Caller is not Recruiter", errors.New("role mismatch"))
		errs := response.ClientResponse(http.StatusBadRequest, "Caller is not Recruiter", nil, errors.New("role mismatch"))
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	UserID := c.Query("user")
	oppositeUser, err := strconv.Atoi(UserID)
	if err != nil {
		au.Logger.Error("Data Getting error", err)
		errs := response.ClientResponse(http.StatusBadRequest, "PostID not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	key, err := au.GRPC_Client.VideoCallKey(userID.(int), oppositeUser)
	if err != nil {

		au.Logger.Error("Error During VideoCallKey RPC call", err)
		errs := response.ClientResponse(http.StatusInternalServerError, "Couldn't not reterive link", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}
	au.Logger.Info("Successfully Get a VideoCallKey And Private Link")
	url := fmt.Sprintf("http://localhost:7000/index?room=%s", key)
	sucess := response.ClientResponse(http.StatusOK, "Successfully Get a VideoCallKey And Private Link", url, nil)
	c.JSON(http.StatusOK, sucess)
}
