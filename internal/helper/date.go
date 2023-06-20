package helper

import (
	"time"
)

func GetDateNow() string {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	return time.Now().In(loc).Format("2006-01-02 15:04:05")
}

func AddMinutes(m int) string {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	return time.Now().In(loc).Add(time.Minute * time.Duration(m)).Format("2006-01-02 15:04:05")
}

func DateFilename() string {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	return time.Now().In(loc).Format("20060102150405")
}
