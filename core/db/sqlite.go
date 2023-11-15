package db

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func ConnDb()(*sql.DB,error){
	db, err := sql.Open("sqlite3", "me.db")
	if err != nil {
		return nil, fmt.Errorf("EEROR: conn db sqlite fail, %v", err.Error())
	}
	return db, nil
}

