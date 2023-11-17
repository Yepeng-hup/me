package db

import (
	"fmt"
	"log"
)


func UpdateMePasswd(passwd string)error{
	database,err := ConnDb()
	if err != nil {
		log.Println(err.Error())
	}

	stmt, err := database.Prepare("UPDATE user SET password = ? WHERE userName = ?")
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	defer stmt.Close()

	// 执行更新操作
	_, err = stmt.Exec(passwd, "me")
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	return nil
}
