package utilloggers

type LeveledLoggerF struct {
	InfoLoggerF  func(foramt string, v ...interface{})
	ErrorLoggerF func(foramt string, v ...interface{})
	MiscLoggerF  func(foramt string, v ...interface{})
	SpamLoggerF  func(foramt string, v ...interface{})
}

func (logger *LeveledLoggerF) LogErrorF(format string, v ...interface{}) {
	if logger.ErrorLoggerF != nil {
		if v == nil {
			logger.ErrorLoggerF(format)
		} else {
			logger.ErrorLoggerF(format, v...)
		}
	}
}

func (logger *LeveledLoggerF) LogInfoF(format string, v ...interface{}) {
	if logger.InfoLoggerF != nil {
		if v == nil {
			logger.InfoLoggerF(format)
		} else {
			logger.InfoLoggerF(format, v...)
		}
	}
}

func (logger *LeveledLoggerF) LogMiscF(format string, v ...interface{}) {
	if logger.MiscLoggerF != nil {
		if v == nil {
			logger.MiscLoggerF(format)
		} else {
			logger.MiscLoggerF(format, v...)
		}
	}
}

func (logger *LeveledLoggerF) LogSpamF(format string, v ...interface{}) {
	if logger.SpamLoggerF != nil {
		if v == nil {
			logger.SpamLoggerF(format)
		} else {
			logger.SpamLoggerF(format, v...)
		}
	}
}
