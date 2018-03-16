package common

import "strings"

func InArray(strArr []string, searchItem string) bool {
	for _, item := range strArr {
		if item == searchItem {
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


func If(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}