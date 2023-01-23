package parsers

import (
	"encoding/csv"
	"fmt"
	"mods/handler/validator"
	"os"
	"strings"
)

func ParserFromCSV(fileName string) (table map[string]string, equations []string, err error) {
	errMsg := []string{"Invalid table structure.",
		"Invalid column/row name: %s/%s.",
		"Incorrect expression in cell: %s."}

	file, err := os.Open(fileName)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ','
	records, err := reader.ReadAll()
	if err != nil {
		return nil, nil, err
	}

	columnName := strings.Split(records[0][0], "\t")
	if !validator.ValidateSequence(columnName) {
		return nil, nil, fmt.Errorf(errMsg[0])
	}

	lenRow := len(columnName)
	table = make(map[string]string, 0)
	prevRowName := ""

	for _, record := range records[1:] {
		row := strings.Split(record[0], "\t")

		if len(row) != lenRow {
			return nil, nil, fmt.Errorf(errMsg[0])
		}

		rowName := row[0]
		if prevRowName == rowName {
			return nil, nil, fmt.Errorf(errMsg[0])
		} else {
			prevRowName = rowName
		}

		for ind, val := range row[1:] {
			if !validator.ValidateStr(columnName[ind+1]) || !validator.ValidateNum(rowName) {
				return nil, nil, fmt.Errorf(errMsg[1], columnName[ind+1], rowName)
			}

			valType, ok := validator.ValidateValue(val)
			if !ok {
				return nil, nil, fmt.Errorf(errMsg[2], columnName[ind+1]+rowName)
			}

			if valType == "string" {
				equations = append(equations, columnName[ind+1]+rowName)
			}
			table[columnName[ind+1]+rowName] = val
		}

	}

	return table, equations, nil
}
