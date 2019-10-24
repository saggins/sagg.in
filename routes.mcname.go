package main

import (
	"github.com/gin-gonic/gin"
)

type postmsg struct {
	MCuser string `form:"MCuser"`
}

func MCName(c *gin.Context) {
	var Mcname postmsg
	c.ShouldBind(&Mcname)
	MCwhitelist(Mcname, "/home/saggins/Documents/projects/test-minecraft/")
}
