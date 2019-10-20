package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func showIndexPage(c *gin.Context) {
	payload := getAllPages()
	c.HTML(http.StatusOK, "nav.html", gin.H{
		"payload": payload,
	})
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title":   "Sagg Web!",
		"payload": payload,
	})

}
