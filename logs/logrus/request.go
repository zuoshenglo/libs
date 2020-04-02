package logrus

import "time"

type RequestLog struct {
	Request string `json:"request"`
	RequestUrl string `json:"request_url"`
	RemoteAddr string `json:"remote_addr"`
	RequestBody string `json:"request_body"`
	RequestTime time.Duration `json:"request_time"`
	Status int `json:"status"`
	UpstreamResponseTime float64 `json:"upstream_response_time"`
	StartTime time.Time `json:"start_time"`
	EndTime time.Time `json:"end_time"`
	ReqMethod string `json:"req_method"`
	ClientIP string `json:"client_ip"`
}

func NewRequestLog() * RequestLog {
	return &RequestLog{}
}