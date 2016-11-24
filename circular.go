package main

import "fmt"

func main() {
	var gridDim int
	fmt.Print("Enter circular array size: ")
	fmt.Scan(&gridDim)

	var cellValue int = 1
	var col1, col2, row1, row2 int = 0, gridDim - 1, 0, gridDim - 1
	var gridLen int = gridDim * gridDim

	// allocate composed 2d array
	grid := make([][]int, gridDim)
	for i := range grid {
		grid[i] = make([]int, gridDim)
	}

	for cellValue <= gridLen {
		for i := col1; i <= col2; i++ {
			grid[row1][i] = cellValue
			cellValue++
		}

		for j := row1 + 1; j <= row2; j++ {
			grid[j][col2] = cellValue
			cellValue++
		}

		for i := col2 - 1; i >= col1; i-- {
			grid[row2][i] = cellValue
			cellValue++
		}

		for j := row2 - 1; j >= row1+1; j-- {
			grid[j][col1] = cellValue
			cellValue++
		}

		col1++
		col2--
		row1++
		row2--
	}

	/* printing the circular matrix */
	println("Here is your circular matrix is:")
	for i := 0; i < gridDim; i++ {
		for j := 0; j < gridDim; j++ {
			fmt.Printf("%d\t", grid[i][j])
		}
		println()
	}

}