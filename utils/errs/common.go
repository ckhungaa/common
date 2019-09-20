package errs

import "github.com/pkg/errors"

var (
	EmptyContext = errors.New("context does not contains metadata")
	InvalidId = errors.New("invalid id")
)
