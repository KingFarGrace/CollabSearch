package util

import log "github.com/sirupsen/logrus"

// ServerLogger
// TODO: Decompose this function.
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

// PlainErrorLogger can log plain errors,
// they are errors that terminate the function
// but won't shut down the system.
func PlainErrorLogger(err error, funcname string) {
	plainErrorEntry := log.WithFields(log.Fields{
		"Func": funcname,
	})
	plainErrorEntry.Error("Error: ", err)
}
