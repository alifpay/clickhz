package slog

import (
	"log"
	"time"

	"github.com/alifpay/sqbuf"
)

var buf *sqbuf.Queue

//InitLog -
func InitLog(b *sqbuf.Queue) {
	buf = b
}

//Log add application logs to queue for saving
func Log(errType, caller, message string) {
	/*
		the order of the arguments is very important, because of data in clickhouse

		INSERT INTO app_logs(errType, caller, message, dt) VALUES (?, ?, ?, ?)
	*/
	err := buf.Add(errType, caller, message, time.Now())
	if err != nil {
		log.Println("internal.Log buf.Add", err)
	}
}
