package main

import (
	"github.com/ayushkkkkkkk/fast-delivery/models"
	"github.com/ayushkkkkkkk/fast-delivery/routes"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	r := gin.Default()
	err := models.InitDB()

	if err != nil {
		panic(err)
	}

	routes.UserRoutes(r)
	r.Run() // listen and serve on 0.0.0.0:8080
}
