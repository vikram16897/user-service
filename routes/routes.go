package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/user-service/controller"
)

func RegisterHealthRoute(r *gin.Engine) {
	r.GET("health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "App is working fine",
		})
	})
}

func RegisterUserRoute(r *gin.Engine, userController *controller.UserController) {
	userGroup := r.Group("v1/users")
	userGroup.POST("/", userController.CreateUser)
	userGroup.GET("/all", userController.GetAllUsers)
	userGroup.GET("/:id", userController.GetUserByID)
	userGroup.GET("/byemail", userController.GetUsersByEmail)
	userGroup.POST("/update", userController.UpdateUser)
}
