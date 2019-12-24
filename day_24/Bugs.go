package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

var bugSymbol = "#"
var emptySpaceSymbol = "."

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func hasLeftNeighbor(i int, j int, grid [5][5]string) bool {
	return i-1 >= 0 && grid[i-1][j] == bugSymbol
}

func hasRightNeighbor(i int, j int, grid [5][5]string) bool {
	return i+1 < 5 && grid[i+1][j] == bugSymbol
}

func hasTopNeighbor(i int, j int, grid [5][5]string) bool {
	return j-1 >= 0 && grid[i][j-1] == bugSymbol
}

func hasBottomNeighbor(i int, j int, grid [5][5]string) bool {
	return j+1 < 5 && grid[i][j+1] == bugSymbol
}

func adjacentBugs(i int, j int, grid [5][5]string) int {
	var bugs = 0
	if hasLeftNeighbor(i, j, grid) {
		bugs++
	}
	if hasRightNeighbor(i, j, grid) {
		bugs++
	}
	if hasTopNeighbor(i, j, grid) {
		bugs++
	}
	if hasBottomNeighbor(i, j, grid) {
		bugs++
	}
	return bugs
}

func minuteChange(grid [5][5]string) [5][5]string {
	var newGrid [5][5]string
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			adjacentBugs := adjacentBugs(i, j, grid)
			if grid[i][j] == bugSymbol && adjacentBugs != 1 {
				newGrid[i][j] = emptySpaceSymbol
			} else if grid[i][j] == emptySpaceSymbol && (adjacentBugs == 1 || adjacentBugs == 2) {
				newGrid[i][j] = bugSymbol
			} else {
				newGrid[i][j] = grid[i][j]
			}
		}
	}
	return newGrid
}

func printGrid(grid [5][5]string) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Print(grid[i][j])
		}
		fmt.Println()
	}
}

func calculateBiodiversity(grid [5][5]string) int {
	var total int = 0
	var index float64 = 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if grid[i][j] == bugSymbol {
				total += int(math.Pow(2, index))
			}
			index++
		}
	}
	return int(total)
}

func isFound(totalResults [][5][5]string, newGrid [5][5]string) bool {
	found := false
	for _, grid := range totalResults {
		if grid == newGrid {
			found = true
			break
		}
	}
	fmt.Println(found, "----found----")
	return found
}

func main() {
	const filePath = "./bugs.txt"
	file, err := os.Open(filePath)
	checkError(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var grid [5][5]string

	var i = 0
	for scanner.Scan() {
		input := strings.Split(scanner.Text(), "")
		for j := 0; j < 5; j++ {
			grid[i][j] = input[j]
		}
		i++
	}

	var totalResults [][5][5]string

	for k := 0; k < 90; k++ {
		var newGrid [5][5]string
		newGrid = minuteChange(grid)
		if isFound(totalResults, newGrid) {
			printGrid(newGrid)
			fmt.Println("BIODIVERSITY --------- : ", calculateBiodiversity(newGrid))
			break
		}
		fmt.Println("-----After ", k+1, " minute ------")
		printGrid(newGrid)
		totalResults = append(totalResults, newGrid)
		grid = newGrid
	}
}
