package handler

import (
	"csvhandler/handler/parsers"
	"csvhandler/handler/validator"
	"fmt"
	"strconv"
)

func SolveTable(table map[string]string, eq []string) error {
	var (
		res, arg1, arg2    int
		op                 string
		expr               []string
		err                error
		statusTableSolving bool
	)
	errMsg := []string{"Invalid expression. ",
		"Unknown %s cell value.",
		"Division by zero in: %s.",
		"Invalid table structure."}

	lenEq := len(eq)
	for lenEq > 0 {

		statusTableSolving = false

		for ind := 0; ind < lenEq; ind++ {
			cell := eq[ind]

			expr, err = parsers.ParserStrToExpr(table[cell])
			if err != nil {
				return err
			}
			if !validator.ValidateNum(expr[0]) {
				str := []rune(table[expr[0]])
				if len(str) == 0 {
					return fmt.Errorf(errMsg[0]+errMsg[1], expr[0])
				}
				if str[0] == '=' {
					continue
				}
				arg1, err = strconv.Atoi(string(str))
			} else {
				arg1, err = strconv.Atoi(expr[0])
			}
			if err != nil {
				return fmt.Errorf(errMsg[0]+errMsg[1], expr[0])
			}

			if len(expr) == 3 {
				if !validator.ValidateNum(expr[2]) {
					str := []rune(table[expr[2]])
					if len(str) == 0 {
						return fmt.Errorf(errMsg[0]+errMsg[1], expr[2])
					}
					if str[0] == '=' {
						continue
					}
					arg2, err = strconv.Atoi(string(str))
				} else {
					arg2, err = strconv.Atoi(expr[2])
				}
				if err != nil {
					return fmt.Errorf(errMsg[0]+errMsg[1], expr[2])
				}
				op = expr[1]
			}

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

			if lenEq > 1 {
				eq[ind] = eq[lenEq-1]
				eq[lenEq-1] = ""
				eq = eq[:lenEq-1]
			}
			statusTableSolving = true
			arg1, arg2, op = 0, 0, ""
			lenEq--
		}

		if !statusTableSolving {
			return fmt.Errorf(errMsg[3])
		}
	}

	return nil
}
