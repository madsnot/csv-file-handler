package main

import (
	"csvhandler/handler"
	"csvhandler/handler/parsers"
	"log"
	"os"
)

func main() {
	var dataTable map[string]string

	fileName := os.Args[1]

	dataTable, equations, err := parsers.ParserFromCSV(fileName)
	if err != nil {
		log.Println(err)
		return
	}

	if len(equations) != 0 {
		err = handler.SolveTable(dataTable, equations)
		if err != nil {
			log.Println(err)
			return
		}
	}

	err = parsers.ParserToCSV(dataTable)
	if err != nil {
		log.Println(err)
		return
	}
}
