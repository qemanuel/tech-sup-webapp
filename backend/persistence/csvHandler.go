package persistence

import (
	"encoding/csv"
	"os"
)

type csvHandler struct {
	csvFile string
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

func (csvHandler *csvHandler) writeAll(tableName string, tableMapSlice []map[string]string) error {
	csvSlice := make([][]string, len(tableMapSlice)+1)
	for i, rowMap := range tableMapSlice {
		keysSlice, valuesSlice := mapToCsv(tableName, rowMap)
		if i == 0 {
			csvSlice[i] = keysSlice
		}
		csvSlice[i+1] = valuesSlice
	}
	f, err := os.OpenFile(csvHandler.csvFile, os.O_CREATE|os.O_WRONLY, 0644)
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

func (csvHandler *csvHandler) write(tableName string, modelMap map[string]string) error {
	_, row := mapToCsv(tableName, modelMap)
	f, err := os.OpenFile(csvHandler.csvFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	// Position on the end of the file
	f.Seek(0, 2)
	w := csv.NewWriter(f)
	w.Write(row)
	w.Flush()
	f.Close()
	return err
}

func (csvHandler *csvHandler) read() ([][]string, error) {
	csvfile, err := os.Open(csvHandler.csvFile)
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
	return csvSlice, nil
}
