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
	id         int
	keys       []string
	nextId     int
	path       string
	csvPath    string
	nextIdPath string
}

func NewDatabase(path string) (*Database, error) {
	var csvPath string
	var nextIdPath string
	if path == "" {
		csvPath = "./database.csv"
		nextIdPath = "./nextId"
	} else {
		csvPath = fmt.Sprintf("%s/database.csv", path)
		nextIdPath = fmt.Sprintf("%s/nextId", path)
	}
	tablesKeys := []string{"id", "csvPath", "keys", "nextId"}
	var err error
	err = write(csvPath, tablesKeys)
	write(nextIdPath, []string{"1"})
	if err != nil {
		return nil, err
	} else {
		return &Database{
			id:         0,
			keys:       tablesKeys,
			nextId:     1,
			path:       path,
			csvPath:    csvPath,
			nextIdPath: nextIdPath,
		}, nil
	}
}

func LoadDatabase(path string) (*Database, error) {
	var csvPath string
	var nextIdPath string
	if path == "" {
		csvPath = "./database.csv"
		nextIdPath = "./nextId"
	} else {
		csvPath = fmt.Sprintf("%s/database.csv", path)
		nextIdPath = fmt.Sprintf("%s/nextId", path)
	}
	databaseNextId, _ := strconv.Atoi(csvFileToSlice(nextIdPath)[0][0])
	databaseKeys := csvFileToSlice(csvPath)[0]
	return &Database{
		id:         0,
		keys:       databaseKeys,
		nextId:     databaseNextId,
		path:       path,
		csvPath:    csvPath,
		nextIdPath: nextIdPath,
	}, nil
}

func (database *Database) update(table Table) error {
	row := []string{fmt.Sprint(table.id), table.csvPath, strings.Join(table.keys, " "), fmt.Sprint(table.nextId)}
	err := write(database.csvPath, row)
	if err != nil {
		return err
	}
	f, err := os.OpenFile(database.nextIdPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	// Reset the file size
	f.Truncate(0)
	// Position on the beginning of the file:
	f.Seek(0, 0)
	w := csv.NewWriter(f)
	w.Write([]string{fmt.Sprint(database.nextId)})
	w.Flush()
	f.Close()
	return nil
}

func (database *Database) NewTable(tableName string, keys []string) (*Table, error) {
	if len(keys) == 0 || tableName == "" {
		return nil, errors.New("[Error]: keys or csvPath missing")
	}
	csvPath := fmt.Sprintf("%s/%s.csv", database.path, tableName)
	err := write(csvPath, keys)
	if err != nil {
		return nil, err
	}
	returnTable := Table{
		id:      database.nextId,
		keys:    keys,
		nextId:  1,
		csvPath: csvPath,
	}
	database.nextId += 1
	updateErr := database.update(returnTable)
	return &returnTable, updateErr
}
