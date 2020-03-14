package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/weibelben/finance_parser/pkg/citibank"
	"github.com/weibelben/finance_parser/pkg/transaction"
)

// parseCitiStatements returns an []Statement type containing all of the
// citibank statement data present in /statements/citibank/
func parseCitiStatements() (transaction.StatementType, error) {
	log.Info("Your citibank data is being parsed...")
	statements, err  := citibank.ParseStatements()
	if err != nil {
		return statements, err
	}

	return statements, nil
}