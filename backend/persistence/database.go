package persistence

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Database struct {
	Path       string
	NextIdFile string
	TablesMap  map[string]*Table
}

var DB *Database

func updateId(tableName string) error {
	systemTable := DB.TablesMap["system"]
	if tableName == "system" {
		systemTable.nextId += 1
		nextIdFile := DB.NextIdFile
		rowSlice := []string{fmt.Sprint(systemTable.nextId)}
		csvSlice := [][]string{rowSlice}
		f, err := os.OpenFile(nextIdFile, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		// Reset the file size
		f.Truncate(0)
		// Position on the beginning of the file
		f.Seek(0, 0)
		w := csv.NewWriter(f)
		w.WriteAll(csvSlice)
		w.Flush()
		f.Close()
		return nil

	} else {
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
			systemTable.csvHandler.writeAll(systemTable.name, updatedMapSlice)
			return nil
		} else {
			return errors.New("[Error]: ID not found")
		}
	}
}

func NewDatabase(path string) (*Database, error) {
	if path == "" {
		return nil, errors.New("[Error]: Database path must be set")
	}
	tableName := "system"
	csvPath := fmt.Sprintf("%s/%s.csv", path, tableName)
	nextIdPath := fmt.Sprintf("%s/nextId", path)

	csvHandler := &csvHandler{
		csvFile: csvPath,
	}

	systemKeys := []string{"id", "path", "keys", "nextId"}
	systemTable := &Table{
		id:         "0",
		name:       tableName,
		keys:       systemKeys,
		nextId:     1,
		csvHandler: csvHandler,
	}
	systemMap := make(map[string]string, len(systemKeys))
	for _, key := range systemKeys {
		systemMap[key] = key
	}
	tablesMap := make(map[string]*Table)
	tablesMap[tableName] = systemTable
	DB = &Database{
		Path:       path,
		NextIdFile: nextIdPath,
		TablesMap:  tablesMap,
	}
	if _, err := os.Stat(csvPath); os.IsNotExist(err) {
		err := systemTable.csvHandler.write(tableName, systemMap)
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
	systemCsv := systemTable.csvHandler.csvFile
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
		f, err = os.OpenFile(systemCsv, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
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
	} else {
		tableRow, err := systemTable.Find(returnTable.id)
		if err != nil {
			return nil, err
		}
		returnTable.nextId, _ = strconv.Atoi(tableRow["nextId"])
	}
	return returnTable, err
}
