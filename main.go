package main

import (
	"gogin/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/static"
)

var (
	router *gin.Engine
)

func main() {
	router = gin.Default()
	router.Use(static.Serve("/static", static.LocalFile("static", true)))
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.LoadHTMLGlob("templates/*")
	initializeRoutes()
	models.InitConnect()
	router.Run(":8080")
}
