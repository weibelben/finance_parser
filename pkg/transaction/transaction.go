package transaction

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
