package handler

import (
	"fmt"
	"mods/handler/parsers"
	"mods/handler/validator"
	"strconv"
)

func SolveTable(table map[string]string, eq []string) error {
	var (
		res, arg1, arg2 int
		expr            []string
		err             error
	)
	errMsg := []string{"Invalid expression. ",
		"This cell is not in the table: %s",
		"Division by zero in: %s"}

	for _, cell := range eq {
		expr, err = parsers.ParserStrToExpr(table[cell])
		if err != nil {
			return err
		}

		if !validator.ValidateNum(expr[0]) {
			arg1, err = strconv.Atoi(table[expr[0]])
		} else {
			arg1, err = strconv.Atoi(expr[0])
		}
		if err != nil {
			return fmt.Errorf(errMsg[0]+errMsg[1], expr[0])
		}

		if !validator.ValidateNum(expr[2]) {
			arg2, err = strconv.Atoi(table[expr[2]])
		} else {
			arg2, err = strconv.Atoi(expr[2])
		}
		if err != nil {
			return fmt.Errorf(errMsg[0]+errMsg[1], expr[2])
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
				return fmt.Errorf(errMsg[0]+errMsg[2], cell)
			} else {
				res = arg1 / arg2
			}
		default:
			res = arg1
		}

		table[cell] = strconv.Itoa(res)
	}

	return nil
}
