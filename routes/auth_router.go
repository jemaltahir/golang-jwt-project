package routes

import "github.com/gin-gonic/gin"

func AuthRoutes(incoming Routes *gin.Engine) {
	incoming.POST("users/signup", controller.Signup())
	incoming.POST("users/login", controller.Login())
}