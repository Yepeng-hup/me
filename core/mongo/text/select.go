package text

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"me/core/mongo"
	"me/global"
)


type Mtext struct {
	FileName string `bson:"filename"`
	Content  string    `bson:"content"`
	Size int64 `bson:"size"`
}

type TitleName struct {
	Name string `bson:"filename"`
}



func TextRead(saveTitle string)(string,string,error){
	mg := mongo.MongoConn()
	defer mg.Disconnect(context.TODO())

	database := mg.Database(global.Cfg.Mongodb.DbName)
	ct := database.Collection("text")

	// 构建查询条件
	filter := bson.M{"filename": saveTitle}

	// 执行查询
	var mText Mtext
	err := ct.FindOne(context.TODO(), filter).Decode(&mText)
	if err != nil {
		return "","",fmt.Errorf(err.Error())
	}
	fmt.Println(mText.Content)
	return mText.FileName, mText.Content, nil
}


func ShowTextAllTitle()([]TitleName, error){
	mg := mongo.MongoConn()
	defer mg.Disconnect(context.TODO())
	database := mg.Database(global.Cfg.Mongodb.DbName)
	ct := database.Collection("text")

	var fileKeyList []TitleName
	// 定义过滤条件，这里为空，表示查询所有文档
	filter := bson.D{{}}

	cursor, err := ct.Find(context.TODO(), filter)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	// 遍历查询结果
	for cursor.Next(context.TODO()) {
		var result TitleName
		err := cursor.Decode(&result)
		if err != nil {
			return nil, fmt.Errorf(err.Error())
		}
		fileKeyList = append(fileKeyList, result)
	}

	return fileKeyList, nil
}
