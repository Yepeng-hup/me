package text

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"io/ioutil"
	"log"
	"me/core/logs"
	"me/core/mongo"
	"me/global"
	"os"
)

func fileToStr(filePath string) string {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Println(err.Error())
	}
	// 字节切片转换为字符串
	return string(content)
}

func TextWrite(filePath, saveTitle string) error {
	mg := mongo.MongoConn()
	defer mg.Disconnect(context.TODO())

	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	defer file.Close()
	// 获取文件的内容长度
	info, err := file.Stat()
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	length := info.Size()

	// new mongo obj
	collection := mg.Database(global.Cfg.Mongodb.DbName).Collection("text")
	if fileToStr(filePath) == "" {
		return fmt.Errorf("read file -> [%s] error, file is nil.", saveTitle)
	}

	document := bson.M{
		"filename": saveTitle,
		"content":  fileToStr(filePath),
		"size":     length,
	}
	_, err = collection.InsertOne(context.TODO(), document)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	l := fmt.Sprintf("filePath  -> %s save mongo success.", filePath)
	logs.Infof(l)
	return nil
}