package provider

import (
	"github.com/weibelben/finance_parser/pkg/transaction"
)

// Provider is an interface that includes all of the functions implemented
// by each finance service provider
type Provider interface {
	ParseStatements(transaction.StatementType, error) ()
}
