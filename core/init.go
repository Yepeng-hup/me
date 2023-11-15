package core

import (
	"fmt"
	"log"
	"me/core/auxiliary"
	"me/core/db"
	"me/core/logs"
)

var tableList = []string{"user",}

func checkTableIfCreate()[]string{
	t := make([]string , 0)
	database, err := db.ConnDb()
	if err != nil {
		log.Fatal(err.Error())
	}
	// 查询所有表的名称
	rows, err := database.Query("SELECT name FROM sqlite_master WHERE type='table';")
	if err != nil {
		log.Fatal(err.Error())
		return nil
	}
	defer rows.Close()

	for rows.Next() {
		var tableName string
		err := rows.Scan(&tableName)
		if err != nil {
			log.Fatal(err.Error())
			return nil
		}
		t = append(t, tableName)
	}
	return t
}


func CreateMeAllTable(){
	database, err := db.ConnDb()
	if err != nil {
		log.Fatal(err.Error())
	}

	list := checkTableIfCreate()
	for _, v := range tableList {
		if auxiliary.IfElement(list, v) {
			l := fmt.Sprintf("table -> [%s] Already exists.", v)
			logs.Warnf(l)
		}else {
			switch v {
			case "user":
				createUserTable := `CREATE TABLE IF NOT EXISTS user (id INTEGER PRIMARY KEY, userName TEXT, password TEXT);`
				_, err = database.Exec(createUserTable)
				if err != nil {
					log.Fatal("ERROR: create table cron fail,", err.Error())
				}
				err = db.SetMeUserAndPasswd()
				if err != nil {
					log.Fatal(err.Error())
				}
			default:
				return
			}
		}
	}
	return
}
