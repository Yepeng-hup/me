package view

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ShowVideo(c *gin.Context){
	c.HTML(http.StatusOK, "video.tmpl", gin.H{})
}
