package routes

import (
	handler "github.com/ARunni/connectHub_gateway/pkg/api/handler/auth"
	"github.com/ARunni/connectHub_gateway/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func AdminRoutes(engine *gin.RouterGroup, adminhandler *handler.AdminHandler) {

	engine.POST("/login", adminhandler.AdminLogin)

	engine.Use(middleware.AdminAuthMiddleware)
	{
		engine.POST("/policy", adminhandler.CreatePolicy)
		engine.PUT("/policy", adminhandler.UpdatePolicy)
		engine.DELETE("/policy", adminhandler.DeletePolicy)
		engine.GET("/policy", adminhandler.GetOnePolicy)
		engine.GET("/policies", adminhandler.GetAllPolicies)

		// Admin Routes jobseeker
		adminJobseeker := engine.Group("/jobseeker")

		adminJobseeker.GET("/all", adminhandler.GetJobseekers)
		adminJobseeker.PATCH("/block", adminhandler.BlockJobseeker)
		adminJobseeker.PATCH("/unblock", adminhandler.UnBlockJobseeker)
		adminJobseeker.GET("", adminhandler.GetJobseekerDetails)

		// Admin Routes recruiter
		adminRecruiter := engine.Group("/recruiter")
		
		adminRecruiter.GET("/all", adminhandler.GetRecruiters)
		adminRecruiter.PATCH("/block", adminhandler.BlockRecruiter)
		adminRecruiter.PATCH("/unblock", adminhandler.UnBlockRecruiter)
		adminRecruiter.GET("", adminhandler.GetRecruiterDetails)
	}
}
