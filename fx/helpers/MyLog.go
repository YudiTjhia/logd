package helpers

import (
	"fmt"
	"log"
	"os"
	"time"
)

const (
	LOG_FILE = "../log/access"
	LOG_ERR = "../log/error"
)

func LogDebug(accountID string, loginID string,
	struc string,
	method string,
	msg ...interface{}) {

	defer func() { r := recover(); if r != nil { fmt.Print("Error detected logging:", r) } }()
	logFile :=  LOG_FILE + "-" + time.Now().Format("20060102") + ".log"
	fp, err := os.OpenFile(logFile , os.O_RDWR | os.O_CREATE | os.O_APPEND,0666)

	if err==nil {

		log.SetOutput(fp)
		log.Println(accountID, ">", loginID, ">", struc, ">", method, ">", msg)
		fmt.Println(accountID, ">", loginID, ">", struc, ">", method, ">", msg)
		_ = fp.Close()

	}
}

func LogErr(accountID string, loginID string, struc string, method string, msg ...interface{}) {

	defer func() { r := recover(); if r != nil { fmt.Print("Error detected logging:", r) } }()

	logFile :=  LOG_ERR + "-" + time.Now().Format("20060102") + ".err"
	fp, err := os.OpenFile(logFile , os.O_RDWR | os.O_CREATE | os.O_APPEND,0666)

	if err==nil {

		log.SetOutput(fp)
		log.Println(accountID, ">", loginID, ">", struc, ">", method, ">ERR=", msg)
		fmt.Println(accountID, ">", loginID, ">", struc, ">", method, ">ERR=", msg)
		_ = fp.Close()

	}

}
