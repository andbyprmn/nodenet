package log

import (
	"log"
	"os"
)

var (
	logger  *log.Logger
	logFile *os.File
)

func InitLogger(logFilePath string) error {
	var err error
	logFile, err = os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	logger = log.New(logFile, "NodeNet: ", log.Ldate|log.Ltime|log.Lshortfile)
	return nil
}

func LogCommand(command string) {
	logger.Printf("Command received: %s", command)
}

func LogError(err error) {
	logger.Printf("Error: %v", err)
}

func CloseLogger() {
	logFile.Close()
}
