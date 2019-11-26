package log

import (
	"log"
)

// Based partially on the following links:
// https://github.com/go-kit/kit/blob/master/log/level/level.go
// https://www.ardanlabs.com/blog/2013/11/using-log-package-in-go.html
var Logger ResourceLogger

type ResourceLogger struct {
	Logger *log.Logger
	Level  int
	ready  bool
}

func (rl *ResourceLogger) InitLoggerFromParam(levelStringValue string) {
	rl.InitLogger(GetValueFromParam(levelStringValue))
}

func (rl *ResourceLogger) InitLogger(levelValue int) {
	if levelValue < MinLevel || levelValue > MaxLevel {
		rl.Level = INFO
	} else {
		rl.Level = levelValue
	}

	rl.Logger = log.New(GetIOHandle(rl.Level), GetPrefixForLogger(rl.Level),
		log.Ldate|log.Ltime)

	rl.ready = true
}

func (rl *ResourceLogger) Debug(vals ...interface{}) {
	rl.log(DEBUG, vals)
}

func (rl *ResourceLogger) Debugf(format string, vals ...interface{}) {
	rl.logf(DEBUG, format, vals)
}

func (rl *ResourceLogger) Info(vals ...interface{}) {
	rl.log(INFO, vals)
}

func (rl *ResourceLogger) Infof(format string, vals ...interface{}) {
	rl.logf(INFO, format, vals)
}

func (rl *ResourceLogger) Warning(vals ...interface{}) {
	rl.log(WARNING, vals)
}

func (rl *ResourceLogger) Warningf(format string, vals ...interface{}) {
	rl.logf(WARNING, format, vals)
}

func (rl *ResourceLogger) Error(vals ...interface{}) {
	rl.log(ERROR, vals)
}

func (rl *ResourceLogger) Errorf(format string, vals ...interface{}) {
	rl.logf(ERROR, format, vals)
}

func (rl *ResourceLogger) log(level int, vals ...interface{}) {
	if !rl.ready {
		rl.InitLogger(INFO)
	}

	if level >= rl.Level && rl.Level != OFF {
		rl.Logger.SetPrefix(GetPrefixForLogger(level))

		rl.Logger.Print(vals)
	}
}

func (rl *ResourceLogger) logf(level int, format string, vals ...interface{}) {
	if !rl.ready {
		rl.InitLogger(INFO)
	}

	if level >= rl.Level && rl.Level != OFF {
		rl.Logger.SetPrefix(GetPrefixForLogger(level))

		rl.Logger.Printf(format, vals)
	}
}
