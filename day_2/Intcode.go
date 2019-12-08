package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func addfn(values []int64, index int64) []int64 {
	fmt.Println(values)
	values[values[index+3]] = values[values[index+1]] + values[values[index+2]]
	fmt.Println(values)
	return values
}

func multiplyFn(values []int64, index int64) []int64 {
	values[values[3+index]] = values[values[1+index]] * values[values[2+index]]
	return values
}

func manipulate(values []int64, index int64) []int64 {
	fmt.Println(values, index)
	var currentOpcode = values[index]
	switch currentOpcode {
	case 1:
		values = addfn(values, index)
		manipulate(values, index+4)
	case 2:
		values = multiplyFn(values, index)
		manipulate(values, index+4)
	case 99:
		return values
	default:
		fmt.Println("invalid")
	}
	return values
}

func toIntArray(arrayOfStrings []string) []int64 {
	var arrayOfIntegers []int64
	for i := 0; i < len(arrayOfStrings); i++ {
		var value, err = strconv.ParseInt(arrayOfStrings[i], 10, 64)
		checkError(err)
		arrayOfIntegers = append(arrayOfIntegers, value)
	}
	return arrayOfIntegers
}

func main() {
	const filePath = "./Intcode.txt"
	file, err := os.Open(filePath)
	checkError(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	values := strings.Split(scanner.Text(), ",")
	integerValues := toIntArray(values)
	integerValues[1] = 12
	integerValues[2] = 2
	fmt.Println(manipulate(integerValues, 0)[0])

}
