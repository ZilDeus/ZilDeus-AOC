package main

import (
	"fmt"
	"os"
	"slices"
	_ "slices"
	"strings"
)

func main() {
	P1()
	P2()
}

func get_last_index(rules [][2]int) int {
	numbers_map := make(map[int]int)
	left_side_map := make(map[int]int)
	for i := 0; i < len(rules); i++ {
		numbers_map[rules[i][1]] = 0
		numbers_map[rules[i][0]] = 0
		left_side_map[rules[i][0]]++
	}
	for k := range numbers_map {
		_, exist_left := left_side_map[k]
		if !exist_left {
			return k
		}
	}
	return 0
}

func convert_rules(rules []string) [][2]int {
	var converted_rules [][2]int
	for _, rule := range rules {

		//??
		if len(rule) == 0 {
			continue
		}

		var numbers [2]int
		fmt.Sscanf(rule, "%d|%d", &numbers[0], &numbers[1])
		converted_rules = append(converted_rules, numbers)
	}
	return converted_rules
}

func convert_ordering(orderings []string) [][]int {
	var converted_orderings [][]int
	for _, ordering := range orderings {

		//??
		if len(ordering) == 0 {
			continue
		}

		numbers_string := strings.Split(ordering, ",")
		var numbers []int

		for _, v := range numbers_string {
			var n int
			fmt.Sscanf(v, "%d", &n)
			numbers = append(numbers, n)
		}
		converted_orderings = append(converted_orderings, numbers)
	}
	return converted_orderings
}

//func get_number_to_left(n int, rules [][2]int) []int {
//	return []int{}
//}

func is_in_order(left int, n int, rules [][2]int) bool {
	//happy ending
	for _, rule := range rules {
		if rule[0] == left && rule[1] == n {
			return true
		}
	}
	return false
	//i don't even want to think about this
	//left_numbers := get_number_to_left(n, rules)
	//return len(left_numbers) != 0 && slices.Contains(left_numbers, left)
}

func is_ordering_valid(ordering []int, rules [][2]int) bool {
	if len(ordering) == 1 {
		return true
	}

	for i := 1; i < len(ordering); i++ {
		if !is_in_order(ordering[i-1], ordering[i], rules) {
			return false
		}
	}
	return true
}
func correctly_order(ordering []int, rules [][2]int) []int {
	is_correct := false
	for {
		is_correct = true
		for _, rule := range rules {

			first_idx := slices.Index(ordering, rule[0])
			second_idx := slices.Index(ordering, rule[1])
			if first_idx == -1 || second_idx == -1 || second_idx > first_idx {
				continue
			}
			is_correct = false
			ordering[second_idx], ordering[first_idx] = ordering[first_idx], ordering[second_idx]
		}
		if is_correct {
			break
		}
	}
	return ordering
}

func P1() {
	input, _ := os.ReadFile("test.txt")
	splited_input := strings.Split(string(input), "\n\n")
	rules := convert_rules(strings.Split(splited_input[0], "\n"))
	last_element := get_last_index(rules)
	orderings := convert_ordering(strings.Split(splited_input[1], "\n"))
	result := 0
	for _, ordering := range orderings {
		if ordering[0] == last_element && len(ordering) != 1 {
			continue
		}
		if is_ordering_valid(ordering, rules) {
			result += ordering[len(ordering)/2]
		}
	}
	fmt.Println(result)
}

func P2() {
	input, _ := os.ReadFile("test.txt")
	splited_input := strings.Split(string(input), "\n\n")
	rules := convert_rules(strings.Split(splited_input[0], "\n"))
	last_element := get_last_index(rules)
	orderings := convert_ordering(strings.Split(splited_input[1], "\n"))
	result := 0
	for _, ordering := range orderings {
		if ordering[0] == last_element && len(ordering) != 1 {
			continue
		}
		if is_ordering_valid(ordering, rules) {
			//i DIDNT READ THE FUCKING FUCK FUCK FUCCKKKKKKKKKK 4 hours of
			//debugging wasted fuck fuck fuck fuck
			//result += ordering[len(ordering)/2]
		} else {
			correct_order := correctly_order(ordering, rules)
			result += correct_order[len(correct_order)/2]
		}
	}
	fmt.Println(result)
}
