package routes

import (
	"github.com/ayushkkkkkkk/fast-delivery/models"
	"github.com/gin-gonic/gin"
)

type User struct {
	Name   string `json:"name"`
	Email  string `json:"email"`
	Gender string `json:"gender"`
	DOB    string `json:"dob"`
}

func UserRoutes(r *gin.Engine) {
	UserGroup := r.Group("/api/user")
	UserGroup.POST("/new", CreateUser)
	UserGroup.GET("/:id", getSingleUser)
}

func getSingleUser(r *gin.Context) {
	userId := r.Param("id")
	var user User
	err := models.Db.QueryRow("SELECT username, email FROM users WHERE id = ?", userId).Scan(&user.Name, &user.Email)
	if err != nil {
		r.JSON(404, gin.H{"error": "User not found"})
		return
	}

	r.JSON(200, gin.H{
		"user": user,
	})

}

func CreateUser(r *gin.Context) {
	var user User
	if err := r.ShouldBindJSON(&user); err != nil {
		r.JSON(400, gin.H{"error": err.Error()})
		return
	}

	_, err := models.Db.Exec(`
        INSERT INTO userg ( name, email, gender, dob)
        VALUES (?, ?, ?, ?)
    `, user.Name, user.Email, user.Gender, user.DOB)

	if err != nil {
		r.JSON(201, gin.H{
			"message": err.Error(),
		})
		return
	}

	r.JSON(200, gin.H{
		"message": "Sucessfully added the user into the database",
		"name":    user.Name,
		"email":   user.Email,
		"Gender":  user.Gender,
		"dob":     user.DOB,
	})
}
