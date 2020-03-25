package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/weibelben/finance_parser/pkg/citibank"
	"github.com/weibelben/finance_parser/pkg/transaction"
)

func main() {
	statements := parseStatements()
}

func parseStatements() {
	citiStatements, err := parseCitiStatements()
	if err != nil {
		log.WithError(err).Error("unable to parse Citibank statements")
	}
}