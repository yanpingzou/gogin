package main

import (
	"gogin/models"
	"gopkg.in/gin-gonic/gin.v1"
)

var (
	router *gin.Engine
)

func main() {
	router = gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.LoadHTMLGlob("templates/*")
	initializeRoutes()
	models.InitConnect()
	router.Run(":8080")
}
