package main

import (
	"testing"
	"time"
)

// 単純に時間設定できるか(実行にはsudo必要)
func TestSetTimeOfDay(t *testing.T) {
	tm := time.Now()
	err := SetTimeOfDay(tm)
	if err != nil {
		t.Errorf("can not set time. %v", err)
	}
}
