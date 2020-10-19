package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"

	_ "github.com/ClickHouse/clickhouse-go"
	"github.com/alifpay/clickhz/slog"
	"github.com/alifpay/clickhz/storages/clickhouse"
	"github.com/alifpay/sqbuf"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	//connects to clickhouse db
	ch, err := clickhouse.Connect("tcp://127.0.0.1:9000?debug=true")
	if err != nil {
		log.Fatalln("clickhouse connect", err)
	}

	wg := sync.WaitGroup{}
	//init new log buffer for application
	qa := sqbuf.New(200, 500, ch.AppLog)
	wg.Add(1)
	qa.Run(ctx, &wg)
	slog.InitLog(qa)

	slog.Log("info", "main.go", "Server is started and running")

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, os.Kill, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		<-sigs
		slog.Log("info", "main.go", "Server is received shutdown signal")
		signal.Stop(sigs)
		close(sigs)
		cancel()
	}()

	wg.Add(1)
	go func() {
		for i := 1; i < 1500; i++ {
			slog.Log("info", "main.go first goroutine", "info message :"+strconv.Itoa(i))
		}
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		for i := 1500; i < 5001; i++ {
			slog.Log("waring", "main.go secong goroutine", "message :"+strconv.Itoa(i))
		}
		wg.Done()
	}()

	for i := 5001; i < 10001; i++ {
		slog.Log("error", "main.go third goroutine", "message :"+strconv.Itoa(i))
	}

	wg.Wait()
	ch.Close()
}
