package main

import (
	"log"
	"log/syslog"
)

func main() {
	logWriter, err := syslog.New(syslog.LOG_WARNING|syslog.LOG_DAEMON, "loggingTestProgram")
	if err != nil {
		log.Fatal(err)
	}
	_ = logWriter.Emerg("emergency sent tot syslog. TEST2")
}

// Демон — это программа, которая постоянно работает в фоновом режиме.
// Журнал демона — это программа, которая отвечает за сбор логов
// приложений
