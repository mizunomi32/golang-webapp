package main

import (
	"fmt"
	"github.com/gin-gonic/gin"

  "./config"
	"./controllers"
)

func main() {
	fmt.Println("Hello world!")
	fmt.Printf("AppName:\t%s\n", config.AppName)
	fmt.Printf("Project is start at http://127.0.0.0%s/\n", config.Port)

	router := gin.Default()
	router.LoadHTMLGlob("views/*.html")
	router.Static("/assets", "./assets")

	router.GET("/", controllers.Index)
	router.GET("/test", controllers.Test)

	router.NoRoute(controllers.NoRoute)

	router.Run(config.Port)
}
