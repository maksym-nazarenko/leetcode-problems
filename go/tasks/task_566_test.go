package tasks

import (
	"fmt"
	"testing"
)

func matrixEqual(matrix1, matrix2 [][]int) (bool, error) {
	if len(matrix1) != len(matrix2) {
		return false, fmt.Errorf("Matrices have different number or rows: %d != %d", len(matrix1), len(matrix2))
	}

	if len(matrix1[0]) != len(matrix2[0]) {
		return false, fmt.Errorf("Matrices have different number or columns: %d != %d", len(matrix1[0]), len(matrix2[0]))
	}

	for i, row := range matrix1 {
		for j := range row {
			if matrix1[i][j] != matrix2[i][j] {
				return false, fmt.Errorf("Matrices have different values at [%d,%d]: %d != %d", i, j, matrix1[i][j], matrix2[i][j])
			}
		}
	}

	return true, nil
}

func printMatrix(matrix [][]int) {
	for i := range matrix {
		for j := range matrix[i] {
			fmt.Printf("%3d", matrix[i][j])
		}
		fmt.Println()
	}
}

func TestSuccess(t *testing.T) {
	var testCases = []struct {
		matrixIn  [][]int
		MatrixIn  [][]int
		row       int
		col       int
		matrixOut [][]int
	}{
		{
			matrixIn:  [][]int{{1, 2}, {3, 4}},
			matrixOut: [][]int{{1, 2, 3, 4}},
			row:       1,
			col:       4,
		},
		{
			matrixIn:  [][]int{{1, 2}, {3, 4}},
			matrixOut: [][]int{{1, 2}, {3, 4}},
			row:       2,
			col:       2,
		},
		{
			matrixIn:  [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}},
			matrixOut: [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}, {9, 10}, {11, 12}},
			row:       6,
			col:       2,
		},
		{
			matrixIn:  [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}},
			matrixOut: [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}},
			row:       1,
			col:       2,
		},
	}

	for caseNumber, testCase := range testCases {
		var ret bool
		var err error
		var matrixResult = matrixReshape(testCase.matrixIn, testCase.row, testCase.col)

		ret, err = matrixEqual(testCase.matrixOut, matrixResult)

		fmt.Printf("        Test case #%d\n\n", caseNumber)
		if err != nil {
			t.Error(err)
			printMatrix(testCase.matrixOut)
			fmt.Println()
			printMatrix(matrixResult)
		}

		if !ret {
			t.Errorf("Matrices are not equal")
		}
	}
}
