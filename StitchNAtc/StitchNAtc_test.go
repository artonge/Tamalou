package stitchnatc

import (
	"fmt"
	"log"
	"testing"
)

func TestStitchNAtc(t *testing.T) {
	stitchIDs := []string{"CID100000085"}
	results, err := GetChemicalFromID(stitchIDs)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(results)
}
