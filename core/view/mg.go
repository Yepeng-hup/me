package view

import (
	"github.com/gin-gonic/gin"
	"me/core/mongo"
	"net/http"
)

func MgConnChecks(c *gin.Context){
	b := mongo.MongoConnCheck()
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