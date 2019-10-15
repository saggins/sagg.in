package main

import "github.com/gin-gonic/gin"

func initializeRoutes(r *gin.Engine) {
	// INDEX is special... Its a special boi
	router.GET("/", showIndexPage)
}
