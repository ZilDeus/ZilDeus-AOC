package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	P1()
	P2()
}

func generate_signal_map(input []string) [][]rune {
	var signal_map [][]rune
	for _, line := range input {
		if len(line) == 0 {
			continue
		}
		var map_line []rune
		for _, char := range line {
			map_line = append(map_line, char)
		}
		signal_map = append(signal_map, map_line)
	}
	return signal_map
}

func set_antinode(x int, y int, antinodes_map [][]rune) {
	if y < 0 || y >= len(antinodes_map) {
		return
	}
	if x < 0 || x >= len(antinodes_map[0]) {
		return
	}
	fmt.Println(x, y)
	antinodes_map[y][x] = '#'
}

func get_antinodes_coords(antenna_coords [][2]int) [][2]int {
	var antinodes_coords [][2]int
	for i := range antenna_coords {
		for j := range antenna_coords {
			if i == j {
				continue
			}
			x := antenna_coords[i][0] + (antenna_coords[i][0] - antenna_coords[j][0])
			y := antenna_coords[i][1] + (antenna_coords[i][1] - antenna_coords[j][1])
			antinodes_coords = append(antinodes_coords, [2]int{x, y})
		}
	}
	return antinodes_coords
}

func get_antinodes_map(signal [][]rune) [][]rune {
	antinodes_map := make([][]rune, len(signal))
	for i := range len(signal) {
		antinodes_map[i] = make([]rune, len(signal[0]))
		for j := 0; j < len(signal[0]); j++ {
			antinodes_map[i][j] = '.'
		}
	}

	signal_map := make(map[rune][][2]int)

	for y, line := range signal {
		for x, char := range line {
			if char != '#' && char != '.' {
				_, exists := signal_map[char]
				if !exists {
					signal_map[char] = make([][2]int, 0)
				}
				signal_map[char] = append(signal_map[char], [2]int{x, y})
			}
		}
	}

	for _, v := range signal_map {
		coords := get_antinodes_coords(v)
		for _, c := range coords {
			set_antinode(c[0], c[1], antinodes_map)
		}
	}

	return antinodes_map
}

func P1() {
	data, _ := os.ReadFile("test.txt")
	signal_map := generate_signal_map(strings.Split(string(data), "\n"))
	antinodes_map := get_antinodes_map(signal_map)
	fmt.Println(antinodes_map)
	result := 0
	for _, line := range antinodes_map {
		fmt.Println(string(line))
		for _, char := range line {
			if char == '#' {
				result++
			}
		}
	}
	fmt.Println(result)
}

func P2() {
	//data, _ := os.ReadFile("test.txt")
	//strings.Split(string(data), "\n")
	//fmt.Println(input)
}
