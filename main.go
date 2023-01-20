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
	}
	handler.SolveTable(dataTable, equations)
	parsers.ParserToCSV(dataTable)
}
