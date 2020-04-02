package logrus

import (
	"github.com/sirupsen/logrus"
	"github.com/zuoshenglo/tools"
	"os"
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
	Log.WithFields(logrus.Fields{
		"type": "pa_access",
		"msg":  logFormat,
	}).Info(args...)
}

func Debug(args ...interface{}) {
	logFormat := logFormat()
	Log.WithFields(logrus.Fields{
		"type": "pa_access",
		"msg":  logFormat,
	}).Debug(args...)
}

func Warn(args ...interface{}) {
	logFormat := logFormat()
	Log.WithFields(logrus.Fields{
		"type": "pa_access",
		"msg":  logFormat,
	}).Warn(args...)
}

func Error(args ...interface{}) {
	logFormat := logFormat()
	Log.WithFields(logrus.Fields{
		"type": "pa_access",
		"msg":  logFormat,
	}).Error(args...)
}
