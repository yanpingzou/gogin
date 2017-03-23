package main

import (
	"gopkg.in/gin-gonic/gin.v1"
)

var router *gin.Engine

func main() {
	router = gin.Default()
	router.LoadHTMLGlob("templates/*")
	initializeRoutes()
	router.Run(":8080")
}
