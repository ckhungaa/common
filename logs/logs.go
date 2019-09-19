package logs

import (
	"context"
	"fmt"
	"log"
)

type Logger struct {
	name string
}

func NewLogger(name string) *Logger{
	return &Logger{name: name}
}

func (o *Logger) Info(ctx context.Context, message string)  {
	log.Printf("[INFO]-%s: %s\n", o.name, message)
}

func (o *Logger) Infof(ctx context.Context, format string, v ...interface{})  {
	log.Printf(fmt.Sprintf("[INFO]-%s: %s\n", o.name, format), v...)
}

func (o *Logger) Error(ctx context.Context, message string)  {
	log.Printf("[ERROR]-%s: %s\n", o.name, message)
}

func (o *Logger) Errorf(ctx context.Context, format string, v ...interface{})  {
	log.Printf(fmt.Sprintf("[ERROR]-%s: %s\n", o.name, format), v...)
}