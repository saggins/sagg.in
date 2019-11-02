package routes

import (	
	"github.com/gin-gonic/gin"
	
	mc "github.com/win32prog/sagg.in/scripts/MC"

	model "github.com/win32prog/sagg.in/web/app/models"
)



func MCName(c *gin.Context) {
	var Mcname model.Postmsg
	c.ShouldBind(&Mcname)

	mc.MCwhitelist(Mcname, "/home/saggins/Documents/projects/test-minecraft/", c)

}
