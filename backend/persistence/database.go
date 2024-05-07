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
	err := databaseTable.write(tablesKeys)
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
	// write table keys definition
	tableContent := [][]string{tableKeys}
	err = database.table.overWrite(csvPath, tableContent)
	if err != nil {
		return nil, err
	}
	// write table nextId state
	nextIdRow := []string{"1"}
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

func LoadTable(path string, tableName string) (*Table, error) {
	// find files
	var csvPath string
	var nextIdPath string
	if path == "" {
		csvPath = fmt.Sprintf("./%s.csv", tableName)
		nextIdPath = fmt.Sprintf("./%s.nextId", tableName)
	} else {
		csvPath = fmt.Sprintf("%s/%s.csv", path, tableName)
		nextIdPath = fmt.Sprintf("%s/%s.nextId", path, tableName)
	}
	// read nextId file value
	nextIdRaw, _ := readAll(nextIdPath)
	nextId, _ := strconv.Atoi(nextIdRaw[0][0])
	// read table file keys
	tableRaw, _ := readAll(csvPath)
	tableKeys := tableRaw[0]
	return &Table{
		id:         0,
		keys:       tableKeys,
		nextId:     nextId,
		path:       path,
		csvPath:    csvPath,
		nextIdPath: nextIdPath,
	}, nil
}

func readAll(csvPath string) ([][]string, error) {

	csvfile, err := os.Open(csvPath)
	if err != nil {
		return nil, err
	}
	csvfile.Seek(0, 0)
	reader := csv.NewReader(csvfile)

	rawCSVdata, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	csvfile.Close()
	return rawCSVdata, nil
}
