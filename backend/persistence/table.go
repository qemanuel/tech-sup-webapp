package persistence

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"

	"github.com/mitchellh/mapstructure"
)

type Table struct {
	id      string
	keys    []string
	nextId  int
	path    string
	csvPath string
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

func writeAll(tableName string, tableMapSlice []map[string]string) error {
	table := DB.TablesMap[tableName]
	csvSlice := make([][]string, len(tableMapSlice)+1)
	fmt.Println(len(tableMapSlice), len(csvSlice))
	for i, rowMap := range tableMapSlice {
		keysSlice, valuesSlice := mapToCsv(tableName, rowMap)
		if i == 0 {
			csvSlice[i] = keysSlice
		}
		csvSlice[i+1] = valuesSlice
	}
	fmt.Println(csvSlice)
	f, err := os.OpenFile(table.csvPath, os.O_CREATE|os.O_WRONLY, 0644)
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
	return err
}

func write(tableName string, modelMap map[string]string) error {
	table := DB.TablesMap[tableName]
	_, row := mapToCsv(tableName, modelMap)
	f, err := os.OpenFile(table.csvPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	// Position on the end of the file
	f.Seek(0, 2)
	w := csv.NewWriter(f)
	w.Write(row)
	w.Flush()
	f.Close()
	return err
}

func overWrite(tableName string, modelMap map[string]string, id string) error {
	tableMapSlice, _ := GetAll(tableName)
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
		writeAll(tableName, updatedMapSlice)
		return nil
	} else {
		return errors.New("[Error]: ID not found")
	}
}

func updateId(tableName string) error {
	var table *Table
	if tableName == "system" {
		table = DB.TablesMap["system"]
		table.nextId += 1
		nextIdFile := DB.NextIdFile
		rowSlice := []string{fmt.Sprint(table.nextId)}
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
		table = DB.TablesMap[tableName]
		tableRow := Find("system", table.id)
		if tableRow == nil {
			return errors.New("[Error]: table not found")
		}
		table.nextId += 1
		nextId := fmt.Sprint(table.nextId)
		tableRow["nextId"] = nextId
		overWrite("system", tableRow, table.id)
	}
	return nil
}

func GetAll(tableName string) ([]map[string]string, error) {
	table := DB.TablesMap[tableName]
	csvfile, err := os.Open(table.csvPath)
	if err != nil {
		return nil, err
	}
	// Position on the begining of the file
	csvfile.Seek(0, 0)
	reader := csv.NewReader(csvfile)
	csvSlice, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	csvfile.Close()
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

func Find(tableName string, id string) map[string]string {
	tableMapSlice, _ := GetAll(tableName)
	tableFound := make(map[string]string, len(tableMapSlice[0]))
	for _, rowMap := range tableMapSlice {
		if rowMap["id"] == id {
			tableFound = rowMap
			break
		}
	}
	return tableFound
}

func Add(tableName string, model interface{}) (int, error) {
	var modelMap map[string]string
	mapstructure.Decode(model, &modelMap)
	table := DB.TablesMap[tableName]
	returnID := table.nextId
	modelMap["id"] = fmt.Sprint(returnID)
	err := write(tableName, modelMap)
	if err != nil {
		return 0, err
	}
	updateId(tableName)
	return returnID, nil
}

func Update(tableName string, model interface{}, id string) error {
	var modelMap = make(map[string]string)
	mapstructure.Decode(model, modelMap)
	err := overWrite(tableName, modelMap, id)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func Remove(tableName string, id string) error {
	tableMapSlice, err := GetAll(tableName)
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
		writeAll(tableName, updatedMapSlice)
		return nil
	} else {
		return errors.New("[Error]: ID not found")
	}
}
