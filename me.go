package main

import (
	"log"
	"me/route"
)

func main(){
	r := route.InitRoute()
	if err := r.Run("127.0.0.1:7070"); err != nil {
		log.Println("ERROR: error start fail", err)
	}
}