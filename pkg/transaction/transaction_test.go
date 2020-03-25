package transaction

import (
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// Not actual unit tests, relies on underlying file system
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