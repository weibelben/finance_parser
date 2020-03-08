package citibank

import (
	"io"
	"io/ioutil"
	"encoding/csv"
	"os"
	"strings"

	"github.com/weibelben/finance_parser/pkg/transaction"
	log "github.com/sirupsen/logrus"
)

// ParseStatements returns the parsed data of all citi statements
func ParseStatements() (transaction.StatementType, error) {
	statementFiles, err := findStatements()
	if err != nil {
		log.WithError(err).Fatal("Failed to collect citibank statements.")
		return nil, err
	} 

	var statementData [][]string
	for _, file := range statementFiles {
		statementData, err = readCSV(file)
	}
	return nil, nil
}

func readCSV(filePath string) ([][]string, error){
	var r io.Reader
	r, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	csvReader := csv.NewReader(r)
	return csvReader.ReadAll()
}

// findStatements returns the names of all the csv files in the
// citibank/ dir
func findStatements() ([]string, error) {
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