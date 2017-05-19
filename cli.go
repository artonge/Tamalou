package main

import (
	"github.com/artonge/Tamalou/Queries"
	"github.com/mkideal/cli"
)

type argT struct {
	cli.Helper
	Query string `cli:"q,query" usage:"ventre AND tete OR hand"`
}

func startCLI() {
	cli.Run(new(argT), func(ctx *cli.Context) error {
		argv := ctx.Argv().(*argT)
		query := Queries.ParseQuery(argv.Query)

		diseases, err := fetchDiseases(query)
		if err != nil {
			ctx.String("Error:\n	%v\n", err)
		}
		drugs, err := fetchDrugs(query)
		if err != nil {
			ctx.String("Error:\n	%v\n", err)
		}

		// Print results
		ctx.String("Diseases (%v): \n", len(diseases))
		for _, r := range diseases {
			ctx.String("	- %v\n", r.Name)
		}
		ctx.String("Drugs (%v): \n", len(drugs))
		for _, d := range drugs {
			ctx.String("	- %v\n", d.Name)
		}

		return nil
	})
}
