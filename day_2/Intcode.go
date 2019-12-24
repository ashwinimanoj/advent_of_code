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
	values[values[index+3]] = values[values[index+1]] + values[values[index+2]]
	return values
}

func multiplyFn(values []int64, index int64) []int64 {
	values[values[3+index]] = values[values[1+index]] * values[values[2+index]]
	return values
}

func manipulate(values []int64, index int64) []int64 {
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
		return values
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
	const expected = 19690720
	var answer int64 = 0
	var noun int64
	var verb int64
	var found bool = false
	for noun = 0; noun <= 99 && !found; noun++ {
		for verb = 0; verb <= 99 && !found; verb++ {
			integerValues := toIntArray(values)
			integerValues[1] = noun
			integerValues[2] = verb
			answer = manipulate(integerValues, 0)[0]
			fmt.Println("answer", answer)
			if answer == expected {
				fmt.Println(100*noun + verb)
				found = true
			}
		}
	}
}
