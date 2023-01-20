package parsers

import (
	"encoding/csv"
	"os"
	"sort"
	"unicode"
)

func takeColumnAndRow(key string) (string, string) {
	column := ""
	row := ""
	for _, symb := range key {
		if unicode.IsDigit(symb) {
			row += string(symb)
		} else {
			column += string(symb)
		}
	}
	return row, column
}

func ParserToCSV(table map[string]string) {
	var columns, rows []string
	var records [][]string
	columnsMap := make(map[string]bool)
	rowsMap := make(map[string]bool)
	columns = append(columns, "")
	for key := range table {
		row, column := takeColumnAndRow(key)
		if _, ok := columnsMap[column]; !ok {
			columnsMap[column] = true
			columns = append(columns, column)
		}
		if _, ok := rowsMap[row]; !ok {
			rowsMap[row] = true
			rows = append(rows, row)
		}
	}
	sort.Strings(columns)
	sort.Strings(rows)
	records = append(records, columns)
	for _, row := range rows {
		recRow := make([]string, len(columns))
		recRow[0] = row
		ind := 1
		for _, column := range columns[1:] {
			recRow[ind] = table[column+row]
			ind++
		}
		records = append(records, recRow)
	}
	writer := csv.NewWriter(os.Stdout)
	writer.Comma = ','
	writer.WriteAll(records)
}
