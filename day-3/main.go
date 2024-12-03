package main

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	P1()
	P2()
}

func P1() {
	input, _ := os.ReadFile("test.txt")
	strs := strings.Split(string(input), "mul")
	var equations []string
	for _, s := range strs {
		sbytes := []byte(s)
		startIndex := bytes.IndexByte(sbytes, '(')
		endIndex := bytes.IndexByte(sbytes, ')')
		if sbytes[0] != '(' || (startIndex == -1 || endIndex == -1) || (endIndex < startIndex) {
			continue
		}

		equation := s[startIndex : endIndex+1]
		match, _ := regexp.Match(`^\([0-9]{1,3},[0-9]{1,3}\)`, []byte(equation))
		if match {
			equations = append(equations, equation)
		}
	}

	result := 0
	for _, equation := range equations {
		var n1, n2 int
		fmt.Sscanf(equation, "(%d,%d)", &n1, &n2)
		result += n1 * n2
	}
	fmt.Println(result)
}

func P2() {
}
