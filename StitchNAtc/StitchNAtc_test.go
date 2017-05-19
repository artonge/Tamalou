package stitchnatc

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"reflect"
	"testing"
)

func TestStitchNAtc(t *testing.T) {
	file, err := os.Open("/media/carl/DATA/Downloads/chemical.sources.v5.0.tsv/chemical.sources.v5.0.tsv")
	if err != nil {
		log.Fatal(err)
	}
	reader := csv.NewReader(file)
	reader.Comma = '\t'
	reader.Comment = '#'
	for i := 0; i < 10; i++ {
		line, err := reader.Read()
		fmt.Println(reflect.TypeOf(err))
		if err != nil {
			switch err {
			case csv.ParseError:
				fmt.Print("it works")
			default:
				fmt.Println("hahahaha")
			}
		}
		fmt.Println(line, err)
	}
}
