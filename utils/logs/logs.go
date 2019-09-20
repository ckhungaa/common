package logs

import (
	"go.uber.org/zap"
	"sync"
)

var (
	mux sync.Mutex
	logFactory *zap.Logger
)

type Log struct {
	logger *zap.Logger
}

func provideLogFactory() (*zap.Logger, error) {
	return zap.NewProduction()
}

func NewLogger(name string) *Log {
	mux.Lock()
	if logFactory == nil {
		logFactory, _ =  provideLogFactory()
		logFactory = logFactory.WithOptions(zap.AddCallerSkip(2))
	}
	mux.Unlock()

	return &Log{logger: logFactory.Named(name)}
}