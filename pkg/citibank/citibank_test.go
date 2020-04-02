package citibank

import (
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/weibelben/finance_parser/pkg/transaction"
)

func TestParseStatements(t *testing.T) {
	cases := []struct {
		name     string
		files    []string
		err      string
		expected []transaction.StatementType
	}{
		{
			name:    "noFiles",
			files:   []string{},
			err:	  "",
			expected: []transaction.StatementType{},
		},
	}

	for _, testcase := range cases {
		t.Run(testcase.name, func(t *testing.T) {
			statements, err := parseStatements(testcase.files)
			if err != nil {
				if testcase.err == "" {
					t.Errorf("test %s failed: %s", testcase.name, err)
				}

				if !strings.Contains(err.Error(), testcase.err) {
					t.Errorf("test %s failed: %s", testcase.name, err)
				}
			} else {
				assert.Equal(t, statements, testcase.expected, "unexected statements")
			}
		})
	}
}

func TestParseStatementEntry(t *testing.T) {
	cases := []struct {
		name     string
		err      string
		row      []string
		expected transaction.RecordType
	}{
		{
			name:     "happyPath",
			err:      "",
			row:      []string{
				"Cleared",
				"04/01/2020",
				"Arbitrary purchase.",
				"100.00",
				"",
			},
			expected: transaction.RecordType{
				Date:        time.Date(2020, 04, 01, 0, 0, 0, 0, time.UTC),
				Amount:      100.00,
				Description: "Arbitrary purchase.",
			},
		},
		{
			name:     "tooFewEntries",
			err:      "unexpected length of citibank statement entry",
			row:      []string{
				"04/01/2020",
				"Arbitrary purchase.",
				"100.00",
				"",
			},
			expected: transaction.RecordType{
				Date:        time.Date(2020, 04, 01, 0, 0, 0, 0, time.UTC),
				Amount:      100.00,
				Description: "Arbitrary purchase.",
			},
		},
		{
			name:     "tooManyEntries",
			err:      "unexpected length of citibank statement entry",
			row:      []string{
				"04/01/2020",
				"Arbitrary purchase.",
				"100.00",
				"",
			},
			expected: transaction.RecordType{
				Date:        time.Date(2020, 04, 01, 0, 0, 0, 0, time.UTC),
				Amount:      100.00,
				Description: "Arbitrary purchase.",
			},
		},
		{
			name:     "creditCharge",
			err:      "credit charge instead of debit",
			row:      []string{
				"Cleared",
				"04/01/2020",
				"Arbitrary purchase.",
				"",
				"100.00",
			},
			expected: transaction.RecordType{
				Date:        time.Date(2020, 04, 01, 0, 0, 0, 0, time.UTC),
				Amount:      100.00,
				Description: "Arbitrary purchase.",
			},
		},
	}

	for _, testcase := range cases {
		t.Run(testcase.name, func(t *testing.T) {
			record, err := parseStatementEntry(testcase.row)
			if err != nil {
				if testcase.err == "" {
					t.Errorf("test %s failed: %s", testcase.name, err)
				}

				if !strings.Contains(err.Error(), testcase.err) {
					t.Errorf("test %s failed: %s", testcase.name, err)
				}
			} else {
				assert.Equal(t, record, testcase.expected, "unexected statements")
			}
		})
	}
}