package view

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ShowText(c *gin.Context){
	c.HTML(http.StatusOK, "text.tmpl", gin.H{})
}
