package db_doc_gen

import (
	"fmt"
)

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
			fmt.Printf("Warning: %s isn't in database \n", item)
		}
	}
	return result
}

func (self *DbManager) excludeTables(allTables []string) []string {
	return ExcludeArray(allTables, self.cfg.Excludes...)
}

func (self *DbManager) getAllTables() []string {
	rows, err := self.db.Query("show tables")

	if (err != nil) {
		panic("Failed to show tables")
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
