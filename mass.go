package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func divideBy3(num int64) int64 {
	return num / 3
}

func subtractBy2(num int64) int64 {
	return num - 2
}

func findFuel(mass int64) int64 {
	var value int64 = subtractBy2(divideBy3(mass))
	if value <= 0 {
		return 0
	}
	return value + findFuel(value)
}

func main() {
	const filePath = "./mass.txt"
	file, err := os.Open(filePath)
	checkError(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var totalFuel int64 = 0

	fmt.Println("for mass 12:", findFuel(12))
	fmt.Println("for mass 14:", findFuel(14))
	fmt.Println("for mass 14:", findFuel(1969))
	fmt.Println("for mass 100756:", findFuel(100756))

	for scanner.Scan() {
		mass, err := strconv.ParseInt(scanner.Text(), 10, 64)
		checkError(err)
		totalFuel = totalFuel + findFuel(mass)
	}
	fmt.Println("for mass in file:", totalFuel)
}
