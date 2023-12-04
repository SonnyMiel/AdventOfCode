package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func solution(path string) int {
	result := 0

	file, err := os.Open(path)

	if err != nil {
		fmt.Println("Error:", err)
		return -1
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// fmt.Println(line)
		runes := []rune(line)

		for first, last := 0, len(runes)-1; first != last; {
			if !unicode.IsDigit(runes[first]) {
				first++
			}
			if !unicode.IsDigit(runes[last]) {
				last--
			}
			if unicode.IsDigit(runes[first]) && unicode.IsDigit(runes[last]) {
				number, err := strconv.Atoi(fmt.Sprintf("%c%c", runes[first], runes[last]))
				if err != nil {
					fmt.Println("Error:", err)
					return -1
				}
        result += number
				break
			}
		}
	}

	return result
}

func main() {
	fmt.Println("Exercise 1")

	result := solution("./input.txt")

	fmt.Println("result: ", result)
}
