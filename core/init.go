package core

import (
	"log"
	"me/core/db"
)


func CreateMeAllTable(){
	database, err := db.ConnDb()
	if err != nil {
		log.Fatal(err.Error())
	}
	createUserTable := `CREATE TABLE IF NOT EXISTS user (id INTEGER PRIMARY KEY, userName TEXT, password TEXT);`
	_, err = database.Exec(createUserTable)
	if err != nil {
		log.Fatal("ERROR: create table cron fail,", err.Error())
	}
	err = db.SetMeUserAndPasswd()
	if err != nil {
		log.Fatal(err.Error())
	}
	//第二张表
	// ......
	return
}
