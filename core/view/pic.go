package view

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ShowPic(c *gin.Context){
	c.HTML(http.StatusOK, "pic.tmpl", gin.H{})
}
