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
	for x := range grid {
		grid[x] = make([]int, gridDim)
	}

	for cellValue <= gridLen {
		for c := col_begin; c <= col_end; c++ {
			grid[row_begin][c] = cellValue
			cellValue++
		}

		for r := row_begin + 1; r <= row_end; r++ {
			grid[r][col_end] = cellValue
			cellValue++
		}

		for c := col_end - 1; c >= col_begin; c-- {
			grid[row_end][c] = cellValue
			cellValue++
		}

		for r := row_end - 1; r >= row_begin +1; r-- {
			grid[r][col_begin] = cellValue
			cellValue++
		}

		col_begin++
		col_end--
		row_begin++
		row_end--
	}

	// printing the circular matrix
	println("\nYour circular matrix is:")
	for i := 0; i < gridDim; i++ {
		for j := 0; j < gridDim; j++ {
			fmt.Printf("%d\t", grid[i][j])
		}
		println()
	}

}