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

// LaunchLogger can log every action to launch or start a service.
// Include: DB, Log, server etc.
func LaunchLogger(service string, isSuccess bool) {
	launchEntry := GetServiceEntry(service)
	launchEntry.Debug("Successfully launched? ", isSuccess)
}

// WarnLogger can log warning messages,
// they are caused by inappropriate operation
// that will lead to safe return or nothing.
func WarnLogger(funcname, msg string) {
	WarnEntry := GetFuncEntry(funcname)
	WarnEntry.Warning("Warning: ", msg)
}

// ErrorLogger can log plain errors,
// they are caused by bugs that only terminate
// running functions but never shut down the system.
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
