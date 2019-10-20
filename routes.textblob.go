package main

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func showIndexPage(c *gin.Context) {
	navpayload := getAllPages()
	page := getRaws("home")

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
			"blobs":      template.HTML(page.Blobs[i]),
			"blobstitle": template.HTML(page.BlobsTitle[i]),
		})
	}
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
			"blobs":      template.HTML(page.Blobs[i]),
			"blobstitle": template.HTML(page.BlobsTitle[i]),
		})
	}

}
