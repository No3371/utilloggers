package utilloggers

type LeveledLoggerF struct {
	InfoLoggerF  func(foramt string, v ...interface{})
	ErrorLoggerF func(foramt string, v ...interface{})
	MiscLoggerF  func(foramt string, v ...interface{})
	SpamLoggerF  func(foramt string, v ...interface{})
}

func (logger *LeveledLoggerF) LogErrorF(format string, v ...interface{}) {
	if logger.ErrorLoggerF != nil {
		logger.ErrorLoggerF(format, v)
	}
}

func (logger *LeveledLoggerF) LogInfoF(format string, v ...interface{}) {
	if logger.InfoLoggerF != nil {
		logger.InfoLoggerF(format, v)
	}
}

func (logger *LeveledLoggerF) LogMiscF(format string, v ...interface{}) {
	if logger.MiscLoggerF != nil {
		logger.MiscLoggerF(format, v)
	}
}

func (logger *LeveledLoggerF) LogSpamF(format string, v ...interface{}) {
	if logger.SpamLoggerF != nil {
		logger.SpamLoggerF(format, v)
	}
}
