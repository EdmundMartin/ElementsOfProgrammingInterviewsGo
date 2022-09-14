package ch8_stacks_queues

import (
	"strconv"
	"strings"
)

func exprAdd(x, y int) int {
	return x + y
}

func exprSub(x, y int) int {
	return x - y
}

func exprDiv(x, y int) int {
	return x / y
}

func exprMul(x, y int) int {
	return x * y
}

func EvaluateRPN(expressions string) int {
	intermediateResults := []int{}

	delimiter := ","
	operators := map[string]func(x, y int) int{
		"+": exprAdd,
		"-": exprSub,
		"/": exprDiv,
		"*": exprMul,
	}

	for _, token := range strings.Split(expressions, delimiter) {
		function, ok := operators[token]
		if ok {
			var first, second int
			first, intermediateResults = intermediateResults[len(intermediateResults)-1], intermediateResults[:len(intermediateResults)-1]
			second, intermediateResults = intermediateResults[len(intermediateResults)-1], intermediateResults[:len(intermediateResults)-1]
			intermediateResults = append(intermediateResults, function(first, second))
		} else {
			res, _ := strconv.Atoi(token)
			intermediateResults = append(intermediateResults, res)
		}
	}
	return intermediateResults[len(intermediateResults)-1]
}
