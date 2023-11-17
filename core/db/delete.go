package db

import (
	"fmt"
	"log"
	"me/core/logs"
)

func DelTextRecord(url string)error{
	database,err := ConnDb()
	if err != nil {
		log.Println(err.Error())
	}

	sql := "DELETE FROM p_text_record WHERE urlName = ?"
	stmt, err := database.Prepare(sql)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(url)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	l := fmt.Sprintf("delete text record ok. name -> [%v].", url)
	logs.Infof(l)
	return nil
}
