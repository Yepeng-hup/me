package view

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
	"log"
	"me/core/auxiliary"
	"me/core/db"
	"me/core/logs"
	"net/http"
	"strings"
)

type (
	CrawlingFrom struct{
		Url string
		SaveXz string
		SaveDataLocal string
		SaveDataEs string
		P1 string
		P2 string
		SaveRecord string
	}

)

func ShowText(c *gin.Context){
	c.HTML(http.StatusOK, "text.tmpl", gin.H{})
}


func CrawlingText(c *gin.Context){
	f := CrawlingFrom{
		Url: c.PostForm("url"),
		SaveXz: c.PostForm("savenum"),
		SaveDataLocal: c.PostForm("localpath"),
		SaveDataEs: c.PostForm("esindex"),
		P1: c.PostForm("c1"),
		P2: c.PostForm("c2"),
		SaveRecord: c.PostForm("hissave"),
	}

	if f.Url == "" {
		logs.Errorf("The input URL cannot be empty.")
		c.Redirect(http.StatusFound, "/svc/text")
		return
	}

	switch f.SaveXz {
		case "本地存储":
			co := colly.NewCollector()
			co.OnHTML("p", func(e *colly.HTMLElement) {
				//保存到本地文件
				err := auxiliary.AppendWrite(e.Text, f.SaveDataLocal)
				if err != nil {
					logs.Errorf(err.Error())
					return
				}
			})
			err := co.Visit(f.Url)
			if err != nil {
				logs.Errorf(err.Error())
				c.Redirect(http.StatusFound, "/svc/text")
				return
			}

			//是否记录此次操作
			if f.SaveRecord == "是" {
				err := db.SetTextRecord(f.Url)
				if err != nil {
					logs.Errorf(err.Error())
				}
				logs.Infof("Record text successfully.")
			}

			c.Redirect(http.StatusFound, "/svc/text")
			return

		case "elasticsearch":
			fmt.Println("-----", f.Url, f.SaveXz, f.SaveDataEs)

		default:
			err := logs.Warnf("not is storage type.")
			if err != nil{
				log.Println(err)
			}
	}


}


func ShowTextRecord(c *gin.Context){
	record := db.SelectTextRecord()
	c.HTML(http.StatusOK, "textrecord.tmpl", gin.H{
		"UrlRecord": record,
	})
}


func DeleteTextRecord(c *gin.Context){
	url := c.PostForm("url")
	urlList := strings.Fields(url)
	err := db.DelTextRecord(urlList[2])
	if err != nil {
		l := fmt.Sprintf("delete text record fail -> [%s].", urlList[2])
		logs.Errorf(l)
		return
	}
	c.Redirect(http.StatusFound, "/svc/text/record")
	return
}