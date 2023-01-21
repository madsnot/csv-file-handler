package parsers

import "fmt"

func ParserStrToExpr(str string) (expr []string, err error) {
	errMsg := "Incorrect expression : %s"
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

	if arg != "" {
		expr = append(expr, arg)
	}

	endInd := len(expr)
	if expr[endInd-1] == "+" || expr[endInd-1] == "-" || expr[endInd-1] == "*" || expr[endInd-1] == "/" {
		return nil, fmt.Errorf(errMsg, str)
	}

	return expr, nil
}
