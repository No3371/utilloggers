package utilloggers

type SimpleLoggerF struct {
	LoggerF func(foramt string, v ...interface{})
}

func (logger *SimpleLoggerF) LogF(format string, v ...interface{}) {
	if logger.LoggerF != nil {
		if v == nil {
			logger.LoggerF(format)
		} else {
			logger.LoggerF(format, v...)
		}
	}
}

type SimpleLogger struct {
	Logger func(v string)
}

func (logger *SimpleLogger) LogF(v string) {
	if logger.Logger != nil {
		logger.Logger(v)
	}
}
