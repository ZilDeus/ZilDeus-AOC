package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	P1()
	//P2()
}

type Equation struct {
	answer  int
	numbers []int
}

const (
	ADD      = iota
	MULTIPLY = iota
)

var operations = []int{ADD, MULTIPLY}

func operate(a int, b int, op int) int {
	switch op {
	case ADD:
		return a + b
	default:
		return a * b
	}
}

func get_permutations_size(n int) uint64 {
	result := uint64(1)
	for i := 1; i <= n; i++ {
		result *= 2
	}
	return result
}

func get_equations(input []string) []Equation {
	var equations []Equation
	for _, line := range input {
		if len(line) == 0 {
			continue
		}
		var equation Equation
		line_splited := strings.Split(line, ":")
		fmt.Sscanf(line_splited[0], "%d", &equation.answer)
		line_splited[1] = line_splited[1][1:]
		numbers_str := strings.Split(line_splited[1], " ")
		for _, str := range numbers_str {
			n := 0
			fmt.Sscanf(str, "%d", &n)
			equation.numbers = append(equation.numbers, n)
		}
		equations = append(equations, equation)
	}
	return equations
}

func generate_operations_permutaions(size int) [][]int {
	size--
	var perms [][]int
	perms_count := get_permutations_size(size)
	for {
		if len(perms) >= int(perms_count) {
			return perms
		}
		perm := []int{}
		if len(perms) == 0 {
			for i := 0; i < size; i++ {
				perm = append(perm, ADD)
			}
			perms = append(perms, perm)
		} else {
			perm = slices.Clone(perms[len(perms)-1])
			for j, p := range perm {
				if p == ADD {
					perm[j] = MULTIPLY
				} else {
					perm[j] = ADD
				}
				if slices.IndexFunc(perms, func(e []int) bool {
					return slices.Equal(e, perm)
				}) != -1 {
					perm[j] = p
				} else {
					perms = append(perms, perm)
				}
			}
		}
	}
}

func is_equation_valid(equation Equation) bool {
	stacks := generate_operations_permutaions(len(equation.numbers))
	result := equation.numbers[0]
	for _, stack := range stacks {
		result = equation.numbers[0]
		for i := 1; i < len(equation.numbers); i++ {
			result = operate(result, equation.numbers[i], stack[i-1])
		}
		if result == equation.answer {
			return true
		}
	}
	return false
}

func P1() {
	data, _ := os.ReadFile("test.txt")
	input := strings.Split(string(data), "\n")
	equations := get_equations(input)
	result := 0
	for _, eq := range equations {
		if is_equation_valid(eq) {
			result += eq.answer
		}
	}
	fmt.Println(result)
}

func P2() {
	data, _ := os.ReadFile("test.txt")
	input := strings.Split(string(data), "\n")
	result := 0
	fmt.Println(input, result)
}
