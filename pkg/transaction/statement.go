package transaction

import (
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
)

// StatementType is a struct
type StatementType struct {
	Provider string // fixme enum
	Records []RecordType
}

// Statement returns a new StatementType struct
func Statement(provider string, records []RecordType) StatementType {
	return StatementType{
		Provider: provider,
		Records: records,
	}
}

// Print prints the underlying data of a Statement
func (s *StatementType) Print() {
	log.Info(fmt.Sprintf("Statement:\n\tProvider: %s\n\tRecords:", s.Provider))
	records := s.Records

	for _, r := range records {
		r.Print()
	}
}

// StartDate returns the start date of a given statement
func (s *StatementType) StartDate() time.Time {
	if s != nil {
		if s.Records != nil {
			if len(s.Records) > 0 {
				return s.Records[0].Date
			}
		}
	}

	log.Info("record contained no start date")
	return time.Now()
}

// EndDate returns the end date of a given statement
func (s *StatementType) EndDate() time.Time {
	if s != nil {
		if len(s.Records) > 0 {
			if s.Records != nil {
				numRecords := len(s.Records)
					return s.Records[numRecords].Date
			}
		}
	}

	log.Info("record contained no final date")
	return time.Now()
}
