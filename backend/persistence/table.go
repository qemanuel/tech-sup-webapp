package persistence

import (
	"encoding/csv"
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
	f, err := os.OpenFile(table.csvPath, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	w := csv.NewWriter(f)
	w.Write(row)
	w.Flush()
	f.Close()
	return nil
}

func (table *Table) updateId() error {
	table.nextId += 1
	nextId := []string{fmt.Sprint(table.nextId)}
	nextIdRow := [][]string{nextId}
	err := table.overWrite(table.nextIdPath, nextIdRow)
	return err
}

func (table *Table) overWrite(csvPath string, row [][]string) error {
	f, err := os.OpenFile(csvPath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	// Reset the file size
	f.Truncate(0)
	// Position on the beginning of the file:
	f.Seek(0, 0)
	w := csv.NewWriter(f)
	w.WriteAll(row)
	w.Flush()
	f.Close()
	return nil
}

func (table *Table) ReadAll() [][]string {

	csvfile, err := os.Open(table.csvPath)
	if err != nil {
		return nil
	}
	//defer csvfile.Close()
	csvfile.Seek(0, 0)
	reader := csv.NewReader(csvfile)

	rawCSVdata, err := reader.ReadAll()
	if err != nil {
		return nil
	}
	csvfile.Close()
	return rawCSVdata
}

func (table *Table) ReadRow(id int) []string {
	idToFind := fmt.Sprint(id)
	csvfile, err := os.Open(table.csvPath)
	if err != nil {
		return nil
	}
	//defer csvfile.Close()
	csvfile.Seek(0, 0)
	reader := csv.NewReader(csvfile)
	rawCSVdata, err := reader.ReadAll()
	csvfile.Close()
	found := []string{}
	if err != nil {
		return nil
	}
	for _, line := range rawCSVdata {
		if line[0] == idToFind {
			found = line
			break
		}
	}
	return found
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
	nextIdFile, err := os.Open(nextIdPath)
	if err != nil {
		return nil, err
	}
	//defer nextIdFile.Close()
	nextIdFile.Seek(0, 0)
	rNextIdFile := csv.NewReader(nextIdFile)
	nextIdRaw, _ := rNextIdFile.ReadAll()
	nextIdFile.Close()
	nextId, _ := strconv.Atoi(nextIdRaw[0][0])
	// read table file keys
	tableFile, err := os.Open(csvPath)
	if err != nil {
		return nil, err
	}
	//defer tableFile.Close()
	tableFile.Seek(0, 0)
	rTableFile := csv.NewReader(tableFile)
	tableRaw, _ := rTableFile.ReadAll()
	tableFile.Close()
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

func (table *Table) AddRow(row []string) (int, error) {
	fullRow := []string{fmt.Sprint(table.nextId)}
	fullRow = append(fullRow, row...)
	table.write(fullRow)
	table.updateId()
	return table.nextId, nil
}

func (table *Table) RemoveRow(id int) error {
	idToFind := fmt.Sprint(id)
	csvfile, err := os.Open(table.csvPath)
	if err != nil {
		return nil
	}
	//defer csvfile.Close()
	csvfile.Seek(0, 0)
	reader := csv.NewReader(csvfile)
	rawCSVdata, err := reader.ReadAll()
	if err != nil {
		return nil
	}
	csvfile.Close()
	var updatedCSVData [][]string
	for index, line := range rawCSVdata {
		if line[0] == idToFind {
			updatedCSVData = append(rawCSVdata[:index], rawCSVdata[index+1:]...)
			break
		}
	}
	table.overWrite(table.csvPath, updatedCSVData)
	return nil
}
