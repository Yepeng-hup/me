package db

import (
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

	Theme struct {
		ThemeName string
		ThemeNum int64
	}
)

func SelectSetUser(){
	database,err := ConnDb()
	var user User
	if err != nil {
		logs.Errorf(err.Error())
	}
	sql := "SELECT * FROM user  WHERE userName = ?"
	rows, err := database.Query(sql, "me")
	if err != nil {
		logs.Errorf("query user table fail,"+err.Error())
	}
	defer rows.Close()
	for rows.Next(){
		err := rows.Scan(&user.Id, &user.UserName, &user.Passwd)
		if err != nil {
			logs.Errorf(err.Error())
		}
	}

	if user.UserName != "me" {
		err := SetMeUserAndPasswd()
		if err != nil{
			logs.Errorf(err.Error())
		}
	}
	return
}


func SelectUserAndPasswd()User{
	database,err := ConnDb()
	var user User
	if err != nil {
		logs.Errorf(err.Error())
	}
	sql := "SELECT * FROM user  WHERE userName = ?"
	rows, err := database.Query(sql, "me")
	if err != nil {
		logs.Errorf("query user table fail,"+err.Error())
	}
	defer rows.Close()
	for rows.Next(){
		err := rows.Scan(&user.Id, &user.UserName, &user.Passwd)
		if err != nil {
			logs.Errorf(err.Error())
		}
	}
	return user
}


func SelectTextRecord()[]Record{
	database,err := ConnDb()
	var record Record
	recordList := make([]Record, 0)
	if err != nil {
		logs.Errorf(err.Error())
	}
	sql := "SELECT * FROM p_text_record"
	rows, err := database.Query(sql, "")
	if err != nil {
		logs.Errorf("query p_text_record table fail,"+err.Error())
	}
	defer rows.Close()
	for rows.Next(){
		err := rows.Scan(&record.Date, &record.Url)
		if err != nil {
			logs.Errorf(err.Error())
		}
		recordList = append(recordList, record)
	}
	return recordList
}


func SelectTextTheme()map[string]int64{
	grapMap := make(map[string]int64, 0)
	database,err := ConnDb()
	var theme Theme
	if err != nil {
		logs.Errorf(err.Error())
	}
	sql := "SELECT themeName, COUNT(*) AS themeName_count FROM text_grap_data GROUP BY themeName;"
	rows, err := database.Query(sql, "")
	if err != nil {
		logs.Errorf("query text_grap_data table fail,"+err.Error())
	}
	defer rows.Close()
	for rows.Next(){
		err := rows.Scan(&theme.ThemeName, &theme.ThemeNum)
		if err != nil {
			logs.Errorf(err.Error())
		}
		grapMap[theme.ThemeName] = theme.ThemeNum
	}
	return grapMap
}
