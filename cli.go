package main

import (
	"fmt"
	"os"

	"github.com/artonge/Tamalou/HPO"
	"github.com/artonge/Tamalou/Omim"
	"github.com/artonge/Tamalou/Queries"
	stitchnatc "github.com/artonge/Tamalou/StitchNAtc"
	"github.com/mkideal/cli"
)

func startCLI() {
	if err := cli.Root(tamalouCMD,
		cli.Tree(help),
		cli.Tree(indexCMD),
		cli.Tree(serverCMD),
	).Run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

var help = cli.HelpCommand("display help information")

// root command
type tamalouCMDT struct {
	cli.Helper
	Name  string `cli:"tamalou" usage:"[index, request]"`
	Query string `cli:"q,query" usage:"ventre AND tete OR hand"`
}

var tamalouCMD = &cli.Command{
	Desc: "this is tamalou root command",
	// Argv is a factory function of argument object
	// ctx.Argv() is if Command.Argv == nil or Command.Argv() is nil
	Argv: func() interface{} { return new(tamalouCMDT) },
	Fn: func(ctx *cli.Context) error {
		argv := ctx.Argv().(*tamalouCMDT)
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
	},
}

// child command
type indexCMDT struct {
	cli.Helper
	Name string `cli:"index" usage:"symptom1 AND sumptom2 OR symptom3"`
}

var indexCMD = &cli.Command{
	Name: "index",
	Desc: "this is the index command",
	Argv: func() interface{} { return new(indexCMDT) },
	Fn: func(ctx *cli.Context) error {
		err := HPO.IndexHPO()
		if err != nil {
			ctx.String("Error:\n	%v\n", err)
		}
		err = Omim.IndexOmim()
		if err != nil {
			ctx.String("Error:\n	%v\n", err)
		}
		err = stitchnatc.IndexStitchNAtc()
		if err != nil {
			ctx.String("Error:\n	%v\n", err)
		}
		return nil
	},
}

// child command
type serverCMDT struct {
	cli.Helper
	Name string `cli:"server" usage:"symptom1 AND sumptom2 OR symptom3"`
}

var serverCMD = &cli.Command{
	Name: "server",
	Desc: "this is the server command",
	Argv: func() interface{} { return new(serverCMDT) },
	Fn: func(ctx *cli.Context) error {
		startServer()
		return nil
	},
}
