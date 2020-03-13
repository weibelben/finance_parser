package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/weibelben/finance_parser/pkg/citibank"
	"github.com/weibelben/finance_parser/pkg/transaction"
)

// citibank package implements provider interface
func parseCitiStatements() transaction.StatementType {
	log.Info("Your citibank data is being compiled...")

	statement := transaction.Statement("citi", nil)
	return statement
}

func main() {
	log.Info("Your citibank data is being compiled...")
	statements, err  := citibank.FindStatements()
	if err != nil {
		log.WithError(err).Fatal("Failed to collect citibank statements.")
	}

	log.Info(statements)
}