package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	// [numRows][numCols]
	// var matrix1 = [3][4]int {
	// 	{0, 1, 2, 3},
	// 	{4, 5, 6, 7},
	// 	{8, 9, 10, 11},
	// }

	// var matrix2 = [4][3]int {
	// 	{0, 1, 2},
	// 	{3, 4, 5},
	// 	{6, 7, 8},
	// 	{9, 10, 11},
	// }

	var matrix1 = [1][3]int {
		{3, 4, 2},
	}
	var matrix2 = [3][4]int {
		{13, 9, 7, 15},
		{8, 7, 4, 6},
		{6, 4, 0, 3},
	}

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