package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	msg "github.com/ARunni/Error_Message"
	"github.com/ARunni/connectHub_gateway/pkg/helper"
	"github.com/ARunni/connectHub_gateway/pkg/utils/response"
	"github.com/gin-gonic/gin"
)

func JobseekerAuthMiddleware(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		err := errors.New("field empty")

		response := response.ClientResponse(http.StatusUnauthorized, msg.NoAuth, nil, err.Error())
		c.JSON(http.StatusUnauthorized, response)
		c.Abort()
		return
	}

	splited := strings.Split(tokenString, " ")

	if len(splited) != 2 {
		fmt.Println("")
		err := errors.New("format not satisfied")
		response := response.ClientResponse(http.StatusUnauthorized, "error in splitting", nil, err.Error())
		c.JSON(http.StatusUnauthorized, response)
		c.Abort()
		return
	}

	tokenPart2 := splited[1]
	tokenPart1 := splited[0]

	if tokenPart1 != "Jobseeker" {
		err := errors.New("role mismatch")
		response := response.ClientResponse(http.StatusUnauthorized, "provided Role is not Jobseeker ", nil, err.Error())
		c.JSON(http.StatusUnauthorized, response)
		c.Abort()
		return
	}

	tokenclaims, err := helper.ValidateTokenJobSeeker(tokenPart2)
	if err != nil {
		response := response.ClientResponse(http.StatusUnauthorized, "Invalid Token Jobseeker", nil, err.Error())
		c.JSON(http.StatusUnauthorized, response)
		c.Abort()
		return
	}
	if tokenclaims.Role != "jobseeker" {
		err := errors.New("invalid role")
		response := response.ClientResponse(http.StatusUnauthorized, "Invalid Token Jobseeker", nil, err.Error())
		c.JSON(http.StatusUnauthorized, response)
		c.Abort()
		return
	}

	c.Set("id", int(tokenclaims.ID))
	c.Set("role", string(tokenclaims.Role))
	c.Next()
}
