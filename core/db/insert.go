package db

import (
	"fmt"
)

func SetMeUserAndPasswd()error{
	database, err := ConnDb()
	sql := `INSERT INTO user (userName, password) VALUES (?, ?);`
	_, err = database.Exec(sql, "me", "me")
	if err != nil {
		return fmt.Errorf("insert data to table[user] fail, %s", err.Error())
	}
	return nil
}


func SetTextRecord(url string)error{
	database, err := ConnDb()
	sql := `INSERT INTO p_text_record (urlName) VALUES (?);`
	_, err = database.Exec(sql, url)
	if err != nil {
		return fmt.Errorf("insert data to table[p_text_record] fail, %s", err.Error())
	}
	return nil
}


func SetTextTheme(theme string)error{
	database, err := ConnDb()
	sql := `INSERT INTO text_grap_data (themeName) VALUES (?);`
	_, err = database.Exec(sql, theme)
	if err != nil {
		return fmt.Errorf("insert data to table[text_grap_data] fail, %s", err.Error())
	}

	return nil
}