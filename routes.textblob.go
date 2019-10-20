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

func getPage(c *gin.Context) {
	navpayload := getAllPages()
	page := getRaws(c.Param("page_id"))

	//just fo nav... we love nav to be loopy
	c.HTML(http.StatusOK, "nav.html", gin.H{
		"payload": navpayload,
	})

	render(c, "page.html", gin.H{
		"title":   page.Title,
		"payload": page,
	})

	for i := 0; i < len(page.Blobs); i++ {
		render(c, "blobs.html", gin.H{
			"blobs":      page.Blobs[i],
			"blobstitle": page.BlobsTitle[i],
		})
	}

}
