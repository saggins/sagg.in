package main

import (
//	"net/http"
	"github.com/gin-gonic/gin"
	db "github.com/win32prog/sagg.in/web/app/routes"
)

var router *gin.Engine

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("./web/templates/*")
	db.InitializeRoutes(router)

	router.Run("172.31.18.164:40000")
	//http.ListenAndServeTLS(":443", "/etc/letsencrypt/live/sagg.in/fullchain.pem", "/etc/letsencrypt/live/sagg.in/privkey.pem", router)

}

