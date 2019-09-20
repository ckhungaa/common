package main

import (
	"github.com/ckhungaa/common/utils/contexts"
	"github.com/ckhungaa/common/utils/logs"
)

func main() {
	ctx := contexts.NewContext("test")
	var log = logs.NewLogger("test")
	log.Infof(ctx, "test begin")
}