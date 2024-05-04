package routes

import (
	"github.com/ayushkkkkkkk/fast-delivery/models"
	"github.com/gin-gonic/gin"
)

type User struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

func UserRoutes(r *gin.Engine) {
	UserGroup := r.Group("/api/user")
	UserGroup.POST("/new", CreateUser)
	UserGroup.GET("/SingleUser/:email", GetUser)
}

func CreateUser(r *gin.Context) {
	var user User

	if err := r.ShouldBindJSON(&user); err != nil {
		r.JSON(400, gin.H{"error": err.Error()})
		return
	}

	_, err := models.Db.Exec("INSERT INTO users (username, email) VALUES (?, ?)", user.Username, user.Email)
	if err != nil {
		r.JSON(500, gin.H{"error": "Failed to insert user into database", "E": err, "data": user.Email, "another": user.Username})
		return
	}
	r.JSON(200, gin.H{
		"message": "User created successfully",
		"user":    user,
	})
}

func GetUser(r *gin.Context) {
	userEmail := r.Param("email")

	var user User
	err := models.Db.QueryRow("SELECT username, email FROM users WHERE email = ?", userEmail).Scan(&user.Username, &user.Email)
	if err != nil {
		r.JSON(404, gin.H{"error": "User not found"})
		return
	}

	r.JSON(200, gin.H{
		"user": user,
	})

}
