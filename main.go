package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	initializeRoutes(router)
	//router.GET("/", func(c *gin.Context) {
	//	c.HTML(http.StatusOK, "index.html", gin.H{
	//		"title": "Sagg Web!"})
	//})
	log.Fatal(autotls.Run(router, "sagg.in"))
}
func render(c *gin.Context, templateName string, data gin.H) {
	switch c.Request.Header.Get("Accept") {
	case "application/json":
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		c.XML(http.StatusOK, data["payload"])
	default:
		c.HTML(http.StatusOK, templateName, data)
	}
}
