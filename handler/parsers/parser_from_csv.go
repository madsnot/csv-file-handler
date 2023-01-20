package parsers

import (
	"encoding/csv"
	"mods/handler/validator"
	"os"
	"strings"
)

func ParserFromCSV(fileName string) (table map[string]string, equations []string, err error) {
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
	table = make(map[string]string, 0)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, nil, err
	}
	for _, record := range records {
		rows := strings.Split(record[0], "\t")
		for ind, val := range rows[1:] {
			str := rows[0]
			if !validator.ValidateStr(columns[ind+1]) || !validator.ValidateNum(str) {
				return nil, nil, err
			}
			valType, ok := validator.ValidateExpr(val)
			if !ok {
				return nil, nil, err
			}
			if valType == "string" {
				equations = append(equations, columns[ind+1]+str)
			}
			table[columns[ind+1]+str] = val
		}
	}
	return table, equations, err
}
