package HPO

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	keywords = []string{
		"id",
		"alt_id",
		"name",
		"def",
		"synonyme",
		"xref",
		"is_a",
	}

	f       *os.File
	scanner *bufio.Scanner
	term    *HPOOBOStruct
)

// Open a connection to the hp.obo file
func init() {
	var err error
	f, err = os.Open("../datas/hpo/hp.obo")
	if err != nil {
		fmt.Println("Error in HPO's obo connector init: ", err)
	}
}

// On a new Query :
//  - Create a new Scanner
//  - Parse the hp.obo file
//  - For each term, submit it to the Query
//    - If the query match, add it to the results
func HPOOBOQuery(query map[string]interface{}) ([]*HPOOBOStruct, error) {
	// Create a new Scanner
	scanner = bufio.NewScanner(f)

	// Ignore the first part of the file
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
	}

	// Init the results array
	results := make([]*HPOOBOStruct, 0, 100)

	// Loop through the Terms
	for {
		term, err := nextTerm()
		if err != nil {
			return nil, fmt.Errorf("Error while parsing obo file\n	==> ", err)
		}
		// End of the file
		if term == nil {
			break
		}

		// Add term that match the query to the results array
		if termMatchesQuery(term, query, "") {
			results = append(results, term)
		}
	}

	return results, nil
}

// Test a Term against a query
func termMatchesQuery(term *HPOOBOStruct, query map[string]interface{}, queryType string) bool {
	var answers []bool

	// For each elements of the query
	for key, value := range query {
		switch key {
		case "and", "or":
			// Recursive call for 'or' and 'and' nodes
			answers = append(
				answers,
				termMatchesQuery(term, value.(map[string]interface{}), key),
			)
		default:
			// Normal comparison for others
			switch key {
			case "id":
				if term.ID == value {
					answers = append(answers, true)
					break
				}
				for _, id := range term.AltIDs {
					if id == value {
						answers = append(answers, true)
						break
					}
				}
				answers = append(answers, false)
			case "name":
				answers = append(answers, term.Name != value)
			case "def":
				answers = append(answers, term.Definition != value)
			case "comment":
				answers = append(answers, term.Comment != value)
			case "synonym":
				for _, synonyme := range term.Synonymes {
					if synonyme == value {
						answers = append(answers, true)
						break
					}
				}
				answers = append(answers, false)
			case "xref":
				for _, xref := range term.Xrefs {
					if xref == value {
						answers = append(answers, true)
						break
					}
				}
				answers = append(answers, false)
			case "is_a":
				answers = append(answers, term.IsA != value)
			case "is_obsolete":
				answers = append(answers, term.Obsolete != value)
			default:
				answers = append(answers, true)
				fmt.Println("Warning: Querying obo file, case not handled<" + key + ">")
			}
		}
	}

	// Depending on the query type ('or' or 'and'), loop through the answers
	switch queryType {
	case "or": // At least one true
		for _, answer := range answers {
			if answer {
				return true
			}
		}
		return false
	case "and": // All must be true
		for _, answer := range answers {
			if !answer {
				return false
			}
		}
		return true
	default: // No type, return first answer
		return answers[0]
	}
}

// Return the next term
func nextTerm() (*HPOOBOStruct, error) {

	term := new(HPOOBOStruct)

	// Continue the file parsing from the last position
	for scanner.Scan() {
		switch scanner.Text() {
		// Start of a Term
		case "[Term]":
			continue
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
			case "replaced_by":
				return nil, nil
			case "created_by", "creation_date", "subset", "is_anonymous", "consider":
			default:
				fmt.Println("Warning: Unexpected field <", lineParts[0], "> during obo file parsing")
			}
		}
	}

	return term, scanner.Err()
}
