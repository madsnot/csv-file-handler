package parsers

import "fmt"

func ParserStrToExpr(str string) (expr []string, err error) {
	errMsg := "Incorrect expression : %s."
	arg := ""

	//проходим выражение, чтобы выделить отдельно первый аргумент,
	//операнд и второй аргумент
	for _, symb := range str[1:] {
		symbStr := string(symb)
		if symbStr == "+" || symbStr == "-" || symbStr == "*" || symbStr == "/" {
			expr = append(expr, arg, symbStr)
			arg = ""
		} else {
			arg += symbStr
		}
	}

	//проверяем переменную хранящую аргумент выражения,
	//если она не пуста, то аргумент надо сохранить в массив (выражение)
	if arg != "" {
		expr = append(expr, arg)
	} else {
		return nil, fmt.Errorf(errMsg, str)
	}

	return expr, nil
}
