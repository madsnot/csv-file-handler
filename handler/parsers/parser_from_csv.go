package parsers

import (
	"csvhandler/handler/models"
	"csvhandler/handler/validator"
	"encoding/csv"
	"fmt"
	"os"
)

func ParserFromCSV(fileName string) (table *models.DataTable, err error) {
	errMsg := []string{"Invalid table structure.",
		"Invalid column name: %s.",
		"Invalid row name: %s.",
		"Incorrect expression in cell: %s."}

	//открываем файл
	file, err := os.Open("./tests/" + fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	//читаем из csv файла все строки
	reader := csv.NewReader(file)
	reader.Comma = ','
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	//берем первую строку считанной из csv таблицы,
	//первая строка - имена столбцов
	columnName := records[0]
	if !validator.ValidateSequence(columnName) {
		return nil, fmt.Errorf(errMsg[0])
	}

	//создаем новую таблицу
	table = models.NewDataTable()

	//добавляем в структуру массив имен столбцов, проверяя имена на правильность
	table.Columns = append(table.Columns, columnName[0])
	for _, str := range columnName[1:] {
		if !validator.ValidateStr(str) {
			return nil, fmt.Errorf(errMsg[1], str)
		} else {
			table.Columns = append(table.Columns, str)
		}
	}

	lenRow := len(columnName)               //количество столбцов
	rowNamesMap := make(map[string]bool, 0) //карта имен строк

	//по одной строке обрабатываем остальные строки таблицы
	for _, row := range records[1:] {

		//проверяем на равность количество столбцов в строке с данными
		//и с количеством столбцов
		if len(row) != lenRow {
			return nil, fmt.Errorf(errMsg[0])
		}

		rowName := row[0]

		//проверяем имя строки, что это число
		if rowName == "" || !validator.ValidateNum(rowName) {
			return nil, fmt.Errorf(errMsg[2], rowName)
		}

		//проверяем на отличие имени данной строки от других строк,
		//добавляем в массив имен строк таблицы
		if _, ok := rowNamesMap[rowName]; ok {
			return nil, fmt.Errorf(errMsg[0])
		} else {
			table.Rows = append(table.Rows, rowName)
			rowNamesMap[rowName] = true
		}

		//проходимся по каждому значению в полученной строке
		for ind, val := range row[1:] {
			//проверяем значение в ячейке (может быть выражением или целым числом)
			valType, ok := validator.ValidateValue(val)
			if !ok {
				return nil, fmt.Errorf(errMsg[3], columnName[ind+1]+rowName)
			}

			//проверяем тип значения в ячейке на выражение,
			//если это выражение, добавляем имя ячейки в список
			if valType == "string" {
				table.Equations = append(table.Equations, columnName[ind+1]+rowName)
			}

			//добавляем значение в таблицу
			table.Table[columnName[ind+1]+rowName] = val
		}

	}

	return table, nil
}
