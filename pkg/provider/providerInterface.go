package provider

import (
	"github.com/weibelben/finance_parser/pkg/transaction"
)

// Provider is an interface that includes all of the functions implemented
// by each finance service provider
type Provider interface {
	ParseStatements([]transaction.StatementType, error) ()
}

// findStatements returns the names of all the csv files in the
// citibank/ dir
func findStatements(folder string) ([]string, error) {
	// get current working dir
	root, err := os.Getwd()
	if err != nil {
		return nil, err
	}

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
			statements = append(statements, file.Name())
		}
	}
	
	return statements, nil
}
