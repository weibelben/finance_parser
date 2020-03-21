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
func ParseStatements() ([]transaction.StatementType, error) {
	statementFiles, err := findStatements("citibank_statements")
	if err != nil {
		return nil, err
		)
	}

	var combinedStatementData []transaction.StatementType
	statementData.provider := "citibank"

	for _, file := range statementFiles {
		statementData, err := csvReader.ReadCSV(file)
		if err != nil {
			return combinedStatementData, err
		}

		combinedStatementData = append(combinedStatementData, statementData)
	}

	parsedStatementData, err := parseRawData(combinedStatementData)
	if err != nil {
		return parsedStatementData, err
	}

	return combinedStatementData, nil
}