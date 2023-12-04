package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
		line := strings.Split(scanner.Text(), ":");

    game_id := line[0]
       


	}

	return result
}

func main() {
	fmt.Println("Exercise 1")

	result := solution("./input.txt")

	fmt.Println("result: ", result)
}
