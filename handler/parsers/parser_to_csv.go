package parsers

import (
	"csvhandler/handler/models"
	"encoding/csv"
	"os"
)

func ParserToCSV(table *models.DataTable) (err error) {
	var records [][]string

	//записываем таблицу в матрицу для дальнейшей записи в csv формат
	records = append(records, table.Columns)
	for _, row := range table.Rows {
		recRow := make([]string, len(table.Columns))
		recRow[0] = row
		ind := 1

		for _, column := range table.Columns[1:] {
			recRow[ind] = table.Table[column+row]
			ind++
		}

		records = append(records, recRow)
	}

	//выводим в консоль получившуюся таблицу в формате csv
	writer := csv.NewWriter(os.Stdout)
	writer.Comma = ','
	err = writer.WriteAll(records)
	if err != nil {
		return err
	}

	return nil
}
