package route

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRoute() *gin.Engine {
	//gin.SetMode("release")
	r := gin.Default()
	r.Static("/sta", "static")
	r.LoadHTMLGlob("templates/*/*.tmpl")
	
	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": 200,
		})
	})

	return r
}
