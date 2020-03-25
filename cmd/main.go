package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/weibelben/finance_parser/pkg/citibank"
	"github.com/weibelben/finance_parser/pkg/transaction"
)

func main() {
	_, err := parseStatements()
	if err != nil {
		log.WithError(err).Error("failed to parse statements")
	}
}

// parseStatments is the default, parse all function
func parseStatements() ([][]transaction.StatementType, error) {
	var aggregatedStatements [][]transaction.StatementType

	citiStatements, err := citibank.ParseStatementFiles()
	if err != nil {
		log.Error("unable to parse Citibank statements")
		return aggregatedStatements, err
	}

	aggregatedStatements = append(aggregatedStatements, citiStatements)

	return aggregatedStatements, nil
}
