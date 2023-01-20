package parsers

import "log"

func ParserStrToExpr(str string) (expr []string) {
	arg := ""
	for _, symb := range str[1:] {
		symbStr := string(symb)
		if symbStr == "+" || symbStr == "-" || symbStr == "*" || symbStr == "/" {
			expr = append(expr, arg, symbStr)
			arg = ""
		} else {
			arg += symbStr
		}
	}
	if arg == "+" || arg == "-" || arg == "*" || arg == "/" {
		log.Fatal("Nonconform table")
	}
	if arg != "" {
		expr = append(expr, arg)
	}

	return expr
}
