package transaction

import (
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRecordConstructor(t *testing.T) {
	cases := []struct {
		name        string
		date        string
		amount      string
		description string
		err         string
		expected    RecordType
	}{
		{
			name:        "happyPath",
			date:        "03/25/2020",
			amount:      "100.00",
			description: "arbitrary purchase",
			err:         "",
			expected:    RecordType{
				Date: time.Date(2020, 03, 25, 0, 0, 0, 0, time.UTC),
				Amount: 100.00,
				Description: "arbitrary purchase",
			},
		},
		{
			name:        "invalidDate",
			date:        "30/25/2020",
			amount:      "100.00",
			description: "arbitrary purchase",
			err:         "parsing time",
			expected:    RecordType{},
		},
		{
			name:        "invalidAmount",
			date:        "03/25/2020",
			amount:      "abc",
			description: "arbitrary purchase",
			err:         "invalid syntax",
			expected:    RecordType{},
		},
	}

	for _, testcase := range cases {
		t.Run(testcase.name, func(t *testing.T) {
			record, err := Record(testcase.date, testcase.amount, testcase.description)
			if err != nil {
				if testcase.err == "" {
					t.Errorf("test %s failed: %s", testcase.name, err)
				}

				if !strings.Contains(err.Error(), testcase.err) {
					t.Errorf("test %s failed: %s", testcase.name, err)
				}
			} else {
				assert.Equal(t, record, testcase.expected, "unexected resulting record")
			}
		})
	}
}

func getMockStatement() StatementType {
	return StatementType {
		Provider: "citibank",
		Records: []RecordType{
			RecordType{
				Date: time.Date(2020, 03, 24, 0, 0, 0, 0, time.UTC),
				Amount: 100.00,
				Description: "arbitrary purchase 1",
			},
			RecordType{
				Date: time.Date(2020, 03, 25, 0, 0, 0, 0, time.UTC),
				Amount: 200.00,
				Description: "arbitrary purchase 2",
			},
			RecordType{
				Date: time.Date(2020, 03, 26, 0, 0, 0, 0, time.UTC),
				Amount: 300.00,
				Description: "arbitrary purchase 3",
			},
		},
	}
}

func getMockRecordSlice() []RecordType {
	return []RecordType{
		RecordType{
			Date: time.Date(2020, 03, 24, 0, 0, 0, 0, time.UTC),
			Amount: 100.00,
			Description: "arbitrary purchase 1",
		},
		RecordType{
			Date: time.Date(2020, 03, 25, 0, 0, 0, 0, time.UTC),
			Amount: 200.00,
			Description: "arbitrary purchase 2",
		},
		RecordType{
			Date: time.Date(2020, 03, 26, 0, 0, 0, 0, time.UTC),
			Amount: 300.00,
			Description: "arbitrary purchase 3",
		},
	}
}

func TestStatementConstructor(t *testing.T) {
	cases := []struct {
		name     string
		provider string
		records  []RecordType
		err      string
		expected StatementType
	}{
		{
			name:     "happyPath",
			provider: "citibank",
			records:  getMockRecordSlice(),
			err:      "",
			expected: getMockStatement(),
		},
		{
			name:     "emptyRecordSlice",
			provider: "citibank",
			records:  []RecordType{},
			err:      "invalid provider",
			expected: StatementType{
				Provider: "citibank",
				Records: []RecordType{},
			},
		},
		{
			name:     "badProvider",
			provider: "unknown",
			records:  getMockRecordSlice(),
			err:      "invalid provider",
			expected: getMockStatement(),
		},
	}

	for _, testcase := range cases {
		t.Run(testcase.name, func(t *testing.T) {
			statement, err := Statement(testcase.provider, testcase.records)
			if err != nil {
				if testcase.err == "" {
					t.Errorf("test %s failed: %s", testcase.name, err)
				}

				if !strings.Contains(err.Error(), testcase.err) {
					t.Errorf("test %s failed: %s", testcase.name, err)
				}
			} else {
				assert.Equal(t, statement, testcase.expected, "unexected resulting statement")
			}
		})
	}
}

func TestIsValidProvider(t *testing.T) {
	cases := []struct {
		name     string
		provider string
		expected bool
	}{
		{
			name:     "happyPathCiti",
			provider: "citibank",
			expected: true,
		},
		{
			name:     "happyPathDiscover",
			provider: "discover",
			expected: true,
		},
		{
			name:     "happyPathBECU",
			provider: "becu",
			expected: true,
		},
		{
			name:     "unknownProvider",
			provider: "unknown",
			expected: false,
		},
	}

	for _, testcase := range cases {
		t.Run(testcase.name, func(t *testing.T) {
			assert.Equal(t, isValidProvider(testcase.provider), testcase.expected, "unexpected result")
		})
	}
}

func TestStartDate(t *testing.T) {
	cases := []struct {
		name      string
		statement StatementType
		err       string
		expected  time.Time
	}{
		{
			name:      "happyPath",
			statement: getMockStatement(),
			err:       "",
			expected:  time.Date(2020, 03, 24, 0, 0, 0, 0, time.UTC),
		},
		{
			name:      "nilStatement",
			statement: StatementType{},
			err:       "record contained no start date",
			expected:  time.Date(2020, 03, 24, 0, 0, 0, 0, time.UTC),
		},
		{
			name:      "emptyRecordSlice",
			statement: StatementType{
				Provider: "citibank",
				Records: []RecordType{},
			},
			err:       "record contained no start date",
			expected:  time.Date(2020, 03, 24, 0, 0, 0, 0, time.UTC),
		},
	}

	for _, testcase := range cases {
		t.Run(testcase.name, func(t *testing.T) {
			date, err := testcase.statement.StartDate()
			if err != nil {
				if testcase.err == "" {
					t.Errorf("test %s failed: %s", testcase.name, err)
				}

				if !strings.Contains(err.Error(), testcase.err) {
					t.Errorf("test %s failed: %s", testcase.name, err)
				}
			} else {
				assert.Equal(t, date, testcase.expected, "unexected resulting date")
			}
		})
	}
}

func TestEndDate(t *testing.T) {
	cases := []struct {
		name      string
		statement StatementType
		err       string
		expected  time.Time
	}{
		{
			name:      "happyPath",
			statement: getMockStatement(),
			err:       "",
			expected:  time.Date(2020, 03, 26, 0, 0, 0, 0, time.UTC),
		},
		{
			name:      "nilStatement",
			statement: StatementType{},
			err:       "record contained no final date",
			expected:  time.Date(2020, 03, 26, 0, 0, 0, 0, time.UTC),
		},
		{
			name:      "emptyRecordSlice",
			statement: StatementType{
				Provider: "citibank",
				Records: []RecordType{},
			},
			err:       "record contained no final date",
			expected:  time.Date(2020, 03, 26, 0, 0, 0, 0, time.UTC),
		},
	}

	for _, testcase := range cases {
		t.Run(testcase.name, func(t *testing.T) {
			date, err := testcase.statement.EndDate()
			if err != nil {
				if testcase.err == "" {
					t.Errorf("test %s failed: %s", testcase.name, err)
				}

				if !strings.Contains(err.Error(), testcase.err) {
					t.Errorf("test %s failed: %s", testcase.name, err)
				}
			} else {
				assert.Equal(t, date, testcase.expected, "unexected resulting date")
			}
		})
	}
}
