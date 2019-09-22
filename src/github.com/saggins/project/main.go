package main

import (
	"net/http"

	gintemplate "github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)

func main() {
	router := gin.Default()

	//new template engine
	router.HTMLRender = gintemplate.New(gintemplate.TemplateConfig{
		Root:         "views",
		Extension:    ".tpl",
		Master:       "template/master",
		DisableCache: true,
	})

	router.GET("/", func(ctx *gin.Context) {

		ctx.HTML(http.SatusOK, "index", gin.H)
	})

	router.Run(":8080") // Listen and Server in 0.0.0.0:8080
}
