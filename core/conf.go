package core

import (
	"bufio"
	"log"
	"me/conf"
	"me/global"
	"os"
	"encoding/json"
)

func InitJsonFile(path string){
	c := &conf.Config{}
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("open json file: ", err.Error())
	}
	defer file.Close()
	f := bufio.NewReader(file)
	configObj := json.NewDecoder(f)
	if err = configObj.Decode(&c); err != nil {
		log.Fatal(err.Error())
		return
	}
	global.Cfg = c
	return
}
