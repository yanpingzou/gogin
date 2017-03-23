package main

import (
	"gogin/handlers"
)

func initializeRoutes(){
	router.GET("/index", handlers.ShowPageIndex)
	router.GET("/article/view/:article_id", handlers.GetArticle)
	router.GET("/cookie", handlers.SetCookie)
	router.GET("/user", handlers.NewUser)
}