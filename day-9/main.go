package main

import (
	"fmt"
	"os"
	"strconv"
)

type file struct {
	start_index int
	size        int
}

func main() {
	data, _ := os.ReadFile("./test.txt")
	disk_map := string(data)

	var disk []int
	file_number := 0
	var files []file
	for i, item := range disk_map {
		size, _ := strconv.Atoi(string(item))
		if i%2 == 0 {
			start_index := len(disk)
			for i := 0; i < size; i++ {
				disk = append(disk, file_number)
			}
			files = append(files, file{start_index: start_index, size: size})
			file_number++
		} else {
			for i := 0; i < size; i++ {
				disk = append(disk, -1)
			}
		}
	}

	for i := len(files) - 1; i >= 0; i-- {
		defrag(i, files[i], disk)
	}

	fmt.Println(checksum(disk))
}

func checksum(numbers []int) int64 {
	sum := int64(0)
	for i, num := range numbers {
		if num == -1 {
			continue
		}

		sum += int64(num) * int64(i)
	}
	return sum
}

func defrag(file_number int, file file, numbers []int) {
	pointer := 0
	for pointer < file.start_index {
		for pointer < file.start_index && numbers[pointer] != -1 {
			pointer++
		}

		if pointer == file.start_index {
			break
		}

		cnt_dots := 0
		start_dots := pointer
		for pointer < file.start_index && numbers[pointer] == -1 {
			cnt_dots++
			pointer++
		}

		if cnt_dots >= file.size {
			move_file(file_number, file, numbers, start_dots)
			break
		}
	}
}

func move_file(file_number int, file file, numbers []int, start_dots int) {
	for i := 0; i < file.size; i++ {
		numbers[start_dots+i] = file_number
		numbers[file.start_index+i] = -1
	}
}
