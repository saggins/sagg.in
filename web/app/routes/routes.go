package routes

import (
	"github.com/gin-gonic/gin"
	
	mc "github.com/win32prog/sagg.in/scripts/MC"
)


func InitializeRoutes(r *gin.Engine) {
	// INDEX is special... Its a special boi
	r.GET("/", showIndexPage)
 
	pageroutes := r.Group("/page")
	{
		pageroutes.GET("/view/:page_id", getPage)

		pageroutes.POST("/view/mcname/", MCName)
		pageroutes.POST("/view/mcshop", MCShop )
	}
	mcroutes := r.Group("/mc")
	{
		mcroutes.GET("view/players", mc.MCPlayers)
		mcroutes.GET("view/shops", mc.MCShop)
		mcroutes.POST("view/shop/del", DMShop)
	}


}
