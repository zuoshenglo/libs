package logrus

import (
	"github.com/sirupsen/logrus"
	"github.com/zuoshenglo/tools"
	"os"
	"runtime"
	"time"
)

var Log = logrus.New()

func init() {
	Log.Out = os.Stdout
	//Log.SetReportCaller(true)
	Log.Formatter = &logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "date",
			logrus.FieldKeyLevel: "level",
			logrus.FieldKeyMsg:   "message",
			logrus.FieldKeyFunc:  "caller",
		},
	}
}

func logFormat() *FormatLog {
	var logFormat = NewFormatLog()
	logFormat.RandomString = tools.GetRandomString(10)
	logFormat.OutPutTime = time.Now()
	return logFormat
}

func Info(args ...interface{}) {
	logFormat := logFormat()
	funcNo, file, line, _ := runtime.Caller(1)
	funcName := runtime.FuncForPC(funcNo).Name()
	Log.WithFields(logrus.Fields{
		"type":     "pa_access",
		"msg":      logFormat,
		"funcName": funcName,
		"file":     file,
		"line":     line,
	}).Info(args...)
}

func Debug(args ...interface{}) {
	logFormat := logFormat()
	funcNo, file, line, _ := runtime.Caller(1)
	funcName := runtime.FuncForPC(funcNo).Name()
	Log.WithFields(logrus.Fields{
		"type":     "pa_access",
		"msg":      logFormat,
		"funcName": funcName,
		"file":     file,
		"line":     line,
	}).Debug(args...)
}

func Warn(args ...interface{}) {
	logFormat := logFormat()
	funcNo, file, line, _ := runtime.Caller(1)
	funcName := runtime.FuncForPC(funcNo).Name()
	Log.WithFields(logrus.Fields{
		"type":     "pa_access",
		"msg":      logFormat,
		"funcName": funcName,
		"file":     file,
		"line":     line,
	}).Warn(args...)
}

func Error(args ...interface{}) {
	logFormat := logFormat()
	funcNo, file, line, _ := runtime.Caller(1)
	funcName := runtime.FuncForPC(funcNo).Name()
	Log.WithFields(logrus.Fields{
		"type":     "pa_access",
		"msg":      logFormat,
		"funcName": funcName,
		"file":     file,
		"line":     line,
	}).Error(args...)
}

func Fatalf(args ...interface{}) {
	logFormat := logFormat()
	funcNo, file, line, _ := runtime.Caller(1)
	funcName := runtime.FuncForPC(funcNo).Name()
	Log.WithFields(logrus.Fields{
		"type":     "pa_access",
		"msg":      logFormat,
		"funcName": funcName,
		"file":     file,
		"line":     line,
	}).Fatal(args...)
}

func WithFields(fields logrus.Fields) {
	logFormat := logFormat()
	funcNo, file, line, _ := runtime.Caller(1)
	funcName := runtime.FuncForPC(funcNo).Name()
	Log.WithFields(logrus.Fields{
		"type":     "pa_access",
		"msg":      logFormat,
		"test-msg": fields,
		"funcName": funcName,
		"file":     file,
		"line":     line,
	}).Info("this is a test")
}
