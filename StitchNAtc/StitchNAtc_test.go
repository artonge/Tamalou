package stitchnatc

import (
	"fmt"
	"log"
	"testing"

	"github.com/artonge/Tamalou/Models"
)

func TestStitchNAtc(t *testing.T) {
	var drugArray []*Models.Drug
	drug := new(Models.Drug)
	drug.STITCH_ID_SIDER = "CID100000085"
	//stitchIDs := []string{"CID100000085"}
	drugArray = append(drugArray, drug)
	err := GetChemicalFromID(drugArray)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(drugArray)
}
