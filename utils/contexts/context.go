package contexts

import (
	"context"
	"google.golang.org/grpc/metadata"
)

//ContextKey meta data key in context
type ContextKey string

const (
	//Stan system tracing audit number
	Stan ContextKey = "x-stan"
	Token ContextKey = "x-token"
)

// String to string
func (key ContextKey) String() string {
	return string(key)
}

//ContextOption context options
type ContextOption func(*ContextMetaData)

//NewContext create new context
func NewContext(stan string, opt ...ContextOption) context.Context {
	curCtx := context.TODO()

	opts := defaultOpts(stan)
	for _, o := range opt {
		o(&opts)
	}
	curCtx = metadata.NewIncomingContext(curCtx, opts.toMetaData())
	return curCtx
}