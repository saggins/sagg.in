package main
import (
	"github.com/gin-gonic/gin"
)

func MCPlayers(c *gin.Context){

	listOplayers := whitelistScan()
	render(c, "page.html", gin.H{
			})
	render(c, "mcnames.html", gin.H{
		"payload":listOplayers,
	})
}