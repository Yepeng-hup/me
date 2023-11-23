package text

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"io"
	"log"
	"me/core/logs"
	"me/core/mongo"
	"me/global"
	"os"
)

func fileToBytes(file *os.File) []byte {
	buf := make([]byte, 51200) // 可以根据实际需求调整缓冲区大小,默认5M
	var totalBytes int64 = 0
	for {
		n, err := file.Read(buf)
		// 读到文件末尾，退出循环
		if err == io.EOF {
			break
		} else if err != nil {
			log.Println(err.Error())
			return nil
		} else {
			// n是实际读取到的字节数，totalBytes表示已读的总字节数。如果n为0，表示已经读完整个文件。
			totalBytes += int64(n)
		}
	}
	return buf
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
	if fileToBytes(file) == nil {
		return fmt.Errorf("read file -> [%s] error, file is nil.", saveTitle)
	}

	document := bson.M{
		"filename": saveTitle,
		"content":  fileToBytes(file),
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
