package zerogokit

import (
	log2 "github.com/go-kit/kit/log"
	"github.com/rs/zerolog"
)

type basicZeroGokit struct {
	level  zerolog.Level
	logger zerolog.Logger
}

func (z basicZeroGokit) Log(keyvals ...interface{}) error {
	logger := z.logger.WithLevel(z.level).Timestamp().Caller(1)
	logger = compose(logger, keyvals)
	logger.Msg("")
	return nil
}

func compose(logger *zerolog.Event, keyvals []interface{}) *zerolog.Event {
	for n := 0; n < len(keyvals); n += 2 {
		if n >= len(keyvals) {
			return logger
		}
		if n+1 == len(keyvals) {
			logger = logger.Interface(handleKey(keyvals[n]), nil)
			return logger
		}
		logger = logger.Interface(handleKey(keyvals[n]), keyvals[n+1])
	}
	return logger
}

func handleKey(key interface{}) string {
	if key == nil {
		return ""
	}
	return key.(string)
}

func NewZeroGokit(logger zerolog.Logger, level zerolog.Level) log2.Logger {
	return &basicZeroGokit{logger: logger, level: level}
}
