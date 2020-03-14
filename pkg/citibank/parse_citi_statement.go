package citibank

import (
	"io"
	"io/ioutil"
	"encoding/csv"
	"os"
	"strings"

	"github.com/weibelben/finance_parser/pkg/transaction"
	"github.com/weibelben/finance_parser/internal/csvReader"
	log "github.com/sirupsen/logrus"
)

// citibank package implements provider interface

// ParseStatements returns the parsed data of all citi statements
func ParseStatements() (transaction.StatementType, error) {
	statementFiles, err := findStatements()
	if err != nil {
		log.WithError(err).Fatal("Failed to collect citibank statements.")
		return statementFiles, err
	}

	for _, file := range statementFiles {
		statementData, err := csvReader.ReadCSV(file)
	}

	return statementFiles, nil
}

// findStatements returns the names of all the csv files in the
// citibank/ dir
func findStatements() ([]string, error) {
	root, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	citiDir := root + "statements/citibank_statements"

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