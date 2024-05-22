package persistence

import (
	"errors"
	"fmt"
	"time"

	"github.com/mitchellh/mapstructure"
)

type Table struct {
	Record
	name       string
	keys       []string
	nextId     int `mapstructure:"next_id"`
	csvHandler *csvHandler
}

func mapAllToCsv(tableName string, tableMapSlice []map[string]string) [][]string {
	csvSlice := make([][]string, len(tableMapSlice)+1)
	for i, rowMap := range tableMapSlice {
		keysSlice, valuesSlice := mapToCsv(tableName, rowMap)
		if i == 0 {
			csvSlice[i] = keysSlice
		}
		csvSlice[i+1] = valuesSlice
	}
	return csvSlice
}

func mapToCsv(tableName string, rowMap map[string]string) (keysSlice []string, valuesSlice []string) {
	table := DB.TablesMap[tableName]
	keysSlice = table.keys
	returnSlice := make([]string, len(keysSlice))
	for i, key := range table.keys {
		returnSlice[i] = rowMap[key]
	}
	valuesSlice = returnSlice
	return keysSlice, valuesSlice
}

func csvToMap(keysSlice []string, valuesSlice []string) map[string]string {
	returnMap := make(map[string]string, len(keysSlice))
	for i, key := range keysSlice {
		returnMap[key] = valuesSlice[i]
	}
	return returnMap
}

func (table *Table) GetAll() ([]map[string]string, error) {
	csvSlice, err := table.csvHandler.read()
	if err != nil {
		return nil, err
	}
	tableMapSlice := make([]map[string]string, len(csvSlice)-1)
	keysSlice := csvSlice[0]
	rowsSlice := csvSlice[1:]
	for i, row := range rowsSlice {
		if len(row) != 0 {
			tableMapSlice[i] = csvToMap(keysSlice, row)
		}
	}
	return tableMapSlice, nil
}

func (table *Table) Find(id string) (map[string]string, error) {
	tableMapSlice, _ := table.GetAll()
	var tableFound map[string]string
	for _, rowMap := range tableMapSlice {
		if rowMap["id"] == id {
			tableFound = rowMap
			break
		}
	}
	if tableFound == nil {
		return nil, errors.New("404 Not Found")
	} else {
		return tableFound, nil
	}
}

func (table *Table) Add(model interface{}) (int, error) {
	var modelMap map[string]string
	mapstructure.Decode(model, &modelMap)
	timeStamp := time.Now()
	returnID := table.nextId
	modelMap["id"] = fmt.Sprint(returnID)
	modelMap["created_at"] = timeStamp.Format(time.DateTime)
	modelMap["updated_at"] = ""
	_, rowSlice := mapToCsv(table.name, modelMap)
	err := table.csvHandler.write(rowSlice)
	if err != nil {
		return 0, err
	}
	updateId(table.name)
	return returnID, nil
}

func (table *Table) Update(model interface{}, id string) error {
	var modelMap = make(map[string]string)
	mapstructure.Decode(model, &modelMap)
	modelMap["id"] = id
	modelMap["updated_at"] = time.Now().Format(time.DateTime)
	tableMapSlice, _ := table.GetAll()
	var updatedMapSlice []map[string]string
	for index, rowMap := range tableMapSlice {
		if rowMap["id"] == id {
			modelMap["created_at"] = rowMap["created_at"]
			updatedMapSlice = append(updatedMapSlice, tableMapSlice[:index]...)
			updatedMapSlice = append(updatedMapSlice, modelMap)
			updatedMapSlice = append(updatedMapSlice, tableMapSlice[index+1:]...)
			break
		}
	}
	if updatedMapSlice != nil {
		csvSlice := mapAllToCsv(table.name, updatedMapSlice)
		table.csvHandler.writeAll(csvSlice)
		return nil
	} else {
		return errors.New("[Error]: ID not found")
	}
}

func (table *Table) Remove(id string) error {
	tableMapSlice, err := table.GetAll()
	if err != nil {
		return nil
	}
	var updatedMapSlice []map[string]string
	for index, modelMap := range tableMapSlice {
		if modelMap["id"] == id {
			updatedMapSlice = append(tableMapSlice[:index], tableMapSlice[index+1:]...)
			break
		}
	}
	if updatedMapSlice != nil {
		csvSlice := mapAllToCsv(table.name, updatedMapSlice)
		table.csvHandler.writeAll(csvSlice)
		return nil
	} else {
		return errors.New("[Error]: ID not found")
	}
}
