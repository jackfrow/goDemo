package main

import (
	"github.com/jackfrow/goDemo/log/cuslog"
	"log"
	"os"
)

func main() {

	cuslog.Info("std log")
	cuslog.SetOptions(cuslog.WithLevel(cuslog.DebugLevel))
	cuslog.Debug("change std log to debug level")
	// 输出到文件
	fd, err := os.OpenFile("test.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln("create file test.log failed")
	}
	defer fd.Close()

	l := cuslog.New(cuslog.WithLevel(cuslog.InfoLevel),
		cuslog.WithOutput(fd),
		cuslog.WithFormatter(&cuslog.JsonFormatter{IgnoreBasicFields: false}),
	)
	l.Info("custom log with json formatter")
}
