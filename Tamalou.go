package main

import (
	"fmt"

	"github.com/artonge/Tamalou/Queries"
	"github.com/mkideal/cli"
)

type argT struct {
	cli.Helper
	Query string `cli:"q,query" usage:"ventre AND tete OR hand"`
}

func main() {
	cli.Run(new(argT), func(ctx *cli.Context) error {
		argv := ctx.Argv().(*argT)
		ctx.String("query=%v\n", fmt.Sprint(Queries.ParseQuery(argv.Query)))
		return nil
	})
}
