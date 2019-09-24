package main

import (
	"web-server/handlers"
)

func initializeRoutes() {
	router.GET("/", handlers.ShowIndexPage)
	router.GET("/article/view/:article_id", handlers.GetArticle)
	userRoutes := router.Group("/u")
	{
		userRoutes.GET("/register", handlers.ShowRegistrationPage)
		userRoutes.POST("/register", handlers.Register)
		userRoutes.GET("/user", handlers.ShowUserPage)
		userRoutes.POST("/logout", handlers.Logout)
	}
}
