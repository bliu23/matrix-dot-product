package main

import (
	"fmt"
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func sliceAtoi(stringNum []string) ([]int, error) {
	intNum := make([]int, 0, len(stringNum))
    for _, a := range stringNum {
        i, err := strconv.Atoi(a)
        if err != nil {
            return intNum, err
        }
        intNum = append(intNum, i)
    }
    return intNum, nil
}

func readMatrix(path string) [][]int {
	resultMatrix := make([][]int, 0)
	
	file, err := os.Open(path)
	if err != nil {
        log.Fatal(err)
    }
	defer file.Close()
	
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Fields(line)
		intArray, err := sliceAtoi(nums)
		if err != nil {
			fmt.Println(err)
		}
		resultMatrix = append(resultMatrix, intArray)
	}

	return resultMatrix
}

func main() {
	var matrix1 = readMatrix("matrix1.txt")
	var matrix2 = readMatrix("matrix2.txt")

	numRows1 := len(matrix1)
	numCols1 := len(matrix1[0])
	numRows2 := len(matrix2)
	numCols2 := len(matrix2[0])

	if numCols1 != numRows2 {
		fmt.Println("Error!")
	}

	// var resultMatrix = [numRows1][numCols2]int
	resultMatrix := make([][]int, numRows1)
    for i := range resultMatrix {
        resultMatrix[i] = make([]int, numCols2)
	}
	
	for resultRow := 0; resultRow < numRows1; resultRow++ {
		for resultCol := 0; resultCol < numCols2; resultCol++ {
			entryTotal := 0
			for i := 0; i < numCols1; i++ {
				entryTotal += (matrix1[resultRow][i] * matrix2[i][resultCol])
			}
			resultMatrix[resultRow][resultCol] = entryTotal
		}
	}

	fmt.Println(resultMatrix)

}





// Using golang version 1.8 or higher, write a program that calculates and outputs the dot product of two integer matrices. The two matrices should be stored in files in a format that you define. The program should take as input parameters the locations of the two files that represent the two matrices. Please document and provide examples of the matrix file format as well as the source code for the program.