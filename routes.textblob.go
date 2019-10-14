package main

import (
	"github.com/gin-gonic/gin"
)

func showIndexPage(c *gin.Context) {

	render(c, gin.H{
		"title": "Sagg Web!",
	}, "index.html")
}

