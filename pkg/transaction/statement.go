package transaction

import (
	"errors"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
)

// PROVIDERS is a list of all the valid providers
var PROVIDERS = []string{"citibank", "discover", "becu"}

// StatementType is a struct
type StatementType struct {
	Provider string // fixme enum
	Records []RecordType
}

// Statement returns a new StatementType struct
func Statement(provider string, records []RecordType) (StatementType, error) {
	var statement StatementType
	if !isValidProvider(provider) {
		return statement, fmt.Errorf("invalid provider, expected one of %v", PROVIDERS)
	}

	return StatementType{
		Provider: provider,
		Records: records,
	}, nil
}

func isValidProvider(provider string) bool {
	for _, prov := range PROVIDERS {
		if prov == provider {
			return true
		}
	}
	return false
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
func (s *StatementType) StartDate() (time.Time, error) {
	if s != nil {
		if s.Records != nil {
			if len(s.Records) > 0 {
				return s.Records[0].Date, nil
			}
		}
	}

	return time.Now(), errors.New("record contained no start date")
}

// EndDate returns the end date of a given statement
func (s *StatementType) EndDate() (time.Time, error) {
	if s != nil {
		if len(s.Records) > 0 {
			if s.Records != nil {
				numRecords := len(s.Records)
					return s.Records[numRecords - 1].Date, nil
			}
		}
	}

	return time.Now(), errors.New("record contained no final date")
}
