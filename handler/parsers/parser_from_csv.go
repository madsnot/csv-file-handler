package parsers

import (
	"encoding/csv"
	"fmt"
	"mods/handler/validator"
	"os"
	"strings"
)

func ParserFromCSV(fileName string) (table map[string]string, equations []string, err error) {
	errMsg := []string{"Invalid column/row name: %s/%s",
		"Incorrect expression in cell: %s"}

	file, err := os.Open(fileName)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ','
	record, err := reader.Read()
	if err != nil {
		return nil, nil, err
	}
	columns := strings.Split(record[0], "\t")

	records, err := reader.ReadAll()
	if err != nil {
		return nil, nil, err
	}

	table = make(map[string]string, 0)
	for _, record := range records {
		rows := strings.Split(record[0], "\t")

		for ind, val := range rows[1:] {
			str := rows[0]
			if !validator.ValidateStr(columns[ind+1]) || !validator.ValidateNum(str) {
				return nil, nil, fmt.Errorf(errMsg[0], columns[ind+1], str)
			}

			valType, ok := validator.ValidateValue(val)
			if !ok {
				return nil, nil, fmt.Errorf(errMsg[1], columns[ind+1]+str)
			}

			if valType == "string" {
				equations = append(equations, columns[ind+1]+str)
			}
			table[columns[ind+1]+str] = val
		}

	}

	return table, equations, nil
}
