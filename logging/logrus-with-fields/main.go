package main

import (
	"github.com/sirupsen/logrus"
)

func main() {
	workerLogger := logrus.WithFields(logrus.Fields{
		"source": "worker",
	})
	workerLogger.Info("worker has finished processed task")

	mysqlLogger := logrus.WithFields(logrus.Fields{
		"source": "db",
		"dbType": "mysql",
	})
	mysqlLogger.Error("impossible to establish connexion")

}
