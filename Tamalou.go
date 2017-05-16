package main

import (
	orpha "github.com/artonge/Tamalou/Orpha"
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

		// Fetch diseases
		results, err := orpha.Query(Queries.ParseQuery(argv.Query))
		// ---

		// Print results
		if err != nil {
			ctx.String("Error:\n	%v\n", err)
		} else {
			ctx.String("Results (%v): \n", len(results))
			for _, r := range results {
				ctx.String("	- %v\n", r.Name)
			}
		}

		return nil
	})
}
