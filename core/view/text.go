package view

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
	"log"
	"me/core/auxiliary"
	"me/core/db"
	"me/core/logs"
	"me/core/mongo/text"
	"net/http"
	"strings"
)

type (
	CrawlingFrom struct{
		Url string
		SaveXz string
		SaveDataLocal string
		SaveDataMGTitle string
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
		SaveDataMGTitle: c.PostForm("mgtl"),
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
				logs.Infof("Record text success.")
			}

			c.Redirect(http.StatusFound, "/svc/text")
			return

		case "MongoDB":
			co := colly.NewCollector()
			co.OnHTML("p", func(e *colly.HTMLElement) {
				//临时把内容保存文件
				err := auxiliary.AppendWrite(e.Text, f.SaveDataMGTitle+".txt")
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

			//save mongodb
			if text.TextWrite(f.SaveDataMGTitle+".txt", f.SaveDataMGTitle) != nil {
				logs.Errorf(err.Error())
				return
			}
			//del temporary file
			if auxiliary.DeleteF(f.SaveDataMGTitle+".txt") != nil {
				logs.Errorf(err.Error())
				return
			}

			if f.SaveRecord == "是" {
				err := db.SetTextRecord(f.Url)
				if err != nil {
					logs.Errorf(err.Error())
				}
				logs.Infof("Record text success.")
			}
			c.Redirect(http.StatusFound, "/svc/text")
			return

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