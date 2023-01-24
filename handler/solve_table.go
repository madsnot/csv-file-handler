package handler

import (
	"csvhandler/handler/models"
	"csvhandler/handler/parsers"
	"csvhandler/handler/validator"
	"fmt"
	"strconv"
)

func SolveTable(table *models.DataTable) error {
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

	//решаем выражения пока все не решим
	lenEq := len(table.Equations)
	for lenEq > 0 {

		//флаг, показывающий возможно ли вообще решить хоть одно из выражений
		//за один проход по списку
		statusTableSolving = false

		//проходим по массиву с именами ячеек, где есть выражения
		for ind := 0; ind < lenEq; ind++ {
			cell := table.Equations[ind]

			//парсим значение в виде строки в удобное для решения выражение
			expr, err = parsers.ParserStrToExpr(table.Table[cell])
			if err != nil {
				return err
			}

			//проверяем первый аргумент выражения и преобразуем в число
			if !validator.ValidateNum(expr[0]) {
				str := []rune(table.Table[expr[0]])

				//проверяем ячейку, что она хранит (выражение/число)
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

			//продолжаем работать с выражением, обрабатывая операнд и второй аргумент,
			//если это выражение вида arg1 op arg2
			if len(expr) == 3 {
				if !validator.ValidateNum(expr[2]) {
					str := []rune(table.Table[expr[2]])

					//проверяем ячейку, что она хранит (выражение/число)
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

			//вычисляем результат выражения на основе опернда
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

			//записываем результат в таблицу
			table.Table[cell] = strconv.Itoa(res)

			//так как выражение решилось, мы его удаляем из списка
			if lenEq > 1 {
				table.Equations[ind] = table.Equations[lenEq-1]
				table.Equations[lenEq-1] = ""
				table.Equations = table.Equations[:lenEq-1]
			}

			statusTableSolving = true
			arg1, arg2, op = 0, 0, ""
			lenEq--
		}

		//проверяем, если за один проход по списку не одно выражение не решилось,
		//то таблицу решить невозможно
		if !statusTableSolving {
			return fmt.Errorf(errMsg[3])
		}
	}

	return nil
}
