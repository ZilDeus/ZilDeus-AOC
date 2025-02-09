package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
)

func main() {
	P1()
	P2()
}
func P1() {

	var list1 []int
	var list2 []int

	data, _ := os.ReadFile("test.txt")

	lines := strings.Split(string(data), "\n")
	lines_count := len(lines) - 1

	for i := 0; i < lines_count; i++ {
		var n1 int
		var n2 int
		fmt.Sscanf(lines[i], "%d   %d", &n1, &n2)
		list1 = append(list1, n1)
		list2 = append(list2, n2)
	}

	slices.Sort(list1)
	slices.Sort(list2)

	result := 0

	for i := 0; i < len(list1); i++ {
		result += int(math.Abs(float64(list1[i] - list2[i])))
	}

	fmt.Println(result)
}

func P2() {

	var list1 []int
	var list2 []int

	ls_map := make(map[int]int)

	data, _ := os.ReadFile("test.txt")

	lines := strings.Split(string(data), "\n")
	lines_count := len(lines) - 1

	for i := 0; i < lines_count; i++ {
		var n1 int
		var n2 int
		fmt.Sscanf(lines[i], "%d   %d", &n1, &n2)
		list1 = append(list1, n1)
		list2 = append(list2, n2)
	}

	for i := 0; i < lines_count; i++ {
		ls_map[list1[i]] = 0
	}

	for i := 0; i < lines_count; i++ {
		_, exists := ls_map[list2[i]]
		if exists {
			ls_map[list2[i]]++
		}
	}

	result := 0

	for i := 0; i < lines_count; i++ {
		result += list1[i] * ls_map[list1[i]]
	}

	fmt.Println(result)
}
