package contexts

import (
	"context"
	"github.com/ckhungaa/common/utils/errs"
	"google.golang.org/grpc/metadata"
)

type ContextMetaData struct {
	Stan string
}

func defaultOpts(stan string) ContextMetaData {
	return ContextMetaData{Stan:stan}
}

func (opts ContextMetaData) toMetaData() metadata.MD {
	return metadata.MD{
		Stan.String() : []string{opts.Stan},
	}
}

func ReadMD(ctx context.Context) (*ContextMetaData, error)  {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		return &ContextMetaData{
			Stan: readMetaDataValue(Stan, md),
		}, nil
	}
	return nil, errs.EmptyContext
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