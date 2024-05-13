package persistence

import (
	"encoding/csv"
	"errors"
	"os"
)

type csvHandler struct {
	csvFile string
}

func newCsvHandler(csvFile string) (*csvHandler, error) {
	if csvFile == "" {
		return nil, errors.New("[Error]: csvFile path is missing")
	} else {
		return &csvHandler{
			csvFile: csvFile,
		}, nil
	}
}

func (csvHandler *csvHandler) writeAll(tableSlice [][]string) error {
	f, err := os.OpenFile(csvHandler.csvFile, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	// Reset the file size
	f.Truncate(0)
	// Position on the beginning of the file
	f.Seek(0, 0)
	w := csv.NewWriter(f)
	w.WriteAll(tableSlice)
	w.Flush()
	f.Close()
	return err
}

func (csvHandler *csvHandler) write(rowSlice []string) error {
	f, err := os.OpenFile(csvHandler.csvFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	// Position on the end of the file
	f.Seek(0, 2)
	w := csv.NewWriter(f)
	w.Write(rowSlice)
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
