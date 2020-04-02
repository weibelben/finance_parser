package citibank

import (
	"fmt"

	"github.com/weibelben/finance_parser/internal/csvReader"
	"github.com/weibelben/finance_parser/pkg/provider"
	"github.com/weibelben/finance_parser/pkg/transaction"
	log "github.com/sirupsen/logrus"
)

// ParseStatementFiles returns the parsed data of all citi statements
func ParseStatementFiles() ([]transaction.StatementType, error) {
	statementFiles, err := provider.FindStatements("citibank_statements")
	if err != nil {
		return nil, err
	}

	return parseStatements(statementFiles)
}

// parseStatements returns a slice of statements
func parseStatements(statementFiles []string) ([]transaction.StatementType, error) {
	combinedStatementData := []transaction.StatementType{}

	// read and parse each statement
	for _, file := range statementFiles {
		rawStatementData, err := csvReader.ReadCSV(file)
		if err != nil {
			return nil, err
		}

		parsedStatementData, err := parseRawStatementData(rawStatementData)
		if err != nil {
			return nil, err
		}

		combinedStatementData = append(combinedStatementData, parsedStatementData)
	}

	return combinedStatementData, nil
}

// parseRawStatementData is inherited from the provider interface.
func parseRawStatementData(rawStatementData [][]string) (transaction.StatementType, error) {
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

	// Ensure all charges are debit charges
	amountStr := row[3]
	if row[4] != "" {
		if row[3] != "" {
			return record, &provider.StatementSyntaxError{
				Message: fmt.Sprintf(
					"both debit and credit amounts present: %s, %s",
					row[3],
					row[4],
				),
			}
		} 
		amountStr = row[4]
 	}

	return transaction.Record(row[1], amountStr, row[2])
}