package handler

import (
	logging "github.com/ARunni/connectHub_gateway/Logging"
	interfaces "github.com/ARunni/connectHub_gateway/pkg/client/chat/interfaces"
	"github.com/ARunni/connectHub_gateway/pkg/helper"
	"github.com/ARunni/connectHub_gateway/pkg/utils/models"
	"github.com/ARunni/connectHub_gateway/pkg/utils/response"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

var upgrade = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var User = make(map[string]*websocket.Conn)

type ChatHandler struct {
	GRPC_Client interfaces.ChatClient
	helper      *helper.Helper
	Logger      *logrus.Logger
	LogFile     *os.File
}

func NewChatHandler(chatClient interfaces.ChatClient, helper *helper.Helper) *ChatHandler {
	logger, logFile := logging.InitLogrusLogger("./Logging/connectHub_gateway.log")
	return &ChatHandler{
		GRPC_Client: chatClient,
		helper:      helper,
		Logger:      logger,
		LogFile:     logFile,
	}
}
func (ch *ChatHandler) SendMessage(c *gin.Context) {

	tokenString := c.Request.Header.Get("Authorization")
	splitToken := strings.Split(tokenString, " ")
	if tokenString == "" {
		ch.Logger.Error("error on split token: ")
		println("error on split ")
		errs := response.ClientResponse(http.StatusUnauthorized, "Missing Authorization header", nil, "")
		c.JSON(http.StatusUnauthorized, errs)
		return
	}

	splitToken[1] = strings.TrimSpace(splitToken[1])
	splitToken[0] = strings.TrimSpace(splitToken[0])
	if splitToken[0] == "Jobseeker" {
		userID, err := ch.helper.ValidateTokenJobseeker(splitToken[1])

		if err != nil {
			ch.Logger.Error("Failed to Validate TokenJ obseeker: ", err)
			errs := response.ClientResponse(http.StatusUnauthorized, "Invalid token", nil, err.Error())
			c.JSON(http.StatusUnauthorized, errs)
			return
		}
		fmt.Println("upgrading ")
		conn, err := upgrade.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			ch.Logger.Error("Failed at upgrading: ", err)
			errs := response.ClientResponse(http.StatusBadRequest, "Websocket Connection Issue", nil, err.Error())
			c.JSON(http.StatusBadRequest, errs)
			return
		}

		defer delete(User, strconv.Itoa(userID))
		defer conn.Close()
		user := strconv.Itoa(userID)
		User[user] = conn

		for {
			fmt.Println("loop starts", userID, User)
			_, msg, err := conn.ReadMessage()
			if err != nil {
				ch.Logger.Error("Failed to Read Message: ", err)
				errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
				c.JSON(http.StatusBadRequest, errs)
				return
			}
			ch.helper.SendMessageToUser(User, msg, user)
		}
	} else if splitToken[0] == "Recruiter" {
		userID, err := ch.helper.ValidateTokenRecruiter(splitToken[1])
		fmt.Println("validate token result ", userID, err)
		if err != nil {
			ch.Logger.Error("Failed to Validate Token Recruiter: ", err)
			errs := response.ClientResponse(http.StatusUnauthorized, "Invalid token", nil, err.Error())
			c.JSON(http.StatusUnauthorized, errs)
			return
		}
		fmt.Println("upgrading ")
		conn, err := upgrade.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			ch.Logger.Error("Failed at upgrading: ", err)
			errs := response.ClientResponse(http.StatusBadRequest, "Websocket Connection Issue", nil, err.Error())
			c.JSON(http.StatusBadRequest, errs)
			return
		}

		defer delete(User, strconv.Itoa(userID))
		defer conn.Close()
		user := strconv.Itoa(userID)
		User[user] = conn

		for {
			fmt.Println("loop starts", userID, User)
			_, msg, err := conn.ReadMessage()
			if err != nil {
				ch.Logger.Error("Failed to Read Message: ", err)
				errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
				c.JSON(http.StatusBadRequest, errs)
				return
			}
			ch.helper.SendMessageToUser(User, msg, user)
			ch.Logger.Info("Send Message To User Successful")

		}
	} else {
		ch.Logger.Error("Invalid token role: ", errors.New("role is not specified"))
		errs := response.ClientResponse(http.StatusUnauthorized, "Invalid token role", nil, errors.New("role is not specified"))
		c.JSON(http.StatusUnauthorized, errs)
		return
	}

}

// GetChat handles the endpoint for retrieving chat details.
// @Summary Get Chat Details
// @Description Retrieves chat details based on the provided request.
// @Tags Recruiter Chat Management
// @Accept json
// @Produce json
// @Security BearerTokenAuth
// @Param id header integer true "User ID"
// @Param chatRequest body models.ChatRequest true "Chat Request Data"
// @Success 200 {object} response.Response "Chat details retrieved successfully"
// @Failure 400 {object} response.Response "Failed to retrieve chat details or incorrect data format"
// @Router /chat [post]
func (ch *ChatHandler) GetChatRecruiter(c *gin.Context) {

	var chatRequest models.ChatRequest
	if err := c.ShouldBindJSON(&chatRequest); err != nil {
		ch.Logger.Error("Failed to Get Data: ", err)

		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	userIDInterface, exists := c.Get("id")
	if !exists {
		ch.Logger.Error("User ID not found in JWT claims: ")
		errs := response.ClientResponse(http.StatusBadRequest, "User ID not found in JWT claims", nil, "")
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	userID := strconv.Itoa(userIDInterface.(int))
	result, err := ch.GRPC_Client.GetChat(userID, chatRequest)

	if err != nil {
		ch.Logger.Error("Failed to Get Data: ", err)
		errs := response.ClientResponse(http.StatusBadRequest, "Failed to get chat details", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	ch.Logger.Info("Successfully retrieved chat details")

	errs := response.ClientResponse(http.StatusOK, "Successfully retrieved chat details", result, nil)
	c.JSON(http.StatusOK, errs)
}


// GetChat handles the endpoint for retrieving chat details.
// @Summary Get Chat Details
// @Description Retrieves chat details based on the provided request.
// @Tags Jobseeker Chat Management
// @Accept json
// @Produce json
// @Security BearerTokenAuth
// @Param id header integer true "User ID"
// @Param chatRequest body models.ChatRequest true "Chat Request Data"
// @Success 200 {object} response.Response "Chat details retrieved successfully"
// @Failure 400 {object} response.Response "Failed to retrieve chat details or incorrect data format"
// @Router /chat [post]
func (ch *ChatHandler) GetChatJobseeker(c *gin.Context) {

	var chatRequest models.ChatRequest
	if err := c.ShouldBindJSON(&chatRequest); err != nil {
		ch.Logger.Error("Failed to Get Data: ", err)

		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	userIDInterface, exists := c.Get("id")
	if !exists {
		ch.Logger.Error("User ID not found in JWT claims: ")
		errs := response.ClientResponse(http.StatusBadRequest, "User ID not found in JWT claims", nil, "")
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	userID := strconv.Itoa(userIDInterface.(int))
	result, err := ch.GRPC_Client.GetChat(userID, chatRequest)

	if err != nil {
		ch.Logger.Error("Failed to Get Data: ", err)
		errs := response.ClientResponse(http.StatusBadRequest, "Failed to get chat details", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	ch.Logger.Info("Successfully retrieved chat details")

	errs := response.ClientResponse(http.StatusOK, "Successfully retrieved chat details", result, nil)
	c.JSON(http.StatusOK, errs)
}
