package view

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
	"log"
	"time"
	"context"
)

func NewPc(c *gin.Context){
	co := colly.NewCollector()
	co.OnRequest(func(r *colly.Request) {
		fmt.Println("url ---> ", r.URL)
	})
	co.OnError(func(_ *colly.Response, err error) {
		log.Println("spider show fail,", err.Error())
	})

	//co.OnHTML("a[href]", func(e *colly.HTMLElement) {
	//	// 提取视频链接
	//	videoURL := e.Attr("href")
	//	fmt.Println("Video URL:", videoURL)
	//})
	co.OnHTML("p", func(e *colly.HTMLElement) {
		fmt.Println("text ---> ", e.Text)
	})
	err := co.Visit("https://www.bilibili.com/read/cv27081966/?from=category_0")
	if err != nil {
		log.Println(err.Error())
		return
	}
}


// 异步函数，模拟耗时任务
func AsyncTask(ctx context.Context, done chan bool) {
	fmt.Println("开始执行异步任务...")
	defer func() {
		done <- true // 通知主handler任务完成
	}()

	select {
	case <-time.After(3 * time.Second): // 模拟任务耗时
		fmt.Println("异步任务执行完成.")
	case <-ctx.Done():
		fmt.Println("异步任务被取消.")
	}
}
