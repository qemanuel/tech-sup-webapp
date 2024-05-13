package persistence

import (
	"errors"
	"fmt"

	"github.com/mitchellh/mapstructure"
)

type Table struct {
	id         string
	name       string
	keys       []string
	nextId     int
	csvHandler *csvHandler
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
	returnID := table.nextId
	modelMap["id"] = fmt.Sprint(returnID)
	err := table.csvHandler.write(table.name, modelMap)
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
	tableMapSlice, _ := table.GetAll()
	var updatedMapSlice []map[string]string
	for index, rowMap := range tableMapSlice {
		if rowMap["id"] == id {
			updatedMapSlice = append(updatedMapSlice, tableMapSlice[:index]...)
			updatedMapSlice = append(updatedMapSlice, modelMap)
			updatedMapSlice = append(updatedMapSlice, tableMapSlice[index+1:]...)
			break
		}
	}
	if updatedMapSlice != nil {
		table.csvHandler.writeAll(table.name, updatedMapSlice)
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
		table.csvHandler.writeAll(table.name, updatedMapSlice)
		return nil
	} else {
		return errors.New("[Error]: ID not found")
	}
}
