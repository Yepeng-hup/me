package view

import (
	"github.com/gin-gonic/gin"
	"me/core/es"
	"net/http"
)

func EsConnChecks(c *gin.Context){
	b := es.EsConnCheck()
	if b == true{
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg": "connect success!",
		})
	}else {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg": "connect fail!",
		})
	}
}
