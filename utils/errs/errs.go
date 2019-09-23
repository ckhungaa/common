package main

import (
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

// testing
/**
gRPC error
	code.Code_INTERNAL database connection err, unknown error, retry
	code.Code_ABORTED like internal server error, implementation error, incorrect data, no retry

 */

func main() {
	//abc := &errdetails.BadRequest{
	//	FieldViolations:      []*errdetails.BadRequest_FieldViolation{
	//		&errdetails.BadRequest_FieldViolation{
	//			Field:                "abc",
	//			Description:          "def",
	//		},
	//	},
	//}
	errStatus, _ := status.New(codes.InvalidArgument, "test").WithDetails(&errdetails.PreconditionFailure{
		Violations:           []*errdetails.PreconditionFailure_Violation{
			&errdetails.PreconditionFailure_Violation{
				Type:                 "test",
				Subject:              "test",
				Description:          "test",
			},
		},
	})
	log.Printf("errStatus.Err(): %s", errStatus.Err())
	log.Printf("errStatus.Err(): %+v", errStatus.Details())

	st := status.Convert(errStatus.Err())
	log.Printf("errStatus.Err(): %+v", st.Details())
}