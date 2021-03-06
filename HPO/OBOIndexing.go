package HPO

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/artonge/Tamalou/indexing"
)

func indexOBO() error {
	// Open the obo file
	file, err := os.Open("datas/hpo/hp.obo")
	if err != nil {
		return fmt.Errorf("Error in HPO's obo connector init\n	Error	==> %v", err)

	}
	// Create a new Scanner to parse the file
	scanner := bufio.NewScanner(file)

	err = indexing.IndexDocs(index, func() (indexing.Indexable, error) {
		return nextTerm(scanner)
	})

	return err
}

// Return the next term
func nextTerm(scanner *bufio.Scanner) (*hpoOBOStruct, error) {
	// Go to the next [Term]
	for scanner.Scan() {
		if scanner.Text() == "[Term]" {
			break
		}
	}

	// Init the new term
	term := new(hpoOBOStruct)

	// Continue the file parsing from the last position
	for scanner.Scan() {
		switch scanner.Text() {
		// End of a Term
		case "":
			return term, nil
		// Properties of a Term
		default:
			lineParts := strings.SplitN(scanner.Text(), ": ", 2)
			switch lineParts[0] {
			case "id":
				term.ID = lineParts[1]
			case "alt_id":
				term.AltIDs = append(term.AltIDs, lineParts[1])
			case "name":
				term.Name = lineParts[1]
			case "def":
				term.Definition = lineParts[1]
			case "comment":
				term.Comment = lineParts[1]
			case "synonym":
				term.Synonymes = append(term.Synonymes, lineParts[1])
			case "xref":
				term.Xrefs = append(term.Xrefs, lineParts[1])
			case "is_a":
				term.IsA = lineParts[1]
			case "is_obsolete": // use with consider ?
				term.Obsolete = true
			case "created_by", "property_value", "replaced_by", "creation_date", "subset", "is_anonymous", "consider":
			default:
				fmt.Println("Warning: Unexpected field <", lineParts[0], "> during obo file parsing")
			}
		}
	}

	return term, io.EOF
}
