package main

import (
	"log"
	"me/core"
	"me/global"
	"me/route"
)

var (
	jsonPath = "me.json"
)

func main(){
	global.InitJsonFile(jsonPath)
	r := route.InitRoute()
	err := core.CreateMeAllTable()
	if err != nil{
		log.Fatal(err)
	}
	if err := r.Run(global.Cfg.Me.Ip+":"+global.Cfg.Me.Port); err != nil {
		log.Println("ERROR: error start fail", err)
	}
}