package util

import log "github.com/sirupsen/logrus"

// GetFuncEntry to get an entry which describe a function.
func GetFuncEntry(funcname string) *log.Entry {
	return log.WithField("Function", funcname)
}

// GetServiceEntry to get an entry which describe a service.
func GetServiceEntry(service string) *log.Entry {
	return log.WithField("Service", service)
}

// LaunchLogger will log every action to launch or start a service.
// Include: DB, Log, server etc.
func LaunchLogger(service string, isSuccess bool) {
	launchEntry := GetServiceEntry(service)
	launchEntry.Debug("Successfully launched? ", isSuccess)
}

// ErrorLogger can log plain errors,
// they are errors that terminate the function
// but won't shut down the system.
func ErrorLogger(err error, funcname string) {
	errorEntry := GetFuncEntry(funcname)
	errorEntry.Error("Error: ", err)
}

// FatalLogger can log fatal errors,
// they will shut down the system or cause fatal error later.
func FatalLogger(err error, funcname string) {
	fatalEntry := GetFuncEntry(funcname)
	fatalEntry.Fatal("Error: ", err)
}
