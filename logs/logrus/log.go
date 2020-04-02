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
	Log.SetReportCaller(true)
	Log.Formatter = &logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "date",
			logrus.FieldKeyLevel: "level",
			logrus.FieldKeyMsg:   "message",
			logrus.FieldKeyFunc:  "caller",
		},
	}

	var requestLog = NewFormatLog()
	requestLog.RandomString = tools.GetRandomString(10)
	requestLog.OutPutTime = time.Now()
	Log.WithFields(logrus.Fields{
		"type": "pa_access",
		"msg":  requestLog,
	}).Info("common-info")
}

func Info(args ...interface{}) {
	var formatLog = NewFormatLog()
	formatLog.RandomString = tools.GetRandomString(10)
	formatLog.OutPutTime = time.Now()
	Log.WithFields(logrus.Fields{
		"type": "pa_access",
		"msg":  formatLog,
	}).Info("common-info")
	Log.Info(args...)
}

func Debug(args ...interface{}) {
	Log.Debug(args...)
}

func Warn(args ...interface{}) {
	Log.Warn(args...)
}

func Error(args ...interface{}) {
	Log.Error(args...)
}

func Test1111111(args ...interface{})  {

}