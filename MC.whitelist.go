package main

import (
	"github.com/gin-gonic/gin"
	"github.com/minotar/minecraft"
)

type Whitelist struct {
	Mcuuid string `json:"uuid"`
	Mcuser string `json:"name"`
}

func MCwhitelist(username postmsg, wJSON string, c *gin.Context) {

	minecrafts := minecraft.NewMinecraft()
	uuid, err := minecrafts.GetUUID(username.MCuser)
	if err == nil {
		msg := Whitelist{
			Mcuuid: uuid,
			Mcuser: username.MCuser,
		}

		postNames(msg)
		exsistingList := whitelistScan()

		//b, err := json.Marshal(exsistingList)
		//if err != nil {
		//	fmt.Println("Error when marshaling")
		//}

		//s := wJSON + "whitelist.json"
		//ioutil.WriteFile(s, b, os.ModePerm)
		RconWhitelist(exsistingList)

		render(c, "error.html", gin.H{
			"msg": "WOrkos! You are added",
		})

	} else {
		render(c, "error.html", gin.H{
			"msg": "Username Not found... rip. Try again",
		})
	}

}
