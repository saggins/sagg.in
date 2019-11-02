package mc

import (
	"github.com/gin-gonic/gin"

	"github.com/minotar/minecraft"

	db "github.com/win32prog/sagg.in/web/app/db"
	model "github.com/win32prog/sagg.in/web/app/models"
	utlity "github.com/win32prog/sagg.in/web/app/utlity"


)

func MCwhitelist(username model.Postmsg, wJSON string, c *gin.Context) {

	minecrafts := minecraft.NewMinecraft()
	uuid, err := minecrafts.GetUUID(username.MCuser)
	if err == nil {
		msg := model.Whitelist{
			Mcuuid: uuid,
			Mcuser: username.MCuser,
		}
		
		db.PostNames(msg)
		//exsistingList := whitelistScan()

		//b, err := json.Marshal(exsistingList)
		//if err != nil {
		//	fmt.Println("Error when marshaling")
		//}

		//s := wJSON + "whitelist.json"
		//ioutil.WriteFile(s, b, os.ModePerm)
		RconWhitelist(msg)

		utlity.Render(c, "error.html", gin.H{
			"msg": "WOrkos! You are added",
		})

	} else {
		utlity.Render(c, "error.html", gin.H{
			"msg": "Username Not found... rip. Try again",
		})
	}

}
