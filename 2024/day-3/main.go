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
	//very lazy solution , but it works
	is_mul_active := true
	data, _ := os.ReadFile("test.txt")
	var equations []string
	for i := 0; i < len(data); i++ {
		if (len(data) - i - 1) <= 7 {
			break
		}

		if strings.Compare(string(data[i:i+4]), "do()") == 0 {
			is_mul_active = true
		}

		if strings.Compare(string(data[i:i+7]), "don't()") == 0 {
			is_mul_active = false
		}
		if is_mul_active && strings.Compare(string(data[i:i+3]), "mul") == 0 {
			sbytes := []byte(data[i+3:])
			startIndex := bytes.IndexByte(sbytes, '(')
			endIndex := bytes.IndexByte(sbytes, ')')
			if sbytes[0] != '(' || (startIndex == -1 || endIndex == -1) || (endIndex < startIndex) {
				continue
			}

			equation := string(sbytes[startIndex : endIndex+1])
			match, _ := regexp.Match(`^\([0-9]{1,3},[0-9]{1,3}\)`, []byte(equation))
			if match {
				equations = append(equations, equation)
			}
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
