package main

import (
	"github.com/ayushkkkkkkk/fast-delivery/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routes.UserRoutes(r)
	r.Run() // listen and serve on 0.0.0.0:8080
}

