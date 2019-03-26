package main

import (
	"os"
	"syscall"
	"time"
)

// SetTimeOfDay Time構造体を受け取りOSの時間設定を行う
func SetTimeOfDay(t time.Time) error {
	tv := syscall.Timeval{Sec: t.Unix(), Usec: 0}
	return os.NewSyscallError("settimeofday", syscall.Settimeofday(&tv))
}
