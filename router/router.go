package router

import (
	"github.com/gin-gonic/gin"
	"jwt-practice/api"
	"jwt-practice/api/user"
	"jwt-practice/middleware"
)

func InitRouter() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	auth := router.Group("/user")
	{
		auth.POST("/register", user.Register)
		auth.POST("/auth", user.Auth)
		//auth.PUT("/emailVerification", user.VerifyEmail)

	}
	// v1
	apiV1 := router.Group("/api/v1")
	apiV1.Use(middleware.JwtAuth(), middleware.ErrorHandler())
	{
		apiV1.GET("/home", api.GetHome)
	}



	router.Run(":8080")
}