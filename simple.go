package utilloggers


type SimpleLoggerF struct {
	LoggerF func(foramt string, v ...interface{})
}

func (logger *SimpleLoggerF) LogF(format string, v ...interface{}) {
	if logger.LoggerF != nil {
		logger.LoggerF(format, v)
	}
}

