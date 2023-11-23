package text

import (
	"context"
	"fmt"
	"me/core/mongo"
	"me/global"
)

func TextDeleteContent(fileTitle string) error {
	mg := mongo.MongoConn()
	defer mg.Disconnect(context.TODO())

	database := mg.Database(global.Cfg.Mongodb.DbName)
	ct := database.Collection("text")

	filter := map[string]interface{}{
		"filename": fileTitle,
	}

	//use del
	rel, err := ct.DeleteOne(context.Background(), filter)
	if err != nil {
		return fmt.Errorf("delete mongo db me/text -> [%s] fail.", fileTitle)
	}
	fmt.Println(rel.DeletedCount)
	return nil
}
