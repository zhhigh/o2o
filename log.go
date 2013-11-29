package o2o

import (
	"strings"
	"github.com/zhhigh/o2o/logs"
	"fmt"
)

// Log levels to control the logging output.
const (
	LevelTrace = iota
	LevelDebug
	LevelInfo
	LevelWarning
	LevelError
	LevelCritical
)

// SetLogLevel sets the global log level used by the simple
// logger.
func SetLevel(l int) {
	Logger.SetLevel(l)
}

// logger references the used application logger.
var Logger *logs.Logger

func init() {
	Logger = logs.NewLogger(10000)
	fmt.Println(LogType)
	fmt.Println(LogFileName)
	switch LogType{
		case "file":
     		filePara := `{"filename":"`+LogFileName+`"}`
			fmt.Println(filePara)
	    	Logger.SetLogger(LogType, filePara)
		default:
		 Logger.SetLogger("console","")
	}

}

// SetLogger sets a new logger.
func SetLogger(adaptername string, config string) {
	Logger.SetLogger(adaptername, config)
}

// Trace logs a message at trace level.
func Trace(v ...interface{}) {
	Logger.Trace(generateFmtStr(len(v)), v...)
}

// Debug logs a message at debug level.
func Debug(v ...interface{}) {
	Logger.Debug(generateFmtStr(len(v)), v...)
}

// Info logs a message at info level.
func Info(v ...interface{}) {
	Logger.Info(generateFmtStr(len(v)), v...)
}

// Warning logs a message at warning level.
func Warn(v ...interface{}) {
	Logger.Warn(generateFmtStr(len(v)), v...)
}

// Error logs a message at error level.
func Error(v ...interface{}) {
	Logger.Error(generateFmtStr(len(v)), v...)
}

// Critical logs a message at critical level.
func Critical(v ...interface{}) {
	Logger.Critical(generateFmtStr(len(v)), v...)
}

func generateFmtStr(n int) string {
	return strings.Repeat("%v ", n)
}
