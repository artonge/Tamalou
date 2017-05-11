package main

import (
	"fmt"

	"github.com/blevesearch/bleve"
	"github.com/mkideal/cli"
)

type argT struct {
	cli.Helper
	Or  []string `cli:"or"  usage:"or logical operator"`
	And []string `cli:"and" usage:"and logical operator"`
}

func main() {
	cli.Run(new(argT), func(ctx *cli.Context) error {
		argv := ctx.Argv().(*argT)
		ctx.String("or=%v, and=%v\n", argv.Or, argv.And)
		return nil
	})
}

// Bleve example
func bleveExample() {
	// open a new index
	mapping := bleve.NewIndexMapping()
	index, err := bleve.New("example.bleve", mapping)
	if err != nil {
		fmt.Println(err)
		return
	}

	data := struct {
		Name string
	}{
		Name: "text",
	}

	// index some data
	index.Index("id", data)

	// search for some text
	query := bleve.NewMatchQuery("text")
	search := bleve.NewSearchRequest(query)
	searchResults, err := index.Search(search)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(searchResults)
}
