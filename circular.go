package main

import "fmt"

// main creates a square matrix of size (n*n) and fills it in a circular fashion
func main() {
	var gridDim int
	fmt.Print("\nEnter circular matrix size: ")
	fmt.Scan(&gridDim)

	var cellValue int = 1
	var col_begin, col_end, row_begin, row_end int = 0, gridDim - 1, 0, gridDim - 1
	var gridLen int = gridDim * gridDim

	// allocate composed 2d array
	grid := make([][]int, gridDim)
	for i := range grid {
		grid[i] = make([]int, gridDim)
	}

	for cellValue <= gridLen {
		for i := col_begin; i <= col_end; i++ {
			grid[row_begin][i] = cellValue
			cellValue++
		}

		for j := row_begin + 1; j <= row_end; j++ {
			grid[j][col_end] = cellValue
			cellValue++
		}

		for i := col_end - 1; i >= col_begin; i-- {
			grid[row_end][i] = cellValue
			cellValue++
		}

		for j := row_end - 1; j >= row_begin +1; j-- {
			grid[j][col_begin] = cellValue
			cellValue++
		}

		col_begin++
		col_end--
		row_begin++
		row_end--
	}

	/* printing the circular matrix */
	println("\nYour circular matrix is:")
	for i := 0; i < gridDim; i++ {
		for j := 0; j < gridDim; j++ {
			fmt.Printf("%d\t", grid[i][j])
		}
		println()
	}

}