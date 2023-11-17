package view

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"me/core/logs"
	"net/http"
	"me/core/db"
	"time"
)

type (
	loginFrom struct {
		userName string
		password string
	}
	pwd struct {
		password01 string
		password02 string
	}
)


func Login(c *gin.Context){
	f := loginFrom{
		userName: c.PostForm("username"),
		password: c.PostForm("password"),
	}
	user := db.SelectUserAndPasswd()
	if f.userName == user.UserName && f.password == user.Passwd {
		l := fmt.Sprintf("user -> [%s] login success.", f.userName)
		err := logs.Infof(l)
		if err != nil{
			log.Println(err)
			return
		}
		//设置cookie
		meCookie := http.Cookie{
			Name:    "user",
			Value:   f.userName,
			Expires: time.Now().Add(24 * time.Hour), // Cookie的过期时间，这里设置为24小时后过期
			Path:    "/",
		}
		http.SetCookie(c.Writer, &meCookie)
		c.Redirect(http.StatusFound, "/svc/index")
		return
	}

	l := fmt.Sprintf("user -> [%s] login fail.", f.userName)
	err := logs.Errorf(l)
	if err != nil {
		log.Println(err)
	}
	c.HTML(http.StatusInternalServerError, "login.tmpl", gin.H{
		"error": "user or passwd error!",
	})

}


func Logout(c *gin.Context){
	meCookie := http.Cookie{
		Name:    "user",
		Expires: time.Now().Add(-time.Hour), // 设置为过去的时间
		Path:    "/",
	}
	http.SetCookie(c.Writer, &meCookie)
	c.Redirect(http.StatusFound, "/user/login")
	return
}


func PwdUpdate(c *gin.Context){
	p := pwd{
		password01: c.PostForm("ypwd"),
		password02: c.PostForm("epwd"),
	}

	if p.password01 == "" || p.password02 == "" {
		logs.Errorf("password input nill.")
		c.Redirect(http.StatusFound, "/svc/index")
		return
	}

	if p.password01 != p.password02 {
		 logs.Errorf("password Inconsistent input,password update fail.")
		c.Redirect(http.StatusFound, "/svc/index")
		 return
	}

	err := db.UpdateMePasswd(p.password02)
	if err != nil {
		logs.Errorf(err.Error())
		c.Redirect(http.StatusFound, "/svc/index")
		return
	}
	logs.Infof("password update success.")
	c.Redirect(http.StatusFound, "/user/logout")
	return
}