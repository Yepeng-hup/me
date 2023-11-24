package core

import (
	"fmt"
	"log"
	"me/core/auxiliary"
	"me/core/db"
	"me/core/logs"
)

var tableList = []string{"user", "p_text_record", "text_grap_data"}

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
			l := fmt.Sprintf("server start init table -> [%s] Already exists.", v)
			logs.Warnf(l)
		}else {
			switch v {
			case "user":
				createUserTable := `CREATE TABLE IF NOT EXISTS user (id INTEGER PRIMARY KEY, userName TEXT, password TEXT);`
				_, err = database.Exec(createUserTable)
				if err != nil {
					log.Fatal("ERROR: create table user fail,", err.Error())
				}
				err = db.SetMeUserAndPasswd()
				if err != nil {
					log.Fatal(err.Error())
				}
				logs.Infof("table -> "+v+" create success.")
			case "p_text_record":
				createTextRrcordTable := `CREATE TABLE IF NOT EXISTS p_text_record (use_date TEXT DEFAULT (strftime('%Y-%m-%d %H:%M', 'now', 'localtime')), urlName TEXT);`
				_, err := database.Exec(createTextRrcordTable)
				if err != nil {
					log.Fatal("ERROR: create table p_text_record fail,", err.Error())
				}
				logs.Infof("table -> "+v+" create success.")
			case "text_grap_data":
				createTextGrapTable := `CREATE TABLE IF NOT EXISTS text_grap_data (use_date TEXT DEFAULT (strftime('%Y-%m-%d %H:%M', 'now', 'localtime')), themeName TEXT);`
				_, err := database.Exec(createTextGrapTable)
				if err != nil {
					log.Fatal("ERROR: create table text_grap_data fail,", err.Error())
				}
				logs.Infof("table -> "+v+" create success.")

			default:
				err := logs.Errorf("not is table -> "+v)
				if err != nil {
					log.Println(err)
					return
				}
				return
			}
		}
	}
	return
}
