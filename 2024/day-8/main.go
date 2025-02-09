package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	//P1()
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

func is_antinode_in_range(x int, y int, _map [][]rune) bool {
	if y < 0 || y >= len(_map) {
		return false
	}
	if x < 0 || x >= len(_map[0]) {
		return false
	}
	return true
}

func set_antinode(x int, y int, antinodes_map [][]rune) bool {
	if is_antinode_in_range(x, y, antinodes_map) {
		antinodes_map[y][x] = '#'
		return true
	}
	return false
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
			if set_antinode(c[0], c[1], antinodes_map) {
			}
		}
	}

	return antinodes_map
}

func generate_the_fucking_nodes(origin [2]int, point [2]int, coords *[][2]int, signal [][]rune) {
	if origin[0] == point[0] && origin[1] == point[1] {
		return
	}
	x := point[0] + (point[0] - origin[0])
	y := point[1] + (point[1] - origin[1])
	if is_antinode_in_range(x, y, signal) {
		*coords = append(*coords, [2]int{x, y})
		generate_the_fucking_nodes(point, [2]int{x, y}, coords, signal)
	}
}
func get_antinodes_map2(signal [][]rune) [][]rune {
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
		coords := make([][2]int, 0)
		for i := range v {
			for j := range v {
				if i != j {
					continue
				}
				x := v[i][0] + (v[i][0] - v[j][0])
				y := v[i][1] + (v[i][1] - v[j][1])
				if is_antinode_in_range(x, y, signal) {
					coords = append(coords, [2]int{x, y})
					generate_the_fucking_nodes(v[i], [2]int{x, y}, &coords, signal)
					generate_the_fucking_nodes(v[j], [2]int{x, y}, &coords, signal)
				}
			}
		}
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
	result := 0
	for _, line := range antinodes_map {
		for _, char := range line {
			if char == '#' {
				result++
			}
		}
	}

	fmt.Println(result)
}

func P2() {
	data, _ := os.ReadFile("test.txt")
	signal_map := generate_signal_map(strings.Split(string(data), "\n"))
	antinodes_map := get_antinodes_map2(signal_map)
	result := 0
	for _, line := range antinodes_map {
		fmt.Println(string(line))
		for _, char := range line {
			if char == '#' {
				result++
			}
		}
	}
	for _, line := range signal_map {
		for _, char := range line {
			if char != '.' {
				result++
			}
		}
	}
	fmt.Println(result)
}
