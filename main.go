package main

import (
	"log"
	"mods/handler"
	"mods/handler/parsers"
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

	err = handler.SolveTable(dataTable, equations)
	if err != nil {
		log.Println(err)
		return
	}

	err = parsers.ParserToCSV(dataTable)
	if err != nil {
		log.Println(err)
		return
	}
}
