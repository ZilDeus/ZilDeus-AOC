package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func isStable(n []int) int {
	if !isInOrder(n) {
		return 0
	}
	for i := 1; i < len(n); i++ {
		diff := int(math.Abs(float64(n[i-1] - n[i])))
		if diff < 1 || diff > 3 {
			return 0
		}
	}
	return 1
}
func isInOrder(n []int) bool {
	l := len(n)
	isAsc := func() bool {
		for i := 1; i < l; i++ {
			if n[i-1] < n[i] {
				return false
			}
		}
		return true
	}()

	isDesc := func() bool {
		for i := 1; i < l; i++ {
			if n[i-1] > n[i] {
				return false
			}
		}
		return true
	}()

	return isAsc || isDesc
}

func main() {
	P1()
	//P2()
}
func P1() {

	var list [][]int

	data, _ := os.ReadFile("test.txt")

	lines := strings.Split(string(data), "\n")
	lines_count := len(lines) - 1

	for i := 0; i < lines_count; i++ {
		number_string := strings.Split(lines[i], " ")
		numbers := make([]int, len(number_string))
		for i, v := range number_string {
			var n int
			fmt.Sscanf(v, "%d", &n)
			numbers[i] = n
		}
		list = append(list, numbers)
	}

	result := 0

	for i := 0; i < len(list); i++ {
		result += isStable(list[i])
	}

	fmt.Println(result)
}

func P2() {
}
