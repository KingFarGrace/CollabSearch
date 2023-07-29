package conf

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func InitLogger() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.WithFields(log.Fields{
		"logrus": "v1.9.3",
		"format": "json",
	}).Info("Logger init successfully.")
}
