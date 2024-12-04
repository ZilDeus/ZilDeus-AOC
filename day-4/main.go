package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	P1()
	//P2()
}

func compare_word(word string) bool {
	xmas := "XMAS"
	samx := "SAMX"
	return strings.Compare(xmas, word) == 0 || strings.Compare(samx, word) == 0
}

func P1() {
	input, _ := os.ReadFile("test.txt")
	letter_board := strings.Split(string(input), "\n")
	letter_board_numbers := make([][]int, len(letter_board))
	for i := range len(letter_board_numbers) {
		letter_board_numbers[i] = make([]int, len(letter_board[i]))
	}
	result := 0
	for y := 0; y < len(letter_board); y++ {
		for x := 0; x < len(letter_board[y]); x++ {
			if letter_board[y][x] != 'X' && letter_board[y][x] != 'S' {
				continue
			}
			is_horizontal_in_bound := len(letter_board[y])-x-1 >= 3
			is_vertical_in_bound := len(letter_board)-y-1 > 3
			is_horizontal_in_bound_left := x >= 3
			if is_horizontal_in_bound {
				word := letter_board[y][x : x+4]
				if compare_word(word) {
					letter_board_numbers[y][x] = 1
					letter_board_numbers[y][x+1] = 1
					letter_board_numbers[y][x+2] = 1
					letter_board_numbers[y][x+3] = 1
					fmt.Printf("h:%s\n", word)
					result++
				}
			}
			if is_vertical_in_bound {
				word := string([]byte{letter_board[y][x], letter_board[y+1][x], letter_board[y+2][x], letter_board[y+3][x]})
				if compare_word(word) {
					letter_board_numbers[y][x] = 1
					letter_board_numbers[y+1][x] = 1
					letter_board_numbers[y+2][x] = 1
					letter_board_numbers[y+3][x] = 1
					fmt.Printf("v:%s\n", word)
					result++
				}
			}
			if is_vertical_in_bound && is_horizontal_in_bound {
				word := string([]byte{letter_board[y][x], letter_board[y+1][x+1], letter_board[y+2][x+2], letter_board[y+3][x+3]})
				if compare_word(word) {
					letter_board_numbers[y][x] = 1
					letter_board_numbers[y+1][x+1] = 1
					letter_board_numbers[y+2][x+2] = 1
					letter_board_numbers[y+3][x+3] = 1
					fmt.Printf("d:%s\n", word)
					result++
				}
			}
			if is_vertical_in_bound && is_horizontal_in_bound_left {
				word := string([]byte{letter_board[y][x], letter_board[y+1][x-1], letter_board[y+2][x-2], letter_board[y+3][x-3]})
				if compare_word(word) {
					letter_board_numbers[y][x] = 1
					letter_board_numbers[y+1][x-1] = 1
					letter_board_numbers[y+2][x-2] = 1
					letter_board_numbers[y+3][x-3] = 1
					fmt.Printf("~d:%s\n", word)
					result++
				}
			}
		}
	}
	for y := 0; y < len(letter_board_numbers); y++ {
		for x := 0; x < len(letter_board_numbers[y]); x++ {
			if letter_board_numbers[y][x] == 0 {
				fmt.Print(".")
			} else {
				fmt.Print(string(letter_board[y][x]))
			}
		}
		fmt.Println()

	}
	fmt.Println(result)
}
