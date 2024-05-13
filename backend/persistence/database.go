package persistence

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Database struct {
	Path      string
	TablesMap map[string]*Table
}

var DB *Database

func updateId(tableName string) error {
	systemTable := DB.TablesMap["system"]
	table := DB.TablesMap[tableName]
	tableRow, err := systemTable.Find(table.id)
	if err != nil {
		return errors.New("[Error]: table not found")
	}
	table.nextId += 1
	nextId := fmt.Sprint(table.nextId)
	tableRow["nextId"] = nextId
	tableMapSlice, _ := systemTable.GetAll()
	var updatedMapSlice []map[string]string
	for index, rowMap := range tableMapSlice {
		if rowMap["id"] == table.id {
			updatedMapSlice = append(updatedMapSlice, tableMapSlice[:index]...)
			updatedMapSlice = append(updatedMapSlice, tableRow)
			updatedMapSlice = append(updatedMapSlice, tableMapSlice[index+1:]...)
			break
		}
	}
	if updatedMapSlice != nil {
		csvSlice := mapAllToCsv(systemTable.name, updatedMapSlice)
		systemTable.csvHandler.writeAll(csvSlice)
		return nil
	} else {
		return errors.New("[Error]: ID not found")
	}
}

func NewDatabase(path string) (*Database, error) {
	if path == "" {
		return nil, errors.New("[Error]: Database path must be set")
	}
	tableName := "system"
	csvPath := fmt.Sprintf("%s/%s.csv", path, tableName)
	csvHandler, _ := newCsvHandler(csvPath)

	tableKeys := []string{"id", "path", "keys", "nextId"}
	tableRow := []string{"0", csvPath, strings.Join(tableKeys, " "), "1"}
	var csvSlice = make([][]string, 2)
	csvSlice[0] = tableKeys
	csvSlice[1] = tableRow
	systemTable := &Table{
		id:         "0",
		name:       tableName,
		keys:       tableKeys,
		nextId:     1,
		csvHandler: csvHandler,
	}
	tablesMap := make(map[string]*Table)
	tablesMap[tableName] = systemTable
	DB = &Database{
		Path:      path,
		TablesMap: tablesMap,
	}
	if _, err := os.Stat(csvPath); os.IsNotExist(err) {
		err := systemTable.csvHandler.writeAll(csvSlice)
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

	csvHandler := &csvHandler{
		csvFile: csvPath,
	}

	returnTable := &Table{
		id:         fmt.Sprint(systemTable.nextId),
		name:       tableName,
		keys:       tableKeys,
		nextId:     1,
		csvHandler: csvHandler,
	}
	database.TablesMap[tableName] = returnTable
	if _, err := os.Stat(csvPath); os.IsNotExist(err) {
		err := returnTable.csvHandler.write(tableKeys)
		if err != nil {
			return nil, err
		}
		// add new table to system table
		tableRow := []string{fmt.Sprint(systemTable.nextId), tablePath, strings.Join(tableKeys, " "), "1"}
		err = systemTable.csvHandler.write(tableRow)
		if err != nil {
			return nil, err
		} // update database nextId state
		err = updateId("system")
		if err != nil {
			return nil, err
		}
	} else {
		tableRow, err := systemTable.Find(returnTable.id)
		if err != nil {
			return nil, err
		}
		returnTable.nextId, _ = strconv.Atoi(tableRow["nextId"])
	}
	return returnTable, err
}
