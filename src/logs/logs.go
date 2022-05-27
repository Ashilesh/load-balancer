package logs

import (
	"io"
	"log"
	"os"

	"github.com/Ashilesh/load-balancer/utils"
)

var (
	infoLogger    *log.Logger
	errorLogger   *log.Logger
	warningLogger *log.Logger
	panicLogger   *log.Logger
)

func init() {
	path, argsErr := utils.GetCmdArgs("-log")

	logPath := path + "LB.log"

	logFile, err := os.OpenFile(logPath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		panic("cannot open log file")
	}
	mw := io.MultiWriter(os.Stdout, logFile)

	logFlags := log.Lmicroseconds | log.LstdFlags

	infoLogger = log.New(mw, "INFO : ", logFlags)
	errorLogger = log.New(mw, "ERROR: ", logFlags)
	warningLogger = log.New(mw, "WARN : ", logFlags)
	panicLogger = log.New(mw, "ERROR: ", logFlags)

	if argsErr != nil {
		Warn("log path not set, logs will be stored in current directory")
	} else {
		Info("log path ->", logPath)
	}
}

func Info(msg ...any) {
	infoLogger.Println(msg...)
}

func Error(errMsg ...any) {
	errorLogger.Println(errMsg...)
}

func Warn(warnMsg ...any) {
	warningLogger.Println(warnMsg...)
}

func Fatal(panicMsg ...any) {
	panicLogger.Fatal(panicMsg...)
}
