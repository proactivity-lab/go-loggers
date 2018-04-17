// Author  Raido Pahtma
// License MIT

package loggers

import "log"
import "io/ioutil"
import "os"

type DIWElog interface {
	SetDebugLogger(logger *log.Logger)
	SetInfoLogger(logger *log.Logger)
	SetWarningLogger(logger *log.Logger)
	SetErrorLogger(logger *log.Logger)
	SetLoggers(loggers *DIWEloggers)
}

type DIWEloggers struct {
	Debug   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
}

func New() *DIWEloggers {
	logger := new(DIWEloggers)
	logger.InitLoggers()
	return logger
}

func (self *DIWEloggers) InitLoggers() {
	self.Debug = log.New(ioutil.Discard, "", 0)
	self.Info = log.New(ioutil.Discard, "", 0)
	self.Warning = log.New(ioutil.Discard, "", 0)
	self.Error = log.New(ioutil.Discard, "", 0)
}

func (self *DIWEloggers) SetDebugLogger(logger *log.Logger) {
	self.Debug = logger
}

func (self *DIWEloggers) SetInfoLogger(logger *log.Logger) {
	self.Info = logger
}

func (self *DIWEloggers) SetWarningLogger(logger *log.Logger) {
	self.Warning = logger
}

func (self *DIWEloggers) SetErrorLogger(logger *log.Logger) {
	self.Error = logger
}

func (self *DIWEloggers) SetLoggers(loggers *DIWEloggers) {
	self.Debug = loggers.Debug
	self.Info = loggers.Info
	self.Warning = loggers.Warning
	self.Error = loggers.Error
}

func BasicLogSetup(debuglevel int) *DIWEloggers {
	logger := New()
	logformat := log.Ldate | log.Ltime | log.Lmicroseconds

	if debuglevel > 1 {
		logformat = logformat | log.Lshortfile
	}

	if debuglevel > 0 {
		logger.SetDebugLogger(log.New(os.Stdout, "DEBUG: ", logformat))
		logger.SetInfoLogger(log.New(os.Stdout, "INFO:  ", logformat))
	} else {
		logger.SetInfoLogger(log.New(os.Stdout, "", logformat))
	}
	logger.SetWarningLogger(log.New(os.Stdout, "WARN:  ", logformat))
	logger.SetErrorLogger(log.New(os.Stdout, "ERROR: ", logformat))
	return logger
}
