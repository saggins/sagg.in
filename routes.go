package main

import "github.com/gin-gonic/gin"

func initializeRoutes(r *gin.Engine) {
	// INDEX is special... Its a special boi
	r.GET("/", showIndexPage)

	pageroutes := r.Group("/page")
	{
		pageroutes.GET("/view/:page_id", getPage)
	}
}
