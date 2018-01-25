package db_doc_gen

func InArray(strArr []string, search_item string) bool {
	for _, item := range strArr {
		if item == search_item {
			return true
		}
	}

	return true
}
