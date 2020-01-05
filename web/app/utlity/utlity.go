package utlity

import (
	"net/http"

	"github.com/gin-gonic/gin"
	model "github.com/win32prog/sagg.in/web/app/models"
)

func Render(c *gin.Context, templateName string, data gin.H) {
	switch c.Request.Header.Get("Accept") {
	case "application/json":
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		c.XML(http.StatusOK, data["payload"])
	default:
		c.HTML(http.StatusOK, templateName, data)
	}
}

func Profile( state string) model.Profile {
	
}