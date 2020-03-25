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

// Statement returns a new StatementType struct
func Statement(provider string, records []RecordType) StatementType {
	return StatementType{
		provider: provider,
		records: records,
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

// StartDate returns the start date of a given statement
func (s *StatementType) StartDate() string { // FIXME use Date class/enum
	if s != nil {
		if s.records != nil {
			if s.records[0] != nil {
				return s.records[0].date
			}
		}
	}

	return ""
}

// EndDate returns the end date of a given statement
func (s *StatementType) EndDate() string { // FIXME use Date class
	if s != nil {
		if s.records != nil {
			numRecords := len(s.records)
			if s.records[numRecords] != nil {
				return s.records[numRecords].date
			}
		}
	}

	return ""
}
