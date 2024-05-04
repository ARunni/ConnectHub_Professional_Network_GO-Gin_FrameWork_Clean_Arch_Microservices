package handler

import (
	interfaces "connectHub_gateway/pkg/client/interface"
	"connectHub_gateway/pkg/utils/models"
	"connectHub_gateway/pkg/utils/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	GRPC_Client interfaces.AdminClient
}

func NewAdminHandler(grpc_client interfaces.AdminClient) *AdminHandler {
	return &AdminHandler{
		GRPC_Client: grpc_client,
	}
}
func (ah *AdminHandler) LoginHandler(c *gin.Context) {
	var adminData models.AdminLogin

	if err := c.ShouldBindJSON(&adminData); err != nil {
		errResp := response.ClientResponse(http.StatusBadRequest, "Incorrect Format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errResp)
		return
	}
	admin, err := ah.GRPC_Client.AdminLogin(adminData)
	if err != nil {
		errResp := response.ClientResponse(http.StatusInternalServerError, "Cannot authenticate Admin", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errResp)
		return
	}
	successResp := response.ClientResponse(http.StatusOK, "Admin Authenticated Successfully", admin, nil)
	c.JSON(http.StatusOK, successResp)

}
