package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	check(err)
	defer file.Close()
	counter := 0
	last := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		number, err := strconv.Atoi(line)
		check(err)

		if number > last && last != 0 {
			counter++
		}
		last = number
	}

	fmt.Println(counter)
	check(scanner.Err())
}
