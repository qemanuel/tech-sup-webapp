package persistence

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Database struct {
	Path       string
	NextIdFile string
	TablesMap  map[string]*Table
}

var DB *Database

func NewDatabase(path string) (*Database, error) {
	if path == "" {
		return nil, errors.New("[Error]: Database path must be set")
	}
	csvPath := fmt.Sprintf("%s/system.csv", path)
	nextIdPath := fmt.Sprintf("%s/nextId", path)
	systemKeys := []string{"id", "path", "keys", "nextId"}
	systemTable := &Table{
		id:      "0",
		keys:    systemKeys,
		nextId:  1,
		path:    path,
		csvPath: csvPath,
	}
	systemMap := make(map[string]string, len(systemKeys))
	for _, key := range systemKeys {
		systemMap[key] = key
	}
	tablesMap := make(map[string]*Table)
	tablesMap["system"] = systemTable
	DB = &Database{
		Path:       path,
		NextIdFile: nextIdPath,
		TablesMap:  tablesMap,
	}
	if _, err := os.Stat(csvPath); os.IsNotExist(err) {
		err := write("system", systemMap)
		if err != nil {
			return nil, err
		}
	}
	return DB, nil
}

func (database *Database) NewTable(tableName string, tableKeys []string) (*Table, error) {
	// validate inputs
	if len(tableKeys) == 0 || tableName == "" || tableName == "system" {
		return nil, errors.New("[Error]: Table name or keys invalid")
	}
	systemTable := database.TablesMap["system"]
	// define table files path
	var err error
	tablePath := fmt.Sprintf("%s/%s", database.Path, tableName)
	csvPath := fmt.Sprintf("%s.csv", tablePath)

	returnTable := &Table{
		id:      fmt.Sprint(systemTable.nextId),
		keys:    tableKeys,
		nextId:  1,
		path:    tablePath,
		csvPath: csvPath,
	}
	database.TablesMap[tableName] = returnTable
	if _, err := os.Stat(csvPath); os.IsNotExist(err) {
		f, err := os.OpenFile(csvPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		// Position on the end of the file
		f.Seek(0, 2)
		w := csv.NewWriter(f)
		w.Write(tableKeys)
		w.Flush()
		f.Close()
		if err != nil {
			return nil, err
		}

		// add new table to system table
		tableRow := []string{fmt.Sprint(systemTable.nextId), tablePath, strings.Join(tableKeys, " "), "1"}
		f, err = os.OpenFile(systemTable.csvPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		// Position on the end of the file
		f.Seek(0, 2)
		w = csv.NewWriter(f)
		w.Write(tableRow)
		w.Flush()
		f.Close()
		if err != nil {
			return nil, err
		} // update database nextId state
		err = updateId("system")
		if err != nil {
			return nil, err
		}
	}
	return returnTable, err
}
