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

const (
	UP    = 0
	RIGHT = 1
	DOWN  = 2
	LEFT  = 3
)
const directions = 4

func print_room(room []string) {
	for _, v := range room {
		fmt.Println(v)
	}
}
func get_direct_obstacle(obstacles [][][2]int, guard_coords [2]int, guard_direction int) [2]int {
	switch guard_direction {
	case UP:
		for i := guard_coords[1] - 1; i >= 0; i-- {
			for _, c := range obstacles[i] {
				if c[0] == guard_coords[0] {
					return c
				}
			}
		}
		return [2]int{guard_coords[0], -1}
	case DOWN:
		for i := guard_coords[1] + 1; i < len(obstacles); i++ {
			for _, c := range obstacles[i] {
				if c[0] == guard_coords[0] {
					return c
				}
			}
		}
		return [2]int{guard_coords[0], -1}
	case LEFT:
		for i := len(obstacles[guard_coords[1]]) - 1; i >= 0; i-- {
			c := obstacles[guard_coords[1]][i]
			if c[0] < guard_coords[0] {
				return c
			}
		}
		return [2]int{-1, guard_coords[1]}
	default:
		for i := 0; i < len(obstacles[guard_coords[1]]); i++ {
			c := obstacles[guard_coords[1]][i]
			if c[0] > guard_coords[0] {
				return c
			}
		}
		return [2]int{-1, guard_coords[1]}
	}
}

func get_walk_coords(guard_coords [2]int, guard_direction int, next_coord [2]int) [][2]int {
	var walk_coords [][2]int
	switch guard_direction {
	case UP:
		for i := guard_coords[1]; i > next_coord[1]; i-- {
			walk_coords = append(walk_coords, [2]int{guard_coords[0], i})
		}
	case DOWN:
		for i := guard_coords[1]; i < next_coord[1]; i++ {
			walk_coords = append(walk_coords, [2]int{guard_coords[0], i})
		}
	case LEFT:
		for i := guard_coords[0]; i > next_coord[0]; i-- {
			walk_coords = append(walk_coords, [2]int{i, guard_coords[1]})
		}
	default:
		for i := guard_coords[0]; i < next_coord[0]; i++ {
			walk_coords = append(walk_coords, [2]int{i, guard_coords[1]})
		}
	}
	return walk_coords
}

func is_coord_in_bound(coord [2]int, room []string) bool {
	return coord[0] >= 0 && coord[0] < len(room[0]) &&
		coord[1] >= 0 && coord[1] < len(room)
}

func replace_at_index(str string, r rune, i int) string {
	out := []rune(str)
	out[i] = r
	return string(out)
}
func clamp(i int, _min int, _max int) int {
	return min(_max, max(i, _min))
}

func process_guard_route(room []string) []string {
	var obstacles_coords [][][2]int
	loop_detection_map := make(map[[2]int]int)
	var guard_coords [2]int
	var guard_direction = UP
	for y, line := range room {
		var line_obstacles [][2]int
		for x, c := range line {
			if c == '#' {
				line_obstacles = append(line_obstacles, [2]int{x, y})
			}
			if c == '^' {
				guard_coords = [2]int{x, y}
			}
		}
		obstacles_coords = append(obstacles_coords, line_obstacles)
	}
	for {
		next_coord := get_direct_obstacle(obstacles_coords, guard_coords, guard_direction)
		if is_coord_in_bound(next_coord, room) {
			gd, exists := loop_detection_map[next_coord]
			if exists && gd == guard_direction {
				fmt.Println("loop detected")
				panic("")
			}
			loop_detection_map[next_coord] = guard_direction
			walked_coords := get_walk_coords(guard_coords, guard_direction, next_coord)
			fmt.Println(guard_coords)
			fmt.Println(next_coord)
			fmt.Println(guard_direction)
			fmt.Println(walked_coords)
			for _, coord := range walked_coords {
				room[coord[1]] = replace_at_index(room[coord[1]], 'X', coord[0])
			}
			guard_coords[0], guard_coords[1] = walked_coords[len(walked_coords)-1][0], walked_coords[len(walked_coords)-1][1]
			guard_direction = (guard_direction + 1) % directions
		} else {
			if next_coord[0] == -1 && guard_direction == RIGHT {
				next_coord[0] = len(room[0]) - 1
			}
			if next_coord[1] == -1 && guard_direction == DOWN {
				next_coord[1] = len(room) - 1
			}
			walked_coords := get_walk_coords(guard_coords, guard_direction, next_coord)
			fmt.Println(guard_coords)
			fmt.Println(next_coord)
			fmt.Println(guard_direction)
			fmt.Println(walked_coords)
			for _, coord := range walked_coords {
				room[coord[1]] = replace_at_index(room[coord[1]], 'X', coord[0])
			}
			print_room(room)
			break
		}
		print_room(room)
	}
	return room
}

func process_guard_route2(room []string) {
	var starting_guard_coords [2]int
	type Key struct {
		x, y int
	}
	result := 0
	original_room := make([]string, len(room))
	for i := range original_room {
		original_room[i] = strings.Clone(room[i])
	}

	for y := range room {
	outer:
		for x := range room[y] {
			if room[y][x] == '#' || room[y][x] == '^' {
				continue
			}
			for i := range room {
				room[i] = strings.Clone(original_room[i])
			}
			room[y] = replace_at_index(room[y], '#', x)
			loop_detection_map := make(map[[2]int]int)
			guard_direction := UP
			guard_coords := [2]int{starting_guard_coords[0], starting_guard_coords[1]}
			var obstacles_coords [][][2]int
			for y, line := range room {
				var line_obstacles [][2]int
				for x, c := range line {
					if c == '#' {
						line_obstacles = append(line_obstacles, [2]int{x, y})
					}
					if c == '^' {
						guard_coords = [2]int{x, y}
					}
				}
				obstacles_coords = append(obstacles_coords, line_obstacles)
			}
			for {
				next_coord := get_direct_obstacle(obstacles_coords, guard_coords, guard_direction)
				if is_coord_in_bound(next_coord, room) {
					gd, exists := loop_detection_map[next_coord]
					if exists && gd == guard_direction {
						result++
						continue outer
					}
					loop_detection_map[next_coord] = guard_direction
					walked_coords := get_walk_coords(guard_coords, guard_direction, next_coord)
					//fmt.Println(guard_coords)
					//fmt.Println(next_coord)
					//fmt.Println(guard_direction)
					//fmt.Println(walked_coords)
					for _, coord := range walked_coords {
						room[coord[1]] = replace_at_index(room[coord[1]], 'X', coord[0])
					}
					guard_coords[0], guard_coords[1] = walked_coords[len(walked_coords)-1][0], walked_coords[len(walked_coords)-1][1]
					guard_direction = (guard_direction + 1) % directions
				} else {
					//fmt.Println("out of bounds")
					continue outer
				}
				//print_room(room)
			}
		}
	}
	fmt.Println(result)
}

func P1() {
	data, _ := os.ReadFile("test.txt")
	input := strings.Split(string(data), "\n")
	result := 0
	after_patrol_input := process_guard_route(input)
	for _, line := range after_patrol_input {
		for _, c := range line {
			if c == 'X' {
				result++
			}
		}
	}
	fmt.Println(result)
}

func P2() {
	data, _ := os.ReadFile("test.txt")
	input := strings.Split(string(data), "\n")
	process_guard_route2(input)
}
