package view

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ShowHome(c *gin.Context){
	c.HTML(http.StatusOK, "index.tmpl", gin.H{

	})
}
