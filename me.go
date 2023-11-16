package main

import (
	"log"
	"me/core"
	"me/core/db"
	"me/global"
	"me/route"
)

var (
	jsonPath = "me.json"
)

func main(){
	core.InitJsonFile(jsonPath)
	r := route.InitRoute()
	core.CreateMeAllTable()
	db.SelectSetUser()
	if err := r.Run(global.Cfg.Me.Ip+":"+global.Cfg.Me.Port); err != nil {
		log.Println("ERROR: error start fail", err)
	}
}