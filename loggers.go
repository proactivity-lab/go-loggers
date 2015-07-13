// Author  Raido Pahtma
// License MIT

package loggers

import "log"
import "io/ioutil"

type DIWElog interface {
	SetDebugLogger(logger *log.Logger)
	SetInfoLogger(logger *log.Logger)
	SetWarningLogger(logger *log.Logger)
	SetErrorLogger(logger *log.Logger)
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
