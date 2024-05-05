package persistence

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
)

type Table struct {
	keys    []string
	nextId  int
	csvPath string
}

func NewTable(tableName string, keys []string) (*Table, error) {

	if len(keys) == 0 || tableName == "" {
		return nil, errors.New("[Error]: keys or csvPath missing")
	}
	csvPath := fmt.Sprintf("./%s.csv", tableName)
	csvFile, err := os.Create(csvPath)
	if err != nil {
		return nil, err
	}
	w := csv.NewWriter(csvFile)
	w.Write(keys)
	w.Flush()
	csvFile.Close()
	return &Table{
		keys:    keys,
		nextId:  1,
		csvPath: csvPath,
	}, nil
}

func (table *Table) GetAllElements() [][]string {
	return CSVFileToSlice(table.csvPath)
}

func (table *Table) GetElement(id int) []string {
	sliceTable := CSVFileToSlice(table.csvPath)
	var found []string
	for _, raw := range sliceTable {
		if raw[0] == fmt.Sprint(id) {
			found = raw
			break
		}
	}
	return found
}

func (table *Table) AddElement(row []string) (int, error) {
	f, err := os.OpenFile(table.csvPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return 0, err
	}
	var fullRow []string
	fullRow = append(fullRow, fmt.Sprint(table.nextId))
	fullRow = append(fullRow, row...)
	w := csv.NewWriter(f)
	w.Write(fullRow)
	w.Flush()
	f.Close()
	table.nextId += 1
	return table.nextId, nil
}

func (table *Table) RemoveElement(id int) error {
	sliceTable := CSVFileToSlice(table.csvPath)
	var newSlice [][]string
	for _, raw := range sliceTable {
		if raw[0] != fmt.Sprint(id) {
			newSlice = append(newSlice, raw)
		}
	}
	f, err := os.OpenFile(table.csvPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	// Reset the file size
	f.Truncate(0)
	// Position on the beginning of the file:
	f.Seek(0, 0)
	w := csv.NewWriter(f)
	w.WriteAll(newSlice)
	w.Flush()
	f.Close()
	return nil
}

func CSVFileToSlice(csvPath string) (returnMap [][]string) {

	csvfile, err := os.Open(csvPath)
	if err != nil {
		return nil
	}
	defer csvfile.Close()
	csvfile.Seek(0, 0)
	reader := csv.NewReader(csvfile)

	rawCSVdata, err := reader.ReadAll()
	if err != nil {
		return nil
	}
	return rawCSVdata
}
