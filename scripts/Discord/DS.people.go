package ds

import (
	"html/template"

	"github.com/gin-gonic/gin"

	utlity "github.com/win32prog/sagg.in/web/app/utlity"
	model "github.com/win32prog/sagg.in/web/app/models"
	db 	"github.com/win32prog/sagg.in/web/app/db"

)

func DSPlayers(c *gin.Context){
	
	id := c.Param("discord_id")
	avatar := c.Param("discord_avatar")

	c.SetCookie("id", id, 3600, "/", "sagg.in", true, false)
	c.SetCookie("avatar", avatar, 3600, "/", "sagg.in", true, false)


	page:= db.GetRaws("discord")

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

func DSResult(c *gin.Context){

	id ,err := c.Cookie("id")
	avatar ,err := c.Cookie("avatar")

	if err != nil {
		utlity.Render(c, "error.html", gin.H{
			"msg":"Please Go through Disocrd $reg",
		})
	}else{
	var Mcname model.DiscordUser
	c.ShouldBind(&Mcname)
	url:=id+"/"+avatar

	newperson := model.Player{
		Userid: url,
		Name: Mcname.Name,
		Ip: c.ClientIP(),
	}
	
	db.PostPlayers(newperson)

	var listOpeople []model.Player
	listOpeople = db.DiscordScan()
	utlity.Render(c, "page.html", gin.H{
			})
	utlity.Render(c, "players.html", gin.H{
		"payload":listOpeople,
	})
	}

	
}
func Listplayers(c *gin.Context){
	var listOpeople []model.Player
	listOpeople = db.DiscordScan()
	utlity.Render(c, "page.html", gin.H{
			})
	utlity.Render(c, "players.html", gin.H{
		"payload":listOpeople,
	})
}
