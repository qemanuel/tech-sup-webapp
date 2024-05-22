package persistence

import (
	"errors"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

type Database struct {
	Path      string
	TablesMap map[string]*Table
}

type Record struct {
	Id        string    `mapstructure:"id" json:"id"`
	CreatedAt time.Time `mapstructure:"created_at" json:"created_at"`
	UpdatedAt time.Time `mapstructure:"updated_at" json:"updated_at"`
}

var DB *Database

func updateId(tableName string) error {
	systemTable := DB.TablesMap["system"]
	table := DB.TablesMap[tableName]
	tableRow, err := systemTable.Find(table.Id)
	if err != nil {
		return errors.New("[Error]: table not found")
	}
	// update next id
	table.nextId += 1
	nextId := fmt.Sprint(table.nextId)
	tableRow["next_id"] = nextId
	// update timeStamp
	timeStamp := time.Now()
	table.UpdatedAt = timeStamp
	tableRow["updated_at"] = timeStamp.Format(time.DateTime)
	// write updated values to system table
	tableMapSlice, _ := systemTable.GetAll()
	var updatedMapSlice []map[string]string
	for index, rowMap := range tableMapSlice {
		if rowMap["id"] == table.Id {
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
	timeStamp := time.Now()
	tableKeys := []string{"id", "created_at", "updated_at", "path", "keys", "next_id"}
	tableRow := []string{"0", timeStamp.Format(time.DateTime), "", csvPath, strings.Join(tableKeys, " "), "1"}
	var csvSlice = make([][]string, 2)
	csvSlice[0] = tableKeys
	csvSlice[1] = tableRow
	systemTable := &Table{
		Record: Record{
			Id:        "0",
			CreatedAt: timeStamp,
		},
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

func (database *Database) NewTable(tableName string, keys []string) (*Table, error) {
	// validate inputs
	if len(keys) == 0 || tableName == "" || tableName == "system" {
		return nil, errors.New("[Error]: Table name or keys invalid")
	}
	// validate table keys
	tableKeys := []string{"id", "created_at", "updated_at"}
	for _, key := range keys {
		if !slices.Contains(tableKeys, key) {
			tableKeys = append(tableKeys, key)
		}
	}

	systemTable := database.TablesMap["system"]
	// define table files path
	var err error
	tablePath := fmt.Sprintf("%s/%s", database.Path, tableName)
	csvPath := fmt.Sprintf("%s.csv", tablePath)
	// init structs
	timeStamp := time.Now()
	csvHandler := &csvHandler{
		csvFile: csvPath,
	}

	returnTable := &Table{
		Record: Record{
			Id:        fmt.Sprint(systemTable.nextId),
			CreatedAt: timeStamp,
		},
		name:       tableName,
		keys:       tableKeys,
		nextId:     1,
		csvHandler: csvHandler,
	}
	database.TablesMap[tableName] = returnTable
	// if csv file doesn't exist, creates it
	if _, err := os.Stat(csvPath); os.IsNotExist(err) {
		err := returnTable.csvHandler.write(tableKeys)
		if err != nil {
			return nil, err
		}
		// add new table to system table
		tableRow := []string{fmt.Sprint(systemTable.nextId), timeStamp.Format(time.DateTime), "", tablePath, strings.Join(tableKeys, " "), "1"}
		err = systemTable.csvHandler.write(tableRow)
		if err != nil {
			return nil, err
		} // update database nextId state
		err = updateId("system")
		if err != nil {
			return nil, err
		}
	} else {
		tableRow, err := systemTable.Find(returnTable.Id)
		if err != nil {
			return nil, err
		}
		returnTable.nextId, _ = strconv.Atoi(tableRow["next_id"])
	}
	return returnTable, err
}
