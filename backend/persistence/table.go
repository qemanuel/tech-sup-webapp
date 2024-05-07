package persistence

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"
)

type Table struct {
	id         int
	keys       []string
	nextId     int
	path       string
	csvPath    string
	nextIdPath string
}

func (table *Table) write(row []string) error {
	f, err := os.OpenFile(table.csvPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	// Position on the end of the file
	f.Seek(0, 2)
	w := csv.NewWriter(f)
	w.Write(row)
	w.Flush()
	f.Close()
	return err
}

func (table *Table) updateId() error {
	table.nextId += 1
	nextId := []string{fmt.Sprint(table.nextId)}
	nextIdRow := [][]string{nextId}
	err := table.overWrite(table.nextIdPath, nextIdRow)
	return err
}

func (table *Table) overWrite(csvPath string, row [][]string) error {
	f, err := os.OpenFile(csvPath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	// Reset the file size
	f.Truncate(0)
	// Position on the beginning of the file
	f.Seek(0, 0)
	w := csv.NewWriter(f)
	w.WriteAll(row)
	w.Flush()
	f.Close()
	return nil
}

func readAll(csvPath string) ([][]string, error) {
	csvfile, err := os.Open(csvPath)
	if err != nil {
		return nil, err
	}
	// Position on the begining of the file
	csvfile.Seek(0, 0)
	reader := csv.NewReader(csvfile)
	rawCSVdata, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	csvfile.Close()
	return rawCSVdata, nil
}

func (table *Table) ReadRow(id int) []string {
	idToFind := fmt.Sprint(id)
	rawCSVdata, err := readAll(table.csvPath)
	if err != nil {
		return nil
	}
	found := []string{}
	for _, line := range rawCSVdata {
		if line[0] == idToFind {
			found = line
			break
		}
	}
	return found
}

func (table *Table) AddRow(row []string) (int, error) {
	fullRow := []string{fmt.Sprint(table.nextId)}
	fullRow = append(fullRow, row...)
	table.write(fullRow)
	returnID := table.nextId
	table.updateId()
	return returnID, nil
}

func (table *Table) RemoveRow(id int) error {
	idToFind := fmt.Sprint(id)
	rawCSVdata, err := readAll(table.csvPath)
	if err != nil {
		return nil
	}
	var updatedCSVData [][]string
	for index, line := range rawCSVdata {
		if line[0] == idToFind {
			updatedCSVData = append(rawCSVdata[:index], rawCSVdata[index+1:]...)
			break
		}
	}
	if updatedCSVData != nil {
		table.overWrite(table.csvPath, updatedCSVData)
		return nil
	} else {
		return errors.New("[Error]: ID not found")
	}
}

func LoadTable(path string, tableName string) (*Table, error) {
	if path == "" || tableName == "" {
		return nil, errors.New("[Error]: Database path and table name must be set")
	}
	csvPath := fmt.Sprintf("%s/%s.csv", path, tableName)
	nextIdPath := fmt.Sprintf("%s/%s.nextId", path, tableName)
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
