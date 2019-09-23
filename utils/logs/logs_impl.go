package logs

import (
	"context"
	"fmt"
	"github.com/ckhungaa/common/utils/contexts"
	"go.uber.org/zap"
)

type logLevel int

const (
	debugLevel logLevel = iota
	infoLevel
	warnLevel
	errorLevel
	fatalLevel
)

func (log *Log) Debugf(ctx context.Context, format string, args ...interface{}) {
	log.printf(ctx, debugLevel, nil, format, args...)
}

func (log *Log) Infof(ctx context.Context, format string, args ...interface{}) {
	log.printf(ctx, infoLevel, nil, format, args...)
}

func (log *Log) Warnf(ctx context.Context, format string, args ...interface{}) {
	log.printf(ctx, warnLevel, nil, format, args...)
}

func (log *Log) Warne(ctx context.Context, err error, format string, args ...interface{}) {
	log.printf(ctx, warnLevel, err, format, args...)
}

func (log *Log) Errorf(ctx context.Context, format string, args ...interface{}) {
	log.printf(ctx, errorLevel, nil, format, args...)
}

func (log *Log) Errore(ctx context.Context, err error, format string, args ...interface{}) {
	log.printf(ctx, errorLevel, err, format, args...)
}

func (log *Log) Fatale(ctx context.Context, err error, format string, args ...interface{}) {
	log.printf(ctx, fatalLevel, err, format, args...)
}

func (log *Log) printf(ctx context.Context, level logLevel, err error, format string, args ...interface{}) {
	var (
		stan string
	)
	defer log.logger.Sync()
	if md, err := contexts.ReadMD(ctx); err == nil {
		if md != nil {
			stan = md.Stan
		}
	}

	switch level {
	case debugLevel:
		log.logger.Debug(fmt.Sprintf(format, args...), zap.String("stan", stan))
	case infoLevel:
		log.logger.Info(fmt.Sprintf(format, args...), zap.String("stan", stan))
	case warnLevel:
		log.logger.Warn(fmt.Sprintf(format, args...), zap.String("stan", stan))
	case errorLevel:
		log.logger.Error(fmt.Sprintf(format, args...), zap.String("stan", stan), zap.Error(err))
	case fatalLevel:
		log.logger.Fatal(fmt.Sprintf(format, args...), zap.String("stan", stan), zap.Error(err))
	default:
		log.logger.Debug(fmt.Sprintf(format, args...), zap.String("stan", stan))
	}
}
