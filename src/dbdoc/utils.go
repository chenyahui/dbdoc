package dbdoc

import "strings"

func InArray(strArr []string, search_item string) bool {
	for _, item := range strArr {
		if item == search_item {
			return true
		}
	}

	return false
}

func ExcludeArray(strArr []string, elems ...string) []string {

	var result []string
	for _, item := range strArr {
		if !InArray(elems, item) {
			result = append(result, item)
		}
	}
	return result
}

func IsBlank(str string) bool {
	return len(strings.TrimSpace(str)) == 0
}
