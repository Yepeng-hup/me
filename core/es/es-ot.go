package es

import (
	"me/core/logs"
	"me/global"
	"net/http"
	"time"
)

func EsConnCheck()bool{
	//es, err := EsConn()
	//if err != nil {
	//	return false
	//}

	client := &http.Client{
		Timeout: time.Second * 3,
	}
	// 创建一个GET请求
	req, err := http.NewRequest("GET", global.Cfg.Elasticsearch.Ssl+"://"+global.Cfg.Elasticsearch.Ip+":"+global.Cfg.Elasticsearch.Port, nil)
	if err != nil {
		logs.Errorf("create http request fail.")
		return false
	}
	resp, err := client.Do(req)
	if err != nil {
		logs.Errorf("check es request fail.")
		return false
	}
	defer resp.Body.Close()

	// 读取响应内容
	//_, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	return false
	//}
	return true

}
