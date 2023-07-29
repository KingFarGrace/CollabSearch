package util

import log "github.com/sirupsen/logrus"

func ServerLogger(ip string, msg string, level string) {
	serverEntry := log.WithField("Server", ip)
	switch level {
	case "Debug":
		serverEntry.Debug(msg)
	case "Info":
		serverEntry.Info(msg)
	case "Warn":
		serverEntry.Warn(msg)
	case "Error":
		serverEntry.Error(msg)
	case "Fatal":
		serverEntry.Fatal(msg)
	case "Panic":
		serverEntry.Panic(msg)
	default:
		serverEntry.Error("Unknown logger level.")
	}
}
