package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/weibelben/finance_parser/pkg/citibank"
)

func main() {
	log.Info("Your citibank data is being compiled...")
	statements, err  := citibank.FindCitiStatements()
	if err != nil {
		log.WithError(err).Fatal("Failed to collect citibank statements.")
	}

	log.Info(statements)
}