package route

import (
	"github.com/gin-gonic/gin"
	"me/core/middleware"
	"me/core/view"
	"net/http"
)

func InitRoute() *gin.Engine {
	//gin.SetMode("release")
	r := gin.Default()
	r.Static("/sta", "static")
	r.LoadHTMLGlob("templates/*/*.tmpl")

	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/user/login")
	})
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"msg": "Not Route",
		})
	})

	user := r.Group("/user")
		user.GET("/login", func(c *gin.Context) {
			c.HTML(http.StatusOK, "login.tmpl", gin.H{})
		})
		user.POST("/login",view.Login)
		user.GET("/logout", view.Logout)

	svc := r.Group("/svc")
		svc.POST("/user/update",middleware.LoginCheck(),view.PwdUpdate)
		svc.GET("/index",middleware.LoginCheck(), view.ShowHome)
		svc.GET("/text", middleware.LoginCheck(), view.ShowText)
		svc.POST("/text/crawling", middleware.LoginCheck(), view.CrawlingText)
		svc.GET("/text/record", middleware.LoginCheck(), view.ShowTextRecord)
		svc.POST("/text/record/del", middleware.LoginCheck(), view.DeleteTextRecord)
		svc.GET("/text/es/ck",middleware.LoginCheck(), view.EsConnChecks)
		svc.GET("/text/mg/ck",middleware.LoginCheck(), view.MgConnChecks)
		svc.GET("/text/mg/content/:titleName",middleware.LoginCheck(), view.ShowTextCat)
		svc.GET("/text/mg/content/list", middleware.LoginCheck(),view.ShowAllText)

		svc.GET("/video", middleware.LoginCheck(), view.ShowVideo)

		svc.GET("/pic", middleware.LoginCheck(), view.ShowPic)

	return r
}
