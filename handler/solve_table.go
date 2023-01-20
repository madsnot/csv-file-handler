package handler

import (
	"log"
	"mods/handler/parsers"
	"mods/handler/validator"
	"strconv"
)

func SolveTable(table map[string]string, eq []string) {
	var res, arg1, arg2 int
	for _, cell := range eq {
		expr := parsers.ParserStrToExpr(table[cell])
		if !validator.ValidateNum(expr[0]) {
			arg1, _ = strconv.Atoi(table[expr[0]])
		} else {
			arg1, _ = strconv.Atoi(expr[0])
		}
		if !validator.ValidateNum(expr[2]) {
			arg2, _ = strconv.Atoi(table[expr[2]])
		} else {
			arg2, _ = strconv.Atoi(expr[2])
		}
		op := expr[1]
		switch op {
		case "+":
			res = arg1 + arg2
		case "-":
			res = arg1 - arg2
		case "*":
			res = arg1 * arg2
		case "/":
			if arg2 == 0 {
				log.Println("Division by zero in: ", cell)
			} else {
				res = arg1 / arg2
			}
		default:
			res = arg1
		}
		table[cell] = strconv.Itoa(res)
	}
}
