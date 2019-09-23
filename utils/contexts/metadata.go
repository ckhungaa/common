package contexts

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type ContextMetaData struct {
	Stan string
}

func defaultOpts(stan string) ContextMetaData {
	return ContextMetaData{Stan: stan}
}

func (opts ContextMetaData) toMetaData() metadata.MD {
	return metadata.MD{
		Stan.String(): []string{opts.Stan},
	}
}

func ReadMD(ctx context.Context) (*ContextMetaData, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		return &ContextMetaData{
			Stan: readMetaDataValue(Stan, md),
		}, nil
	}
	return nil, status.Errorf(codes.Aborted, "context is empty")
}

func readMetaDataValue(key ContextKey, md metadata.MD) string {
	val, mdOK := md[key.String()]
	if mdOK {
		if len(val) > 0 {
			return val[(len(val) - 1)]
		}
	}
	return ""
}
