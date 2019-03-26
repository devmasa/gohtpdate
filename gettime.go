package main

import (
	"net/http"
	"time"
)

// GetTimeFromHost 指定されたホストに対してhttpのHEADを要求して,
// レスポンスのHTTPHeader:Dateを取得しtime.Timeへ変換して返す
func GetTimeFromHost(host string) (time.Time, error) {
	resp, err := http.Head("http://" + host)
	if err != nil {
		return time.Time{}, err
	}
	httpDate := resp.Header.Get("Date")
	t, err := time.Parse(time.RFC1123, httpDate)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}
