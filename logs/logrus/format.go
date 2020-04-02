package logrus

import "time"

type FormatLog struct {
	RandomString string    `json:"random"`
	OutPutTime   time.Time `json:"out_put_time"`
}

func NewFormatLog() *FormatLog {
	return &FormatLog{}
}
