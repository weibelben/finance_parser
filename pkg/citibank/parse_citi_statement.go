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
	}

	return parseStatements(statementFiles)
}

// ParseRawStatementData is inherited from the provider interface.
func ParseRawStatementData([][]string rawStatementData) (transaction.StatementType, error) {
	var statementData transaction.StatementType
	statementData.provider := "citibank"
	
	var records []statement.RecordType
	for i, row := range rawStatementData {
		// skip first row as it only contains headers
		if i == 0 {
			continue
		}

		// each row is a record
		record, err := parseStatementEntry(row)
		if err != nil {
			return statementData, err
		}

		records := append(records, record)
	}

	statementData.records := records

	return nil, nil
}

// parseStatementEntry is inherited from the provider interface.
// Citibank statements are of the form:
// date, description, amount
func parseStatementEntry([]string row) (transaction.RecordType, error) {
	var record transaction.RecordType
	if len(row) != 5 { // fixme idk how long an entry is
		return record, Errors.New("unexpected length of citibank statement entry") // FIXME create an error type for this
	}

	record.date := row[0]
	record.amount := row[2]
	record.description := row[1]

	return record
}