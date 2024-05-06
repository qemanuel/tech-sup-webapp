package persistence

import (
	"encoding/csv"
	"os"
)

func write(csvPath string, row []string) error {
	f, err := os.OpenFile(csvPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	w := csv.NewWriter(f)
	w.Write(row)
	w.Flush()
	f.Close()
	return nil
}

func csvFileToSlice(csvPath string) (returnMap [][]string) {

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
