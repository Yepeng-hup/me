package auxiliary

import (
	"fmt"
	"log"
	"os"
	"time"
)

func IfElement(slice []string, element string) bool {
	for _, v := range slice {
		if v == element {
			return true
		}
	}
	return false
}

func AppendWrite(line, filePath string)error{
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		return fmt.Errorf("open file error: %s",err.Error())
	}
	defer file.Close()
	writeText := line+"\n"
	if _, err := file.WriteString(writeText); err != nil {
		log.Println("write file error: ",err.Error())
		return fmt.Errorf("write file error: %s",err.Error())
	}
	return nil
}


func GeneratePrefix()string{
	t := time.Now()
	return t.Format("2006-01-02 15:04:05")
}


func DeleteF(path string)error{
	err := os.Remove(path)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	return nil
}


func DeleteD(path string)error{
	err := os.RemoveAll(path)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	return nil
}


