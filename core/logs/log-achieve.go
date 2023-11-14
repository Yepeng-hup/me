package logs

import (
	"fmt"
	"log"
	"me/global"
	"os"
	"time"
)

type MeLogs interface {
	Info() bool
	Error() bool
	Warn() bool
}

type MeLog struct {
	DateTime string
	Loglevel string
	Describe string
}

var t = time.Now()

func (b *MeLog) Info() bool {
	file, err := os.OpenFile(global.Cfg.Me.LogInfo, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		log.Println("open file error: ",err.Error())
		return false
	}
	defer file.Close()
	writeText := b.DateTime+" "+b.Loglevel+" "+b.Describe+"\n"
	if _, err := file.WriteString(writeText); err != nil {
		log.Println("write file error: ",err.Error())
		return false
	}
	return true
}

func (b *MeLog) Error() bool {
	file, err := os.OpenFile(global.Cfg.Me.LogError, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		log.Println("open file error: ",err.Error())
		return false
	}
	defer file.Close()
	writeText := b.DateTime+" "+b.Loglevel+" "+b.Describe+"\n"
	if _, err := file.WriteString(writeText); err != nil {
		log.Println("write file error: ",err.Error())
		return false
	}
	return true
}


func (b *MeLog) Warn() bool {
	file, err := os.OpenFile(global.Cfg.Me.LogWarn, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		log.Println("open file error: ",err.Error())
		return false
	}
	defer file.Close()
	writeText := b.DateTime+" "+b.Loglevel+" "+b.Describe+"\n"
	if _, err := file.WriteString(writeText); err != nil {
		log.Println("write file error: ",err.Error())
		return false
	}
	return true
}

func Infof(log string) error {
	logs := MeLog{
		DateTime: t.Format("2006-01-02 15:04:05"),
		Loglevel: "INFO",
		Describe: log,
	}
	b := logs.Info()
	if b == false{
		return fmt.Errorf("level [INFO] log write fail")
	}
	return nil
}

func Errorf(log string) error {
	logs := MeLog{
		DateTime: t.Format("2006-01-02 15:04:05"),
		Loglevel: "ERROR",
		Describe: log,
	}
	b := logs.Error()
	if b == false{
		return fmt.Errorf("level [ERROR] log write fail")
	}
	return nil
}

func Warnf(log string) error {
	logs := MeLog{
		DateTime: t.Format("2006-01-02 15:04:05"),
		Loglevel: "WARN",
		Describe: log,
	}
	b := logs.Warn()
	if b == false{
		return fmt.Errorf("level [WARN] log write fail")
	}
	return nil
}

