package citibank

import (
	"github.com/weibelben/finance_parser/pkg/provider"
	"github.com/weibelben/finance_parser/pkg/transaction"
	log "github.com/sirupsen/logrus"
)

// citibank package implements provider interface

// ParseStatementFiles returns the parsed data of all citi statements
func ParseStatementFiles() ([]transaction.StatementType, error) {
	var citiProvider provider.Provider
	statementFiles, err := provider.FindStatements("citibank_statements")
	if err != nil {
		return nil, err
	}

	return provider.ParseStatements(citiProvider, statementFiles)
}

// ParseRawStatementData is inherited from the provider interface.
func ParseRawStatementData(rawStatementData [][]string) (transaction.StatementType, error) {
	log.Info("parsing statement")
	var statementData transaction.StatementType
	statementData.Provider = "citibank"
	
	var records []transaction.RecordType
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

		records = append(records, record)
	}

	statementData.Records = records

	return statementData, nil
}

// parseStatementEntry is inherited from the provider interface.
// Citibank statements are of the form:
// status, date, description, debit amount, credit amount
func parseStatementEntry(row []string) (transaction.RecordType, error) {
	var record transaction.RecordType
	if len(row) != 5 {
		return record, &provider.StatementSyntaxError{
			Message: "unexpected length of citibank statement entry",
		}
	}

	if row[4] != "" {
		if row[3] == "" {
			return record, &provider.StatementSyntaxError{
				Message: "credit charge instead of debit",
			}
		}
 	}

	return transaction.Record(row[1], row[3], row[2])
}