package transaction

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

// StatementType is a struct
type StatementType struct {
	provider string // fixme enum
	records []RecordType
}

// RecordType is a struct
type RecordType struct {
	date string
	amount int
	description string
}

// Statement returns a new StatementType struct
func Statement(provider string, records []RecordType) StatementType {
	return StatementType{
		provider: provider,
		records: records,
	}
}

// Record returns a new RecordType struct
func Record(date string, amount int, description string) RecordType {
	return RecordType{
		date: date,
		amount: amount,
		description: description,
	}
}

// Print prints the underlying data of a Statement
func (s *StatementType) Print() {
	log.Info(fmt.Sprintf("Statement:\n\tProvider: %s\n\tRecords:", s.provider))
	records := s.records

	for _, r := range records {
		r.Print()
	}
}

// Print prints the underlying data of a Record
func (r *RecordType) Print() {
	log.Info(fmt.Sprintf(
		"\tDate: %s  Amount: %d  Description %s",
		r.date,
		r.amount,
		r.description,
	))
}
