package main

import (
	"github.com/benreic/dingus/config"
	"github.com/benreic/dingus/controllers"
	"github.com/benreic/dingus/middleware"
	"github.com/benreic/dingus/utils"
	"github.com/gin-gonic/gin"
)

func main() {

	utils.Log("Dingus starting")

	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	router.LoadHTMLGlob("templates/*.tmpl")
	router.Static("/public", "./public")

	public := router.Group("/")
	{
		public.GET("/", controllers.Hello)
		public.GET("/sign-up", controllers.SignUp)
		public.POST("/sign-up", controllers.ProcessSignUp)
		public.POST("/login", controllers.Login)
		public.GET("/automatic-redirect", controllers.AutomaticRedirect)
	}

	protected := router.Group("/", middleware.AuthRequired)
	{
		protected.GET("/dashboard", controllers.Dashboard)
		protected.POST("/logout", controllers.Logout)
	}

	// Listen and server on 0.0.0.0:8080
	err := router.Run(":" + config.Port())

	utils.Log("Dingus stopping")

	if err != nil {
		utils.LogError(err)
	}
}
