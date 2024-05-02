package routes

import "github.com/gin-gonic/gin"

func UserRoutes(r *gin.Engine) {
	UserGroup := r.Group("/api/user")
	UserGroup.GET("/", CreateUser)
}

func CreateUser(r *gin.Context) {
	r.JSON(200, gin.H{
		"message": "created user sucessfully",
	})
}
