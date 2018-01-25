package db_doc_gen

import "log"

func (self *DbManager) filterTables() []string {
	allTables := self.getAllTables()

	if len(self.cfg.Includes) > 0 {
		return self.includeTables(allTables)
	}
	if len(self.cfg.Excludes) > 0 {
		return self.excludeTables(allTables)
	}

	return allTables
}

func (self *DbManager) includeTables(allTables []string) []string {
	var result []string
	for _, item := range self.cfg.Includes {
		if InArray(allTables, item) {
			result = append(result, item)
		} else {
			log.Fatalf("table %s isn't in database", item)
		}
	}
	return result
}

func (self *DbManager) excludeTables(allTables []string) []string {
	return ExcludeArray(allTables, self.cfg.Excludes...)
}

func (self *DbManager) getAllTables() []string {
	if (self.db == nil) {
		log.Println("db is nil")
		panic("")
	}

	db := self.db

	rows, err := db.Query("show tables")

	if (err != nil) {
		log.Println("error")
		panic(err)
	}
	defer rows.Close()

	var result []string
	for rows.Next() {
		var item string
		rows.Scan(&item)

		result = append(result, item)
	}

	return result
}
