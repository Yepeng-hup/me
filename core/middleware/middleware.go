package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"me/core/logs"
	"net/http"
)

func LoginCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, err := c.Cookie("user"); err != nil {
			err := logs.Errorf(err.Error())
			if err != nil {
				log.Println(err)
			}
			c.Redirect(http.StatusMovedPermanently, "/user/login")
			c.Abort()
			return
		}
	}
}
