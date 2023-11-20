package es

import (
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"me/core/logs"
	"me/global"
)

func EsConn() (*elasticsearch.Client, error) {
	// new es clien config
	NewEsConf := elasticsearch.Config{
		Addresses: []string{global.Cfg.Elasticsearch.Ssl+"://"+global.Cfg.Elasticsearch.Ip+":"+global.Cfg.Elasticsearch.Port},
	}

	// new es client
	es, err := elasticsearch.NewClient(NewEsConf)
	if err != nil {
		logs.Errorf(err.Error())
		return nil, fmt.Errorf(err.Error())
	}
	return es, nil
}