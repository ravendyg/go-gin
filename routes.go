package main

import (
	"web-server/handlers"
	"web-server/middlewares"
)

func initializeRoutes() {
	router.GET("/", handlers.ShowIndexPage)
	articleRoutes := router.Group("/article")
	{
		articleRoutes.GET("/view/:article_id", handlers.GetArticle)
		articleRoutes.GET("/new", middlewares.GetUser(), middlewares.RequireAuth(), handlers.GetNewArticleForm)
		articleRoutes.POST("/new", middlewares.GetUser(), middlewares.RequireAuth(), handlers.CreateArticle)
	}
	userRoutes := router.Group("/u")
	{
		userRoutes.GET("/register", middlewares.GetUser(), handlers.ShowRegistrationPage)
		userRoutes.POST("/register", handlers.Register)
		userRoutes.GET("/login", handlers.ShowLoginPage)
		userRoutes.POST("/login", handlers.Login)
		userRoutes.POST("/logout", handlers.Logout)
	}
	router.GET("/user", middlewares.GetUser(), middlewares.RequireAuth(), handlers.ShowUserPage)
}
