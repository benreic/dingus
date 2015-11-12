package utils

import (
	"log"
	"os"
	"path/filepath"
	"time"
)

var globalLog *log.Logger

/**
 * Sets up the global logger for the app
 *
 * @author Ben Reichelt <ben.reichelt@gmail.com>
 *
 * @return  Logger
**/

func NewLogger() *log.Logger {

	if globalLog != nil {
		return globalLog
	}

	t := time.Now()
	format := "20060102"

	logDir := "/var/log/dingus"
	err := os.MkdirAll(logDir, 0755)
	if err != nil {
		panic(err)
	}

	filePart := t.Format(format)
	filePath := filepath.Join(logDir, filePart+".log")
	var fi *os.File
	if !fileExists(filePath) {
		fi, err = os.Create(filePath)
	} else {
		fi, err = os.OpenFile(filePath, os.O_RDWR|os.O_APPEND, 0755)
	}

	if err != nil {
		panic(err)
	}

	globalLog = log.New(fi, "", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
	return globalLog
}

func LogError(err error) {
	logger := NewLogger()
	logger.Println(err.Error())
}

func Log(message string) {
	logger := NewLogger()
	logger.Println(message)
}
