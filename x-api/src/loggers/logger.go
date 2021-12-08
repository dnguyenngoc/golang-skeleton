package loggers

import (
	"log"
	"os"
)

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

func InitLoggerConfig() {
	log.Println("Start init logger .......")
	file, err := os.OpenFile("./loggers/x-api.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Panic(err)
	}
	InfoLogger = log.New(file, "[INFO]: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(file, "[WARNING]: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(file, "[ERROR]: ", log.Ldate|log.Ltime|log.Lshortfile)
	log.Println("Complete setup logger!")
}
