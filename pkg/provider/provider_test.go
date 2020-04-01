package provider

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/weibelben/finance_parser/pkg/transaction"
)

// Not actual unit tests, relies on underlying file system
func TestFindStatements(t *testing.T) {
	cases := []struct {
		name     string
		folder   string
		err      string
		expected []string
	}{
		{
			name:     "nonExistantFolder",
			folder:   "test_statements/nonexistant",
			err:	  "no such file or directory",
			expected: nil,
		},
		{
			name:     "emptyFolder",
			folder:   "test_statements/emptytestfolder",
			err:      "",
			expected: nil,
		},
		{
			name:     "populatedFolder",
			folder:   "test_statements/populatedtestfolder",
			err:	  "",
			expected: []string{"statement1.csv", "statement2.csv"},
		},
		{
			name:     "populatedFolderWithNoise",
			folder:   "test_statements/noisytestfolder",
			err:	  "",
			expected: []string{"statement1.csv", "statement2.csv"},
		},
	}

	for _, testcase := range cases {
		t.Run(testcase.name, func(t *testing.T) {
			folderNames, err := FindStatements(testcase.folder)
			if err != nil {
				if testcase.err == "" {
					t.Errorf("test %s failed: %s", testcase.name, err)
				}

				if !strings.Contains(err.Error(), testcase.err) {
					t.Errorf("test %s failed: %s", testcase.name, err)
				}
			} else {
				assert.Equal(t, folderNames, testcase.expected, "unexected folder names")
			}
		})
	}
}

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

	var prov Provider

	for _, testcase := range cases {
		t.Run(testcase.name, func(t *testing.T) {
			statements, err := ParseStatements(prov, testcase.files)
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
