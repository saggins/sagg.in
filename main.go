package main

import (
	"net/http"
	"log"
	"github.com/gin-gonic/gin"
	db "github.com/win32prog/sagg.in/web/app/routes"
	"github.com/zalando/gin-oauth2"
	"github.com/zalando/gin-oauth2/zalando"
)

var router *gin.Engine

func redirect(w http.ResponseWriter, req *http.Request) {
    // remove/add not default ports from req.Host
    target := "https://" + req.Host + req.URL.Path 
    if len(req.URL.RawQuery) > 0 {
        target += "?" + req.URL.RawQuery
    }
    log.Printf("redirect to: %s", target)
    http.Redirect(w, req, target,
            // see comments below and consider the codes 308, 302, or 301
            http.StatusTemporaryRedirect)
}

func main() {
	router := gin.Default()
	router.Use(ginglog.Logger(3 * time.Second))
	router.Use(ginoauth2.RequestLogger([]string{"uid"}, "data"))
	router.Use(gin.Recovery())


	router.LoadHTMLGlob("./web/templates/*")
	db.InitializeRoutes(router)
	go http.ListenAndServe(":80", http.HandlerFunc(redirect))
	//router.Run("172.31.18.164:40000")
	http.ListenAndServeTLS(":443", "/etc/letsencrypt/live/sagg.in/fullchain.pem", "/etc/letsencrypt/live/sagg.in/privkey.pem", router)


	
}

