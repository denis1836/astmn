package log

import (
	"io"
	"fmt"
	"os"
	"path/filepath"

	config "astmn/config"
	ui "astmn/ui"
)

var logger *log.Logger

func InitLogger() (io.Writer, error){
	logDir := config.GetLogDir()
	if err := os.MkdirAll(logDir, 0755); err != nil {
		ui.pError("can't create log dir: " + err.Error())
		return nil, err
	}

	logFileName := time.Now().Format("2006-01-02_15-04") + "-astmn.log"

	fullPath := filepath.Join(logDir, logFileName)
	absPath, err := filepath.Abs(fullPath)
	if err != nil {
		ui.pError("can't get abs path: " + err.Error())
		return nil, err
	}

	logFile, err := os.OpenFile(absPath, os.O_CREATE|os.OWRONLY|os.O_APPEND, 0644)
	if err != nil {
		ui.pError("can't create log file: " + err.Error())
		return nil, err
	}

	logger = log.New(logFile, "[ASTMN] ", log.LstdFlags|log.Lshortfile)

	return logFile, nil
}

func Info(msg string) {
	if logger != nil {
		logger.Println("[INFO] ", msg)
	}
}

func Infof(format string, v ...any) {
	if logger != nil {
		logger.Printf("[INFO] "+format, v...)
	}
}

func Error(msg string) {
	if logger != nil {
		logger.Println("[ERROR]", msg)
	}
}

func Errorf(msg string, v ...any) {
	if logger != nil {
		logger.Printf("[ERROR]"+error, v...)
	}
}

