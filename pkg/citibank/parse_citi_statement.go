package citibank

import (
	"io/ioutil"
	"os"
	"strings"
)

// FindCitiStatements returns the names of all the csv files in the
// citibank/ dir
func FindCitiStatements() ([]string, error) {
	root, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	citiDir := root + "/citibank_statements"

	files, err := ioutil.ReadDir(citiDir)
	if err != nil {
		return nil, err
	}

	var statements []string
	for _, file := range files {
		// only include csv files
		if strings.HasSuffix(file.Name(), ".csv") {
			statements = append(statements, file.Name())
		}
	}
	
	return statements, nil
}