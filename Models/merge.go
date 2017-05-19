package Models

import "fmt"

// Merge - Merge two arrays of Diseases
// Support "or" and "and" logic operator
func Merge(list1 []*Disease, list2 []*Disease, operator string) []*Disease {
	switch operator {
	case "and":
		return andMerge(list1, list2)
	case "or":
		return orMerge(list1, list2)
	default:
		fmt.Println("Operator <" + operator + "> not suported for merging")
		return nil
	}
}

// Merge two Disease list with the 'and' operator
func andMerge(list1 []*Disease, list2 []*Disease) []*Disease {
	var mergeResult []*Disease
	// Put all item of list1 contained in list2 in mergeResult
	for _, item1 := range list1 {
		for _, item2 := range list2 {
			if item1.Name == item2.Name {
				mergeResult = append(mergeResult, item1)
			}
		}
	}

	// Put all item of list2 contained in list1 in mergeResult
	// Check that the item is not allready in mergeResult
	for _, item2 := range list2 {
		for _, item1 := range list1 {
			if item1.Name == item2.Name {
				contains := false
				for _, itemR := range mergeResult {
					if itemR.Name == item2.Name {
						contains = true
						break
					}
				}
				if !contains {
					mergeResult = append(mergeResult, item2)
				}
			}
		}
	}

	return mergeResult
}

// Merge two Disease list with the 'or' operator
func orMerge(list1 []*Disease, list2 []*Disease) []*Disease {
	var mergeResult []*Disease
	// Put all item of list1 in mergeResult
	for _, item1 := range list1 {
		mergeResult = append(mergeResult, item1)
	}
	// Put all item if list2 in mergeResult
	// Check that the item is not allready in mergeResult
	for _, item2 := range list2 {
		contains := false
		for _, itemR := range mergeResult {
			if itemR.Name == item2.Name {
				contains = true
				break
			}
		}
		// Add the item to mergeResult if mergeResult does not contains it
		if !contains {
			mergeResult = append(mergeResult, item2)
		}
	}

	return mergeResult
}
