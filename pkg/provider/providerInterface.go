package provider

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/weibelben/finance_parser/pkg/transaction"
)

// Provider is an interface that includes all of the functions implemented
// by each finance service provider
type Provider interface {
	ParseStatementFiles([]transaction.StatementType, error) ()
	parseStatements(statementFiles []string) ([]transaction.StatementType, error)
	parseRawStatementData([][]string) (transaction.StatementType, error)
	parseStatementEntry([]string) (transaction.RecordType, error)
}

// StatementSyntaxError is raised when an input csv statement has bad syntax
type StatementSyntaxError struct {
    Message string
}

func (e *StatementSyntaxError) Error() string {
    return fmt.Sprintf("statement syntax error: %s", e.Message)
}

// FindStatements returns the names of all the csv files in the
// citibank/ dir
func FindStatements(folder string) ([]string, error) {
	// get current working dir
	root, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	rootFolder := "finance_parser"
	root = strings.Split(root, rootFolder)[0] + rootFolder + "/"

	// read all file names in working dir
	citiDir := root + fmt.Sprintf("statements/%s", folder)
	files, err := ioutil.ReadDir(citiDir)
	if err != nil {
		return nil, err
	}

	// filter to only include csv files
	var statements []string
	for _, file := range files {
		// only include csv files
		if strings.HasSuffix(file.Name(), ".csv") {
			filePath := root + "statements/" + folder + "/" + file.Name()
			statements = append(statements, filePath)
		}
	}

	return statements, nil
}
