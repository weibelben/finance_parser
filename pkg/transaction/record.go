package transaction

import (
	"fmt"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
)

// RecordType is a struct
type RecordType struct {
	Date time.Time
	Amount float64
	Description string
}

// Record returns a new RecordType struct
func Record(dateString string, amountString string, description string) (RecordType, error) {
	var record RecordType
	date, err := time.Parse("01/02/2006", dateString)
	if err != nil {
		return record, err
	}

	amount, err := strconv.ParseFloat(amountString, 32)
	if err != nil {
		return record, err
	}

	return RecordType{
		Date: date,
		Amount: amount,
		Description: description,
	}, nil
}

// Classify classifies
func (r *RecordType) Classify() {
	// FIXME TODO
}

// Print prints the underlying data of a Record
func (r *RecordType) Print() {
	log.Info(fmt.Sprintf(
		"\tDate: %s  Amount: %f  Description %s",
		r.Date,
		r.Amount,
		r.Description,
	))
}
