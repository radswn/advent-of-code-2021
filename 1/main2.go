package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func swapArray(array []int, newNumber int) {
	array[0] = array[1]
	array[1] = array[2]
	array[2] = newNumber
}

func sumArray(array []int) int {
	result := 0
	for i := 0; i < len(array); i++ {
		result += array[i]
	}
	return result
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.Open("input.txt")
	check(err)
	defer file.Close()
	counter := 0

	lastThree := []int{0, 0, 0}
	previousSum := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		number, err := strconv.Atoi(line)
		check(err)

		if lastThree[0] == 0 {
			lastThree[0] = number
		} else if lastThree[1] == 0 {
			lastThree[1] = number
		} else if lastThree[2] == 0 {
			lastThree[2] = number
			previousSum = sumArray(lastThree)
		} else {
			swapArray(lastThree, number)
			currentSum := sumArray(lastThree)

			if currentSum > previousSum {
				counter++
			}
			previousSum = currentSum
		}
	}

	fmt.Println(counter)
	check(scanner.Err())
}
