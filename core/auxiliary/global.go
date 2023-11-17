package auxiliary

import (
	"fmt"
	"log"
	"os"
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


