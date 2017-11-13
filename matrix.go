// Using golang version 1.8 or higher, write a program that calculates and outputs the dot product of two integer matrices. The two matrices should be stored in files in a format that you define. The program should take as input parameters the locations of the two files that represent the two matrices. Please document and provide examples of the matrix file format as well as the source code for the program.


/*

** Running the program **
	Given two files matrix1.txt and matrix2.txt that contain the matrices you wish to multiply,
	you can run the program like so:
		`go run matrix.go matrix1.txt matrix2.txt`
	
	This will multiply the two matrices together and outputs to stdout the dot product as a 2D array of integers.


** Matrix file format **
	Each row contain numbers separated by a single whitespace.
	Each newline constitutes a new row.

	For example a matrix with two rows and three columns would look like this:
	0 1 2
	3 4 5
	
*/

package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

// Helper function to convert string array to int array
func sliceAtoi(stringNum []string) ([]int, error) {
	intNum := make([]int, 0, len(stringNum))
	for _, num := range stringNum {
		i, err := strconv.Atoi(num)
		if err != nil {
			return intNum, err
		}
		intNum = append(intNum, i)
	}
	return intNum, nil
}

// Read a matrix given the specified format, return as 2D array of type int.
func readMatrix(path string) ([][]int, error) {
	resultMatrix := make([][]int, 0)
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file.")
		return nil, err
	}
	defer file.Close()
	
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Convert line into string array of nums, and then into int array of nums
		line := scanner.Text()
		numsAsString := strings.Fields(line)
		numsAsInt, err := sliceAtoi(numsAsString)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		resultMatrix = append(resultMatrix, numsAsInt)
	}

	return resultMatrix, nil
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Requires two matrices as input.")
		return
	}
	matrix1, err := readMatrix(os.Args[1])
	if err != nil {
		fmt.Println("Error opening file.")
		return
	}
	matrix2, err := readMatrix(os.Args[2])
	if err != nil {
		fmt.Println("Error opening file.")
		return
	}

	numRows1 := len(matrix1)
	numCols1 := len(matrix1[0])
	numRows2 := len(matrix2)
	numCols2 := len(matrix2[0])

	validDotProduct := numCols1 == numRows2
	if !validDotProduct {
		fmt.Println("Number of columns in matrix 1 must equal number of rows in matrix 2.")
		return
	}

	resultMatrix := make([][]int, numRows1)
	for i := range resultMatrix {
		resultMatrix[i] = make([]int, numCols2)
	}
	
	// Calculate dot product
	for resultRow := 0; resultRow < numRows1; resultRow++ {
		for resultCol := 0; resultCol < numCols2; resultCol++ {
			entryTotal := 0
			// Run the main calculation
			for i := 0; i < numCols1; i++ {
				entryTotal += (matrix1[resultRow][i] * matrix2[i][resultCol])
			}
			resultMatrix[resultRow][resultCol] = entryTotal
		}
	}

	fmt.Println(resultMatrix)
}
