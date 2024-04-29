package main

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

func main() {
	log := logrus.New()

	f, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("error opening file")
	}

	mw := io.MultiWriter(f, os.Stdout)
	log.SetOutput(mw)

	// установление уровня логгирования
	// (по умолчанию трассировка и отладка не выводятся)
	log.SetLevel(logrus.TraceLevel)

	log.Trace("going to mars")
	log.Debug("connected, received buffer from worker")
	log.Info("connection successful to db")
	log.Warn("something went wrong with x")
	log.Error("an error occurred in worker x")
	log.Fatal("impossible to continue exec")
	log.Panic("system emergency shutdown")
}
