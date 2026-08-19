package log

import (
	"fmt"
	"io"
	stdlog "log"
	"os"
	"path/filepath"
	"time"

	"astmn/internal/ui"
)

var logger *stdlog.Logger

func InitLogger(logDir string) (*os.File, error) {
	if err := os.MkdirAll(logDir, 0755); err != nil {
		ui.PError("can't create log dir: " + err.Error())
		return nil, err
	}

	logFileName := time.Now().Format("2006-01-02_15-04") + "-astmn.log"

	fullPath := filepath.Join(logDir, logFileName)
	absPath, err := filepath.Abs(fullPath)
	if err != nil {
		ui.PError("can't get abs path: " + err.Error())
		return nil, err
	}

	logFile, err := os.OpenFile(absPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		ui.PError("can't create log file: " + err.Error())
		return nil, err
	}

	multiWriter := io.MultiWriter(os.Stdout, logFile)

	logger = stdlog.New(multiWriter, "[ASTMN] ", stdlog.LstdFlags|stdlog.Lshortfile)

	return logFile, nil
}

func Info(msg string) {
	if logger != nil {
		logger.Output(2, "[INFO] "+msg)
	}
}

func Infof(format string, v ...any) {
	if logger != nil {
		logger.Output(2, fmt.Sprintf("[INFO] "+format, v...))
	}
}

func Error(msg string) {
	if logger != nil {
		logger.Output(2, "[ERROR] "+msg)
	}
}

func Errorf(format string, v ...any) {
	if logger != nil {
		logger.Output(2, fmt.Sprintf("[ERROR] "+format, v...))
	}
}
