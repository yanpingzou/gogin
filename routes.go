package main

import (
	"ginblog/handlers"
)

func initializeRoutes(){
	router.GET("/index", handlers.ShowPageIndex)
	router.GET("/article/view/:article_id", handlers.GetArticle)
}