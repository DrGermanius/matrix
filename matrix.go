package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var (
	errMatrixConsistsNonIntegerElems = errors.New("only integers allowed im matrix")
)

// isMatrixSquare checks if the number of elements in each row is equal to the number of rows
func isMatrixSquare(matrix [][]string) bool {
	numRows := len(matrix)
	for _, row := range matrix {
		if len(row) != numRows {
			return false
		}
	}
	return true
}

// matrixToString represents matrix as string
func matrixToString(matrix [][]string) string {
	var result string
	for _, row := range matrix {
		result = fmt.Sprintf("%s%s\n", result, strings.Join(row, ","))
	}
	return result
}

// matrixToFlatString converts matrix to flat representation
func matrixToFlatString(matrix [][]string) string {
	var strs []string
	for i := range matrix {
		strs = append(strs, strings.Join(matrix[i], ","))
	}
	return strings.Join(strs, ",")
}

// sumIntMatrix gets the sum of int matrix elements
func sumIntMatrix(matrix [][]string) (int, error) {
	var total int
	intMatrix, err := stringMatrixToInt(matrix)
	if err != nil {
		return 0, err
	}
	for i := range intMatrix {
		for j := range intMatrix[i] {
			total += intMatrix[i][j]
		}
	}
	return total, nil
}

// multiplyIntMatrix gets the product of int matrix elements
func multiplyIntMatrix(matrix [][]string) (int, error) {
	total := 1 // in case of multiplying the initial value should be 1
	intMatrix, err := stringMatrixToInt(matrix)
	if err != nil {
		return 0, err
	}
	for i := range intMatrix {
		for j := range intMatrix[i] {
			total *= intMatrix[i][j]
		}
	}
	return total, nil
}

// stringMatrixToInt converts string matrix to int matrix
func stringMatrixToInt(matrix [][]string) ([][]int, error) {
	n := len(matrix)
	m := len(matrix[0])
	res := make([][]int, m)
	for i := range res {
		res[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			elem, err := strconv.Atoi(matrix[i][j])
			if err != nil {
				return nil, errMatrixConsistsNonIntegerElems
			}
			res[i][j] = elem
		}
	}
	return res, nil
}

// invertMatrix Inverts the rows and columns of the matrix
func invertMatrix(matrix [][]string) [][]string {
	n := len(matrix)
	m := len(matrix[0])
	inverted := make([][]string, m)
	for i := range inverted {
		inverted[i] = make([]string, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			inverted[j][i] = matrix[i][j]
		}
	}
	return inverted
}
