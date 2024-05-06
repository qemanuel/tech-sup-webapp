package persistence

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Table struct {
	id      int
	keys    []string
	nextId  int
	csvPath string
}

func (table *Table) GetAllElements() [][]string {
	return csvFileToSlice(table.csvPath)
}

func (table *Table) GetElement(id int) []string {
	sliceTable := csvFileToSlice(table.csvPath)
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
	sliceTable := csvFileToSlice(table.csvPath)
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
