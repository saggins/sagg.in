package main

import (
	"net/http"

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
	//router.Run("172.31.18.164:40000")
	http.ListenAndServeTLS(":443", "/etc/letsencrypt/live/sagg.in/fullchain.pem", "/etc/letsencrypt/live/sagg.in/privkey.pem", router)
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
