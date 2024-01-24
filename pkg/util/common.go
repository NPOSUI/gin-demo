package util

func ListDiff(firstList []string, secondList []string) []string {
	var diff []string
	for _, item := range firstList {
		if !IsInList(item, secondList) {
			diff = append(diff, item)
		}
	}
	return diff
}

func IsInList(element string, list []string) bool {
	for _, item := range list {
		if item == element {
			return true
		}
	}
	return false
}
