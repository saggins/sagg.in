package routes

import (
	"html/template"

	"github.com/gin-gonic/gin"
	
	db "github.com/win32prog/sagg.in/web/app/db"
	utlity "github.com/win32prog/sagg.in/web/app/utlity"
)

func showIndexPage(c *gin.Context) {

	//just fo nav... we love nav to be loopy

	utlity.Render(c, "home.html", gin.H{
	})

}

func getPage(c *gin.Context) {
	page := db.GetRaws(c.Param("page_id"))

	//just fo nav... we love nav to be loopy
	utlity.Render(c, "page.html", gin.H{
		"title":   page.Title,
		"payload": page,
	})

	for i := 0; i < len(page.Blobs); i++ {
		utlity.Render(c, "blobs.html", gin.H{
			"blobs":      template.HTML(page.Blobs[i]),
			"blobstitle": template.HTML(page.BlobsTitle[i]),
		})
	}

}
