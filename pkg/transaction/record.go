package transaction

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

// RecordType is a struct
type RecordType struct {
	date string
	amount int
	description string
}

// Record returns a new RecordType struct
func Record(date string, amount int, description string) RecordType {
	return RecordType{
		date: date,
		amount: amount,
		description: description,
	}
}

func (r *RecordType) Classify() {
	// FIXME TODO
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
