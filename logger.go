package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"time"
)

func Init() {
	logsDir := getLogDir()
	setLogFile(logsDir)
}

func setLogFile(logsDir string) {
	logFileName := fmt.Sprintf("log_%s.log", time.Now().Format("2006-01-02"))
	logFilePath := path.Join(logsDir, logFileName)
	f, err := os.OpenFile(logFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	wrt := io.MultiWriter(os.Stdout, f)
	log.SetOutput(wrt)
}

func getLogDir() string {

	exec, err := os.Executable()
	if err != nil {
		panic(err)
	}

	logsDir := path.Join(path.Dir(exec), "logs")

	if _, err := os.Stat(logsDir); err == os.ErrNotExist {
		err := os.Mkdir(logsDir, 0777)
		if err != nil {
			log.Println(err)
			log.Fatalf("Error: can't create log directory %v", logsDir)
		}
	}

	return logsDir
}
