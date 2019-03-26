package main

import (
	"testing"
)

// 取得できてパースできるか
func TestGetTime(t *testing.T) {
	host := "www.google.co.jp"
	_, err := GetTimeFromHost(host)
	if err != nil {
		t.Errorf("can not parse time format. %v", err)
	}
}
