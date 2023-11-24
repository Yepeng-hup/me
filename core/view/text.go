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
		HtmlLable string
		Theme string
	}
	TextCat struct {
		TextTitle string
	}

)

func ShowText(c *gin.Context){
	themeMap := db.SelectTextTheme()
	c.HTML(http.StatusOK, "text.tmpl", gin.H{
		"themeMap": themeMap,
	})
}


func CrawlingText(c *gin.Context){
	f := CrawlingFrom{
		Url: c.PostForm("url"),
		SaveXz: c.PostForm("savenum"),
		SaveDataLocal: c.PostForm("localpath"),
		SaveDataMGTitle: c.PostForm("mgtl"),
		SaveRecord: c.PostForm("hissave"),
		HtmlLable: c.PostForm("htmlLable"),
		Theme:c.PostForm("theme"),
	}

	if f.Url == "" {
		logs.Errorf("The input URL cannot be empty.")
		c.Redirect(http.StatusFound, "/svc/text")
		return
	}

	err := db.SetTextTheme(f.Theme)
	if err != nil {
		logs.Errorf(err.Error())
		c.Redirect(http.StatusFound, "/svc/text")
		return
	}

	co := colly.NewCollector()
	strSlice := auxiliary.StrConvertSlice(f.HtmlLable)

	switch f.SaveXz {
		case "本地存储":
			if len(strSlice) == 1 {
				go func(htmlLable string) {
					co.OnHTML(htmlLable, func(e *colly.HTMLElement) {
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

					if f.SaveRecord == "是" {
						err := db.SetTextRecord(f.Url)
						if err != nil {
							logs.Errorf(err.Error())
						}
						logs.Infof("Record text success.")
					}
				}(f.HtmlLable)
			}else {
				for _,v := range strSlice {
					//异步，没有顺序
					go func(htmlLable string) {
						co.OnHTML(htmlLable, func(e *colly.HTMLElement) {
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

						if f.SaveRecord == "是" {
							err := db.SetTextRecord(f.Url)
							if err != nil {
								logs.Errorf(err.Error())
							}
							logs.Infof("Record text success.")
						}
					}(v)
				}
			}

			c.Redirect(http.StatusFound, "/svc/text")
			return

		case "MongoDB":
			if len(strSlice) == 1 {
				go func(htmlLable string) {
					co.OnHTML(htmlLable, func(e *colly.HTMLElement) {
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
						logs.Infof(f.Url+"  Record text success.")
					}
				}(f.HtmlLable)
			}else {
				for _,v := range strSlice {
					go func(htmlLable string) {
						co.OnHTML(htmlLable, func(e *colly.HTMLElement) {
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

						if text.TextWrite(f.SaveDataMGTitle+".txt", f.SaveDataMGTitle) != nil {
							logs.Errorf(err.Error())
							return
						}

						if auxiliary.DeleteF(f.SaveDataMGTitle+".txt") != nil {
							logs.Errorf(err.Error())
							return
						}

						if f.SaveRecord == "是" {
							err := db.SetTextRecord(f.Url)
							if err != nil {
								logs.Errorf(err.Error())
							}
							logs.Infof(f.Url+"  Record text success.")
						}
					}(v)
				}
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


func DeleteTextContent(c *gin.Context){
	t := TextCat{
		TextTitle: c.PostForm("titleName"),
	}
	tlList := strings.Fields(t.TextTitle)
	err := text.TextDeleteContent(tlList[0])
	if err != nil {
		logs.Errorf(err.Error())
		return
	}
	logs.Infof(tlList[0]+" content delete success.")
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
	})
}


func ShowTextCat(c *gin.Context){
	t := TextCat{
		TextTitle: c.Param("titleName"),
	}
	//tList := strings.Fields(t.TextTitle)
	fileName,content,err := text.TextRead(t.TextTitle)
	if err != nil {
		logs.Errorf(err.Error())
		c.Redirect(http.StatusFound, "/svc/text")
		return
	}
	c.HTML(http.StatusOK, "textcat.tmpl", gin.H{
			"fileName": fileName,
			"content": content,
	})
}


func ShowAllText(c *gin.Context){
	k, err := text.ShowTextAllTitle()
	if err != nil {
		logs.Errorf(err.Error())
		return
	}
	c.HTML(http.StatusOK, "texttitlelist.tmpl", gin.H{
		"keyName": k,
	})
}