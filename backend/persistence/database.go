package persistence

import (
	"errors"
	"fmt"
	"strings"
)

type Database struct {
	table *Table
}

func NewDatabase(path string) (*Database, error) {
	var csvPath string
	var nextIdPath string
	if path == "" {
		csvPath = "./database.csv"
		nextIdPath = "./database.nextId"
	} else {
		csvPath = fmt.Sprintf("%s/database.csv", path)
		nextIdPath = fmt.Sprintf("%s/database.nextId", path)
	}
	tablesKeys := []string{"id", "path", "keys"}
	databaseTable := &Table{
		id:         0,
		keys:       tablesKeys,
		nextId:     1,
		path:       path,
		csvPath:    csvPath,
		nextIdPath: nextIdPath,
	}
	var err error
	err = databaseTable.write(tablesKeys)
	if err != nil {
		return nil, err
	}
	return &Database{
		table: databaseTable,
	}, nil
}

func LoadDatabase(path string) (*Database, error) {
	tableName := "database"
	databaseTable, err := LoadTable(path, tableName)
	if err != nil {
		return nil, err
	}
	return &Database{
		table: databaseTable,
	}, nil
}

func (database *Database) NewTable(tableName string, tableKeys []string) (*Table, error) {
	// validate inputs
	if len(tableKeys) == 0 || tableName == "" {
		return nil, errors.New("[Error]: keys or csvPath missing")
	}
	// define table files path
	var err error
	tablePath := fmt.Sprintf("%s/%s", database.table.path, tableName)
	csvPath := fmt.Sprintf("%s.csv", tablePath)
	nextIdPath := fmt.Sprintf("%s.nextId", tablePath)
	fmt.Println(tablePath, csvPath, nextIdPath)
	// write table keys definition
	tableContent := [][]string{tableKeys}
	err = database.table.overWrite(csvPath, tableContent)
	if err != nil {
		return nil, err
	}
	// write table nextId state
	nextIdRow := []string{fmt.Sprint(database.table.nextId)}
	nextIdContent := [][]string{nextIdRow}
	err = database.table.overWrite(nextIdPath, nextIdContent)
	if err != nil {
		return nil, err
	}
	returnTable := Table{
		id:         database.table.nextId,
		keys:       tableKeys,
		nextId:     1,
		path:       tablePath,
		csvPath:    csvPath,
		nextIdPath: nextIdPath,
	}
	// Update database table with the added table definition
	newTableRow := []string{fmt.Sprint(database.table.nextId), tablePath, strings.Join(tableKeys, " ")}
	err = database.table.write(newTableRow)
	if err != nil {
		return nil, err
	}
	// update database nextId state
	err = database.table.updateId()
	if err != nil {
		return nil, err
	}
	return &returnTable, err
}
