package routes

import (
	"html/template"
	"net/http"
	"fmt"
	"io/ioutil"
	"log"
	"encoding/json"
	"os"


	"golang.org/x/oauth2"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/contrib/sessions"
	"golang.org/x/oauth2/google"


	discord "github.com/win32prog/sagg.in/scripts/Discord"
	mc "github.com/win32prog/sagg.in/scripts/MC"
	db "github.com/win32prog/sagg.in/web/app/db"
	model "github.com/win32prog/sagg.in/web/app/models"
	utlity "github.com/win32prog/sagg.in/web/app/utlity"
	auth "github.com/win32prog/sagg.in/web/app/auth"
)

var conf = oauth2.Config{}

func init() {
	var c model.OauthCreds
    file, err := ioutil.ReadFile("/home/ubuntu/saggweb/src/github.com/win32prog/sagg.in/oauth.json")
    if err != nil {
        fmt.Printf("File error: %v\n", err)
        os.Exit(1)
    }
	json.Unmarshal(file, &c)
	
	conf = oauth2.Config{
		ClientID:     c.Cid,
		ClientSecret: c.Csecret,
		RedirectURL:  "https://sagg.in/login/auth",
		Scopes:       []string{"profile", "email", "openid"},
		Endpoint: google.Endpoint,
	}
}

func InitializeRoutes(r *gin.Engine) {
	// INDEX is special... Its a special boi
	r.GET("/", showIndexPage)

	pageroutes := r.Group("/page")
	{
		pageroutes.GET("/view/:page_id", getPage)

		pageroutes.POST("/view/mcname/", MCName)
		pageroutes.POST("/view/mcshop", MCShop)
	}
	mcroutes := r.Group("/mc")
	{
		mcroutes.GET("view/players", mc.MCPlayers)
		mcroutes.GET("view/shops", mc.MCShop)
		mcroutes.POST("view/shop/del", DMShop)
	}
	discordroutes := r.Group("/discord")
	{
		discordroutes.GET("/view/invite/:discord_id/:discord_avatar", discord.DSPlayers)
		discordroutes.POST("/view/invite/dsname", discord.DSResult)
		discordroutes.GET("view/invitelist", discord.Listplayers)
	}
	user := r.Group("/login")
	{
		user.GET("/", loginHandler)
		user.GET("/auth", authHandler)//TODO auth Handler
	}

}
// Login/Oauth

func loginHandler(c *gin.Context) {
    state := auth.Randomtoken()
    session := sessions.Default(c)
    session.Set("state", state)
    session.Save()
	utlity.Render(c, "login.html", gin.H{
		"State": auth.GetLoginURL(conf, state),
	})
}

func authHandler(c *gin.Context) {
	// Check state validity.
	
    session := sessions.Default(c)
    retrievedState := session.Get("state")
    if retrievedState != c.Query("state") {
        c.AbortWithError(http.StatusUnauthorized, fmt.Errorf("Invalid session state: %s", retrievedState))
        return
    }
    // Handle the exchange code to initiate a transport.
  	tok, err := conf.Exchange(oauth2.NoContext, c.Query("code"))
  	if err != nil {
  		c.AbortWithError(http.StatusBadRequest, err)
          return
  	}
    // Construct the client.
    client := conf.Client(oauth2.NoContext, tok)
    resp, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	log.Println("Resp body: ", string(data))
	if () 
	utlity.Render(c, "page.html", gin.H{
		"title": ,
		"payload": ,
	})
	
}

// Special Index page
func showIndexPage(c *gin.Context) {

	//just fo nav... we love nav to be loopy

	utlity.Render(c, "home.html", gin.H{})

}

// Generic database to page handler
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


// Scan names database and add to whitelist
func MCName(c *gin.Context) {
	var Mcname model.Postmsg
	c.ShouldBind(&Mcname)
	mc.MCwhitelist(Mcname, "/home/saggins/Documents/projects/test-minecraft/", c)

}

// Scan all apps and build html
func MCShop(c *gin.Context) {
	var MCShopPost model.MCSPost
	c.ShouldBind(&MCShopPost)

	mc.MCShopH(MCShopPost, c)
}

// Delete a shop
func DMShop(c *gin.Context) {
	var MCShopPost model.MCSPost
	c.ShouldBind(&MCShopPost)

	mc.MCShopD(MCShopPost, c)
}
