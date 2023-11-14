package global

import (
	"bufio"
	"log"
	"me/conf"
	"os"
	"encoding/json"
)

var (
	Cfg *conf.Config
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
	Cfg = c
	return
}