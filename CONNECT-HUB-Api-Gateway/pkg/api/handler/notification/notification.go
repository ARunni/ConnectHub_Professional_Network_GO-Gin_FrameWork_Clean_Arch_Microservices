package handler

import (
	"net/http"
	"os"
	"strconv"

	logging "github.com/ARunni/connectHub_gateway/Logging"
	"github.com/ARunni/connectHub_gateway/pkg/client/notification/interfaces"
	"github.com/ARunni/connectHub_gateway/pkg/utils/models"
	"github.com/ARunni/connectHub_gateway/pkg/utils/response"
	"github.com/gin-gonic/gin"

	"github.com/sirupsen/logrus"
)

type NotificationHandler struct {
	GRPC_Client interfaces.NotificationClient
	Logger      *logrus.Logger
	LogFile     *os.File
}

func NewNotificationHandler(grpc_client interfaces.NotificationClient) *NotificationHandler {
	logger, logFile := logging.InitLogrusLogger("./Logging/connectHub_gateway.log")
	return &NotificationHandler{
		GRPC_Client: grpc_client,
		Logger:      logger,
		LogFile:     logFile,
	}
}


// GetNotification retrieves notifications for a user with pagination support.
// @Summary Get notifications
// @Description Retrieves notifications for a user with pagination support
// @Tags Jobseeker Notification Management
// @Accept json
// @Produce json
// @Security BearerTokenAuth
// @Param limit query int false "Number of notifications to retrieve (default: 1)"
// @Param offset query int false "Offset for pagination (default: 10)"
// @Success 200 {object} response.Response "Successfully retrieved notifications"
// @Failure 400 {object} response.Response "Failed to retrieve notifications: missing or incorrect parameters"
// @Failure 500 {object} response.Response "Internal server error: failed to retrieve notifications"
// @Router /notifications [get]
func (n *NotificationHandler) GetNotification(c *gin.Context) {
	n.Logger.Info("GetNotification at NotificationHandler started")

	pageStr := c.DefaultQuery("limit", "1")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		n.Logger.Error("page number not in right format", err)
		errorRes := response.ClientResponse(http.StatusBadRequest, "page number not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	countStr := c.DefaultQuery("offset", "10")
	pageSize, err := strconv.Atoi(countStr)
	if err != nil {
		n.Logger.Error("user count in a page not in right format", err)
		errorRes := response.ClientResponse(http.StatusBadRequest, "user count in a page not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	notificationRequest := models.NotificationPagination{
		Limit:  page,
		Offset: pageSize,
	}

	userID, exists := c.Get("id")
	if !exists {
		n.Logger.Error("User ID not found in JWT claims")
		errs := response.ClientResponse(http.StatusBadRequest, "User ID not found in JWT claims", nil, "")
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	UserID, _ := userID.(int)

	data, err := n.GRPC_Client.GetNotification(UserID, notificationRequest)
	if err != nil {
		n.Logger.Error("Error during GetNotification rpc call", err)
		errs := response.ClientResponse(http.StatusBadRequest, "Failed to get notification details", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	n.Logger.Info("Successfully retrieved Notifications")
	res := response.ClientResponse(http.StatusOK, "Successfully retrieved Notifications", data, nil)
	c.JSON(http.StatusOK, res)
}


// ReadNotification marks a notification as read.
// @Summary Read notification
// @Description Marks a notification as read
// @Tags Jobseeker Notification Management
// @Accept json
// @Produce json
// @Security BearerTokenAuth
// @Param notification_id query string true "Notification ID"
// @Success 200 {object} response.Response "Notification successfully marked as read"
// @Failure 400 {object} response.Response "Failed to mark notification as read: missing or incorrect parameters"
// @Failure 500 {object} response.Response "Internal server error: failed to mark notification as read"
// @Router /notifications [patch]
func (n *NotificationHandler) ReadNotification(c *gin.Context) {
	noificationIdStr := c.Query("notification_id")
	noificationId, err := strconv.Atoi(noificationIdStr)
	if err != nil {
		n.Logger.Error("conversion error notification id ", err)
		errs := response.ClientResponse(http.StatusBadRequest, "conversion error notification id ", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	userId, ok := c.Get("id")
	if !ok {
		n.Logger.Error("User ID not found in JWT claims")
		errs := response.ClientResponse(http.StatusBadRequest, "User ID not found in JWT claims", nil, "")
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	UserID, _ := userId.(int)

	result, err := n.GRPC_Client.ReadNotification(noificationId, UserID)
	if err != nil {
		n.Logger.Error("Error during ReadNotification rpc call", err)
		errs := response.ClientResponse(http.StatusBadRequest, "Failed to read notification ", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	n.Logger.Info(" Notification Successfully marked as read")
	res := response.ClientResponse(http.StatusOK, "Notification Successfully marked as read", result, nil)
	c.JSON(http.StatusOK, res)

}


// MarkAllAsRead marks all notifications as read for a user.
// @Summary Mark all notifications as read
// @Description Marks all notifications as read for a user
// @Tags Jobseeker Notification Management
// @Accept json
// @Produce json
// @Security BearerTokenAuth
// @Success 200 {object} response.Response "Notifications successfully marked as read"
// @Failure 400 {object} response.Response "Failed to mark notifications as read: User ID not found or incorrect parameters"
// @Failure 500 {object} response.Response "Internal server error: failed to mark notifications as read"
// @Router /notifications/all [patch]
func (n *NotificationHandler) MarkAllAsRead(c *gin.Context) {
	n.Logger.Info("MarkAllAsRead at NotificationHandler started")

	userId, ok := c.Get("id")
	if !ok {
		n.Logger.Error("User ID not found in JWT claims")
		errs := response.ClientResponse(http.StatusBadRequest, "User ID not found in JWT claims", nil, "")
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	UserID, _ := userId.(int)
	n.Logger.Info("MarkAllAsRead at client rpc  started")
	result, err := n.GRPC_Client.MarkAllAsRead(UserID)
	if err != nil {
		n.Logger.Error("Error during ReadNotification rpc call", err)
		errs := response.ClientResponse(http.StatusBadRequest, "Failed to read notification ", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	n.Logger.Info("MarkAllAsRead at client rpc  finished")
	n.Logger.Info(" Notifications Successfully marked as read")
	res := response.ClientResponse(http.StatusOK, "Notification Successfully marked as read", result, nil)
	c.JSON(http.StatusOK, res)

}




// GetAllNotifications retrieves all notifications for a user.
// @Summary Get all notifications
// @Description Retrieves all notifications for a user
// @Tags Jobseeker Notification Management
// @Accept json
// @Produce json
// @Security BearerTokenAuth
// @Success 200 {object} response.Response "Successfully retrieved all notifications"
// @Failure 400 {object} response.Response "Failed to retrieve notifications: User ID not found or incorrect parameters"
// @Failure 500 {object} response.Response "Internal server error: failed to retrieve notifications"
// @Router /notifications/all [get]
func (n *NotificationHandler) GetAllNotifications(c *gin.Context) {
	n.Logger.Info("GetAllNotifications at NotificationHandler started")

	userId, ok := c.Get("id")
	if !ok {
		n.Logger.Error("User ID not found in JWT claims")
		errs := response.ClientResponse(http.StatusBadRequest, "User ID not found in JWT claims", nil, "")
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	UserID, _ := userId.(int)
	n.Logger.Info("GetAllNotifications at client rpc  started")
	result, err := n.GRPC_Client.GetAllNotifications(UserID)
	if err != nil {
		n.Logger.Error("Error during GetAllNotifications rpc call", err)
		errs := response.ClientResponse(http.StatusBadRequest, "Failed to GetAllNotifications ", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	n.Logger.Info("GetAllNotifications at client rpc  finished")
	n.Logger.Info(" Notifications Successfully Fetched")
	res := response.ClientResponse(http.StatusOK, "Get All Notifications Successfull ", result, nil)
	c.JSON(http.StatusOK, res)

}
