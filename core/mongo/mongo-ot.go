package mongo

import (
	"fmt"
	"me/core/logs"
	"me/global"
	"net"
	"time"
)

func MongoConnCheck()bool{
	address := fmt.Sprintf("%s:%s", global.Cfg.Mongodb.Ip, global.Cfg.Mongodb.Port)
	fmt.Println(address, global.Cfg.Elasticsearch.Ssl)
	// new conn
	conn, err := net.DialTimeout("tcp", address, 3*time.Second)
	if err != nil {
		logs.Errorf("check mongo request fail.")
		return false
	}
	defer conn.Close()

	return true
}
