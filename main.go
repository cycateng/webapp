package main

import (
	"./routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLFiles("views/*.html")
	router.Static("/assets", "./assets")

	user := router.Group("/user")
	{
		user.POST("/signup", routes.UserSignUp)
		user.POST("/login", routes.UserLogIn)
	}

	router.GET("/", routes.Home)
	route.GET("/login", routes.LogIn)
	route.GET("/signup", routes.SignUp)
	router.NoRoute(routes.NoRoute)

	router.Run(":80")
}