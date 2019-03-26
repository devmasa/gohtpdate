package main

import (
	"flag"
	"log/syslog"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	log "github.com/sirupsen/logrus"
)

// EnvProxy プロキシのOS環境変数
const (
	EnvProxy = "HTTP_PROXY"
)

var (
	defaultProxy string
	proxy        string
	host         string
	interval     int64
	resetProxy   = func() {}
)

// OSの環境変数を一時的に上書きして元に戻す関数を返す
func setTempEnv(key, val string) func() {
	if key == "" {
		return func() {}
	}
	preVal := os.Getenv(key)
	os.Setenv(key, val)
	return func() {
		os.Setenv(key, preVal)
	}
}

// 処理本体
func gohtpdate() {
	// 指定されたドメインからhttpで時間取得してパース
	t, err := GetTimeFromHost(host)
	if err != nil {
		log.Errorf("Failed GetTimeFromHost  : %v\n", err)
		return
	}
	log.Infof("local : %v\n", time.Now())
	log.Infof("%v : %v\n", host, t)
	// os時間設定(ここでルート権限必要)
	if err := SetTimeOfDay(t); err != nil {
		log.Errorf("failed SetTimeOfDay : %v\n", err)
		return
	}
	log.Infof("Adjusted : %v\n", time.Now())
}

func init() {
	// ログ初期化
	log.SetLevel(log.InfoLevel)
	logger, err := syslog.New(syslog.LOG_NOTICE|syslog.LOG_USER, "gohtpdate")
	if err != nil {
		panic(err)
	}
	log.SetOutput(logger)

	// コマンド引数初期化
	defaultProxy = os.Getenv(EnvProxy)
	flag.StringVar(&proxy, "p", defaultProxy, "proxy&port")
	flag.StringVar(&host, "h", "www.google.co.jp", "host&port")
	flag.Int64Var(&interval, "i", 60, "interval minutes")
	flag.Parse()
	// プロキシ指定されてたら戻し関数をセット
	if proxy != defaultProxy {
		resetProxy = setTempEnv(EnvProxy, proxy)
	}
}

func main() {
	// 色々チャンネル設定
	quitChannel := make(chan os.Signal)
	signal.Notify(
		quitChannel,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	shutdownChannel := make(chan bool)
	waitGroup := &sync.WaitGroup{}
	waitGroup.Add(1)
	intervalChannel := time.Tick(time.Duration(interval) * time.Minute)

	go func() {
		defer waitGroup.Done()
		for {
			select {
			case <-shutdownChannel:
				log.Info("Received quit.")
				return
			case <-intervalChannel:
				gohtpdate()
			}
		}
	}()

	// Wait until we get the quit message
	<-quitChannel
	shutdownChannel <- true
	// Block until wait group counter gets to zero
	waitGroup.Wait()
	resetProxy()
	os.Exit(0)
}
