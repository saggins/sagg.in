package mc
import (
	"github.com/gin-gonic/gin"

	utlity "github.com/win32prog/sagg.in/web/app/utlity"

	db 	"github.com/win32prog/sagg.in/web/app/db"
	
	model "github.com/win32prog/sagg.in/web/app/models"
)

func MCPlayers(c *gin.Context){
	var listOplayers []model.Whitelist
	listOplayers = db.WhitelistScan()
	utlity.Render(c, "page.html", gin.H{
			})
	utlity.Render(c, "mcnames.html", gin.H{
		"payload":listOplayers,
	})
	
} 