package middleware

// import (
// 	"errors"
// 	"fmt"
// 	"net/http"
// 	"strings"

// 	"github.com/ARunni/connectHub_gateway/pkg/helper"
// 	"github.com/ARunni/connectHub_gateway/pkg/utils/response"

// 	msg "github.com/ARunni/Error_Message"
// 	"github.com/gin-gonic/gin"
// )

// func AuthMiddleware(c *gin.Context) {
// 	tokenString := c.GetHeader("Authorization")
// 	if tokenString == "" {
// 		err := errors.New("field empty")

// 		response := response.ClientResponse(http.StatusUnauthorized, msg.NoAuth, nil, err.Error())
// 		c.JSON(http.StatusUnauthorized, response)
// 		c.Abort()
// 		return
// 	}

// 	splited := strings.Split(tokenString, " ")

// 	if len(splited) != 3 {
// 		fmt.Println("")
// 		err := errors.New("format not satisfied")
// 		response := response.ClientResponse(http.StatusUnauthorized, "error in splitting", nil, err.Error())
// 		c.JSON(http.StatusUnauthorized, response)
// 		c.Abort()
// 		return
// 	}

// 	tokenPart1 := splited[1]
// 	tokenPart2 := splited[2]

// 	switch tokenPart1 {
// 	case "Admin":
// 		tokenclaims, err := helper.ValidateTokenAdmin(tokenPart2)
// 		if err != nil {
// 			response := response.ClientResponse(http.StatusUnauthorized, "Invalid Token admin", nil, err.Error())
// 			c.JSON(http.StatusUnauthorized, response)
// 			c.Abort()
// 			return
// 		}
// 		if tokenclaims.Role != "admin" {
// 			err := errors.New("invalid role")
// 			response := response.ClientResponse(http.StatusUnauthorized, "Invalid Token admin", nil, err.Error())
// 			c.JSON(http.StatusUnauthorized, response)
// 			c.Abort()
// 			return
// 		}
// 		c.Set("tokenClaims", tokenclaims)

// 		c.Next()
// 	case "Jobseeker":
// 		tokenclaims, err := helper.ValidateTokenJobSeeker(tokenPart2)
// 		if err != nil {
// 			response := response.ClientResponse(http.StatusUnauthorized, "Invalid Token Jobseeker", nil, err.Error())
// 			c.JSON(http.StatusUnauthorized, response)
// 			c.Abort()
// 			return
// 		}
// 		if tokenclaims.Role != "jobseeker" {
// 			err := errors.New("invalid role")
// 			response := response.ClientResponse(http.StatusUnauthorized, "Invalid Token Jobseeker", nil, err.Error())
// 			c.JSON(http.StatusUnauthorized, response)
// 			c.Abort()
// 			return
// 		}

// 		c.Set("id", int(tokenclaims.ID))
// 		c.Set("role", string(tokenclaims.Role))
// 		c.Next()

// 	case "Recruiter":
// 		tokenclaims, err := helper.ValidateTokenRecruiter(tokenPart2)
// 		if err != nil {
// 			response := response.ClientResponse(http.StatusUnauthorized, "Invalid Token Recruiter", nil, err.Error())
// 			c.JSON(http.StatusUnauthorized, response)
// 			c.Abort()
// 			return
// 		}
// 		if tokenclaims.Role != "recruiter" {
// 			err := errors.New("invalid role")
// 			response := response.ClientResponse(http.StatusUnauthorized, "Invalid Token Recruiter", nil, err.Error())
// 			c.JSON(http.StatusUnauthorized, response)
// 			c.Abort()
// 			return
// 		}
// 		c.Set("id", int(tokenclaims.ID))
// 		c.Set("role", string(tokenclaims.Role))
// 		c.Next()

// 	default:
// 		err := errors.New("privileges not met")
// 		response := response.ClientResponse(http.StatusUnauthorized, "Invalid Token format", nil, err.Error())
// 		c.JSON(http.StatusUnauthorized, response)
// 		c.Abort()
// 		return
// 	}

// }


