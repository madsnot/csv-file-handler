package main

import (
	"csvhandler/handler"
	"csvhandler/handler/models"
	"csvhandler/handler/parsers"
	"log"
	"os"
)

func main() {
	var dataTable *models.DataTable

	fileName := os.Args[1]

	//парсим csv файл в удобную таблицу для работы с данными,
	//получаем эту таблицу - dataTable и equations - список ячеек,
	//которые содержат выражения
	dataTable, err := parsers.ParserFromCSV(fileName)
	if err != nil {
		log.Println(err)
		return
	}

	//проверка на наличие выражений в таблице
	if len(dataTable.Equations) != 0 {

		//решаем выражения в таблице
		err = handler.SolveTable(dataTable)
		if err != nil {
			log.Println(err)
			return
		}
	}

	//парсим обработанную таблицу в формат csv и
	//выводим результат в консоль
	err = parsers.ParserToCSV(dataTable)
	if err != nil {
		log.Println(err)
		return
	}
}
