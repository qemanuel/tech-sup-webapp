package main

import (
	"fmt"
	"os"

	"github.com/qemanuel/tech-sup-webapp/backend/persistence"
)

func initDatabase() (*persistence.Database, map[string]*persistence.Table) {
	databasePath := "/Users/qemanuel/IT/personal/github/tech-sup-webapp/backend/main/database"
	tables := map[string][]string{
		"workers":   {"id", "name", "email", "phone"},
		"customers": {"id", "name", "email", "phone"},
		"devices":   {"id", "owner", "kind", "brand", "model", "serial"},
		//		"jobs":       {""},
		//		"incidences": {""},
	}
	tablesMap := map[string]*persistence.Table{}
	_, err := os.Stat(fmt.Sprintf("%s/database.csv", databasePath))
	if err != nil {
		newDatabase, _ := persistence.NewDatabase(databasePath)
		for tableName, keys := range tables {
			table, _ := newDatabase.NewTable(tableName, keys)
			tablesMap[tableName] = table
		}
		return newDatabase, tablesMap
	} else {
		loadDatabase, _ := persistence.LoadDatabase(databasePath)
		for tableName, _ := range tables {
			table, _ := persistence.LoadTable(databasePath, tableName)
			tablesMap[tableName] = table
		}
		return loadDatabase, tablesMap
	}
}

func main() {
	database, tables := initDatabase()
	fmt.Println(database, tables)
}
