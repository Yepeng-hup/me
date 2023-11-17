package auxiliary

import (
	"time"
)

func GeneratePrefix()string{
	t := time.Now()
	t.Format("2006-01-02 15:04:05")
	return ""
}
