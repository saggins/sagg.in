package mc
import (
	"github.com/gin-gonic/gin"

	utlity "github.com/win32prog/sagg.in/web/app/utlity"

	db 	"github.com/win32prog/sagg.in/web/app/db"
	
	model "github.com/win32prog/sagg.in/web/app/models"
)


func MCShop(c *gin.Context) {
	var listOshops []model.MCShop
	listOshops = db.ShopScan()
	utlity.Render(c, "page.html", gin.H{
			})
	utlity.Render(c, "mcshop.html", gin.H{
		"payload":listOshops,
	})

}

//MCSHOP Handle
func MCShopH(post model.MCSPost, c *gin.Context) {
	msg := model.MCShop{
		Name: post.Name,
		Item: post.Item,
		Price: post.Price,
	}
	
	db.PostShops(msg)

	utlity.Render(c, "error.html", gin.H{
		"msg": "WOrkos! You are added. Go ahead and find it My Minecraft Shanngins -> All Shops",
	})
}

func MCShopD(post model.MCSPost, c *gin.Context)  {
	msg := model.MCShop{
		Name: post.Name,
		Item: post.Item,
		Price: post.Price,
	}
	
	db.DelShops(msg)

	utlity.Render(c, "error.html", gin.H{
		"msg": "WOrkos! It's Delted!",
	})
}