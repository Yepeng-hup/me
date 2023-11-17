package db

import (
	lo "log"
	"me/core/logs"
)

type (
	User struct {
		Id string
		UserName string
		Passwd string
	}

	Record struct {
		Date string
		Url string
	}
)

func SelectSetUser(){
	database,err := ConnDb()
	var user User
	if err != nil {
		lo.Fatal(err.Error())
	}
	sql := "SELECT * FROM user  WHERE userName = ?"
	rows, err := database.Query(sql, "me")
	if err != nil {
		lo.Fatal("query user table fail,",err.Error())
	}
	defer rows.Close()
	for rows.Next(){
		err := rows.Scan(&user.Id, &user.UserName, &user.Passwd)
		if err != nil {
			lo.Fatal(err.Error())
		}
	}

	if user.UserName != "me" {
		err := SetMeUserAndPasswd()
		if err != nil{
			err := logs.Errorf(err.Error())
			if err != nil{
				lo.Fatal(err.Error())
			}

		}
	}
	return
}


func SelectUserAndPasswd()User{
	database,err := ConnDb()
	var user User
	if err != nil {
		lo.Fatal(err.Error())
	}
	sql := "SELECT * FROM user  WHERE userName = ?"
	rows, err := database.Query(sql, "me")
	if err != nil {
		lo.Fatal("query user table fail,",err.Error())
	}
	defer rows.Close()
	for rows.Next(){
		err := rows.Scan(&user.Id, &user.UserName, &user.Passwd)
		if err != nil {
			lo.Fatal(err.Error())
		}
	}
	return user
}


func SelectTextRecord()[]Record{
	database,err := ConnDb()
	var record Record
	recordList := make([]Record, 0)
	if err != nil {
		lo.Fatal(err.Error())
	}
	sql := "SELECT * FROM p_text_record"
	rows, err := database.Query(sql, "")
	if err != nil {
		lo.Fatal("query p_text_record table fail,",err.Error())
	}
	defer rows.Close()
	for rows.Next(){
		err := rows.Scan(&record.Date, &record.Url)
		if err != nil {
			lo.Fatal(err.Error())
		}
		recordList = append(recordList, record)
	}
	return recordList
}
