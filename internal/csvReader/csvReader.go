package csvReader

import (
	"io"
	"encoding/csv"
	"os"
)

// ReadCSV returns the data contained in a csv file as a [][]string
func ReadCSV(filePath string) ([][]string, error){
	var r io.Reader
	r, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	csvReader := csv.NewReader(r)
	return csvReader.ReadAll()
}