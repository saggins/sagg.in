package routes

import (	
	"github.com/gin-gonic/gin"
	
	mc "github.com/win32prog/sagg.in/scripts/MC"

	model "github.com/win32prog/sagg.in/web/app/models"
)



func MCShop(c *gin.Context) {
	var MCShopPost model.MCSPost
	c.ShouldBind(&MCShopPost)

	mc.MCShopH(MCShopPost, c)
}
func DMShop(c *gin.Context)  {
	var MCShopPost model.MCSPost
	c.ShouldBind(&MCShopPost)

	mc.MCShopD(MCShopPost, c)
}
