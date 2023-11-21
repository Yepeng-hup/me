package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"me/core/logs"
	"me/global"
)


func MongoConn()*mongo.Client{
	mgUrl := "mongodb://"+global.Cfg.Mongodb.DbUser+":"+global.Cfg.Mongodb.DbPaaaword+"@"+global.Cfg.Mongodb.Ip+":"+global.Cfg.Mongodb.Port+"/"+global.Cfg.Mongodb.DbName
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mgUrl))
	if err != nil {
		logs.Errorf(err.Error())
		return nil
	}
	return client
}
