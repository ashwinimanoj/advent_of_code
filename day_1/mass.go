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

func findFuel(num int64) int64 {
	return subtractBy2(divideBy3(num))
}

func main() {
	const filePath = "./mass.txt"
	file, err := os.Open(filePath)
	checkError(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var totalFuel int64 = 0

	fmt.Println(findFuel(12))
	fmt.Println(findFuel(14))
	fmt.Println(findFuel(1969))
	fmt.Println(findFuel(100756))

	for scanner.Scan() {
		num, err := strconv.ParseInt(scanner.Text(), 10, 64)
		checkError(err)
		totalFuel = totalFuel + findFuel(num)
	}
	fmt.Println(totalFuel)
}
