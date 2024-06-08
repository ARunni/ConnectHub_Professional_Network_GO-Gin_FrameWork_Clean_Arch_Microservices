package handler

import (
	logging "connectHub_gateway/Logging"
	interfaces "connectHub_gateway/pkg/client/chat/interfaces"
	"connectHub_gateway/pkg/helper"
	"connectHub_gateway/pkg/utils/models"
	"connectHub_gateway/pkg/utils/response"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrade = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var User = make(map[string]*websocket.Conn)

type ChatHandler struct {
	GRPC_Client interfaces.ChatClient
	helper      *helper.Helper
}

func NewChatHandler(chatClient interfaces.ChatClient, helper *helper.Helper) *ChatHandler {
	return &ChatHandler{
		GRPC_Client: chatClient,
		helper:      helper,
	}
}
func (ch *ChatHandler) SendMessage(c *gin.Context) {

	logrusLogger, logrusLogFile := logging.InitLogrusLogger("./Logging/connectHub_gateway.log")
	defer logrusLogFile.Close()

	tokenString := c.Request.Header.Get("Authorization")
	splitToken := strings.Split(tokenString, " ")
	if tokenString == "" {
		logrusLogger.Error("error on split token: ")
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
			logrusLogger.Error("Failed to Validate TokenJ obseeker: ", err)
			errs := response.ClientResponse(http.StatusUnauthorized, "Invalid token", nil, err.Error())
			c.JSON(http.StatusUnauthorized, errs)
			return
		}
		fmt.Println("upgrading ")
		conn, err := upgrade.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			logrusLogger.Error("Failed at upgrading: ", err)
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
				logrusLogger.Error("Failed to Read Message: ", err)
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
			logrusLogger.Error("Failed to Validate Token Recruiter: ", err)
			errs := response.ClientResponse(http.StatusUnauthorized, "Invalid token", nil, err.Error())
			c.JSON(http.StatusUnauthorized, errs)
			return
		}
		fmt.Println("upgrading ")
		conn, err := upgrade.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			logrusLogger.Error("Failed at upgrading: ", err)
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
				logrusLogger.Error("Failed to Read Message: ", err)
				errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
				c.JSON(http.StatusBadRequest, errs)
				return
			}
			ch.helper.SendMessageToUser(User, msg, user)
			logrusLogger.Info("Send Message To User Successful")

		}
	} else {
		logrusLogger.Error("Invalid token role: ", errors.New("role is not specified"))
		errs := response.ClientResponse(http.StatusUnauthorized, "Invalid token role", nil, errors.New("role is not specified"))
		c.JSON(http.StatusUnauthorized, errs)
		return
	}

}

func (ch *ChatHandler) GetChat(c *gin.Context) {

	logrusLogger, logrusLogFile := logging.InitLogrusLogger("./Logging/connectHub_gateway.log")
	defer logrusLogFile.Close()

	var chatRequest models.ChatRequest
	if err := c.ShouldBindJSON(&chatRequest); err != nil {
		logrusLogger.Error("Failed to Get Data: ", err)

		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	userIDInterface, exists := c.Get("id")
	if !exists {
		logrusLogger.Error("User ID not found in JWT claims: ")
		errs := response.ClientResponse(http.StatusBadRequest, "User ID not found in JWT claims", nil, "")
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	userID := strconv.Itoa(userIDInterface.(int))
	result, err := ch.GRPC_Client.GetChat(userID, chatRequest)

	if err != nil {
		logrusLogger.Error("Failed to Get Data: ", err)
		errs := response.ClientResponse(http.StatusBadRequest, "Failed to get chat details", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	logrusLogger.Info("Successfully retrieved chat details")

	errs := response.ClientResponse(http.StatusOK, "Successfully retrieved chat details", result, nil)
	c.JSON(http.StatusOK, errs)
}