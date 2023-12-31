package main

import "fmt"

func SudokuSolver(Sudoku [][]int, i int, j int) ([][]int, bool) {
	if i >= 8 && j >= 9 {
		return Sudoku, true
	}
	if j >= 9 {
		return SudokuSolver(Sudoku, i+1, 0)
	}
	if Sudoku[i][j] != 0 {

		return SudokuSolver(Sudoku, i, j+1)
	}
	guess := 1
	for guess < 10 {

		if (!CheckConditionInRow(Sudoku, i, guess)) || (!CheckConditionInColumn(Sudoku, j, guess)) || (!CheckConditionInSquare(Sudoku, i, j, guess)) {
			guess++

		} else {
			Sudoku[i][j] = guess

			if su, temp := SudokuSolver(Sudoku, i, j+1); !temp {
				Sudoku[i][j] = 0
				guess++
			} else {
				return su, temp
			}
		}

	}

	return Sudoku, false
}
func CheckConditionInRow(Sudoku [][]int, Row int, guess int) bool {
	for i := 0; i < 9; i++ {
		if Sudoku[Row][i] == guess {
			return false
		}
	}
	return true
}
func CheckConditionInColumn(Sudoku [][]int, Column int, guess int) bool {
	for i := 0; i < 9; i++ {
		if Sudoku[i][Column] == guess {
			return false
		}
	}
	return true
}
func CheckConditionInSquare(Sudoku [][]int, Row int, Column int, guess int) bool {

	r := Row / 3
	c := Column / 3
	for i := r * 3; i < r*3+3; i++ {
		for j := c * 3; j < c*3+3; j++ {
			if Sudoku[i][j] == guess {
				return false
			}
		}
	}
	return true
}
func main() {
	Sudoku := [][][]int{
		{
			{0, 3, 0, 5, 8, 0, 0, 9, 0},
			{4, 8, 0, 9, 0, 0, 7, 0, 0},
			{0, 0, 0, 0, 0, 4, 0, 0, 8},
			{0, 0, 3, 0, 0, 0, 5, 0, 7},
			{0, 0, 8, 1, 5, 7, 3, 0, 0},
			{7, 0, 5, 0, 0, 0, 6, 0, 0},
			{2, 0, 0, 8, 0, 0, 0, 0, 0},
			{0, 0, 4, 0, 0, 5, 0, 6, 2},
			{0, 5, 0, 0, 4, 1, 0, 7, 0},
		},

		{

			{9, 0, 5, 6, 3, 0, 0, 0, 8},
			{0, 4, 0, 8, 9, 0, 0, 0, 0},
			{0, 6, 0, 0, 0, 1, 3, 0, 0},
			{4, 1, 0, 0, 0, 0, 0, 0, 0},
			{6, 0, 3, 0, 0, 0, 2, 0, 5},
			{0, 0, 0, 0, 0, 0, 0, 8, 1},
			{0, 0, 4, 1, 0, 0, 0, 9, 0},
			{0, 0, 0, 0, 4, 9, 0, 2, 0},
			{8, 0, 0, 0, 2, 6, 1, 0, 4},
		},

		{
			{2, 0, 4, 0, 7, 9, 5, 0, 0},
			{0, 0, 0, 0, 5, 0, 0, 0, 6},
			{0, 6, 0, 0, 2, 3, 1, 0, 0},
			{0, 0, 3, 0, 0, 0, 0, 4, 0},
			{6, 0, 5, 0, 0, 0, 9, 0, 1},
			{0, 2, 0, 0, 0, 0, 3, 0, 0},
			{0, 0, 6, 5, 9, 0, 0, 1, 0},
			{1, 0, 0, 0, 4, 0, 0, 0, 0},
			{0, 0, 8, 3, 1, 0, 2, 0, 7},
		},

		{
			{5, 0, 0, 1, 0, 0, 6, 0, 0},
			{3, 7, 0, 0, 0, 9, 0, 0, 4},
			{9, 1, 0, 0, 0, 4, 0, 5, 0},
			{0, 0, 7, 9, 0, 0, 0, 0, 0},
			{0, 0, 9, 7, 0, 6, 8, 0, 0},
			{0, 0, 0, 0, 0, 2, 4, 0, 0},
			{0, 2, 0, 8, 0, 0, 0, 4, 5},
			{6, 0, 0, 2, 0, 0, 0, 1, 8},
			{0, 0, 1, 0, 0, 5, 0, 0, 2},
		},

		{
			{4, 0, 6, 0, 2, 0, 0, 0, 9},
			{0, 5, 2, 1, 0, 0, 0, 0, 0},
			{0, 0, 9, 0, 0, 4, 8, 0, 0},
			{0, 4, 0, 9, 8, 7, 0, 0, 0},
			{9, 0, 0, 0, 0, 0, 0, 0, 4},
			{0, 0, 0, 5, 4, 2, 0, 6, 0},
			{0, 0, 8, 4, 0, 0, 6, 0, 0},
			{0, 0, 0, 0, 0, 9, 1, 2, 0},
			{3, 0, 0, 0, 7, 0, 4, 0, 8},
		}, {
			{8, 0, 5, 4, 2, 1, 0, 0, 0},
			{0, 7, 0, 0, 9, 0, 0, 0, 0},
			{0, 9, 0, 0, 3, 0, 0, 0, 5},
			{0, 4, 6, 2, 0, 0, 0, 0, 0},
			{0, 5, 9, 0, 0, 0, 2, 6, 0},
			{0, 0, 0, 0, 0, 9, 3, 5, 0},
			{5, 0, 0, 0, 8, 0, 0, 7, 0},
			{0, 0, 0, 0, 4, 0, 0, 3, 0},
			{0, 0, 0, 9, 1, 5, 4, 0, 6},
		},

		{

			{0, 1, 0, 0, 7, 8, 2, 6, 0},
			{0, 0, 6, 0, 0, 0, 0, 3, 5},
			{0, 0, 0, 0, 0, 6, 8, 0, 1},
			{6, 0, 0, 0, 0, 5, 0, 0, 0},
			{0, 0, 4, 2, 0, 3, 1, 0, 0},
			{0, 0, 0, 7, 0, 0, 0, 0, 2},
			{4, 0, 5, 6, 0, 0, 0, 0, 0},
			{3, 6, 0, 0, 0, 0, 9, 0, 0},
			{0, 8, 9, 1, 3, 0, 0, 4, 0},
		},
		{
			{0, 0, 8, 0, 0, 0, 5, 0, 0},
			{6, 0, 0, 7, 0, 5, 0, 0, 3},
			{0, 9, 0, 8, 3, 2, 0, 0, 0},
			{0, 0, 4, 0, 1, 0, 0, 0, 0},
			{3, 8, 0, 4, 0, 7, 0, 5, 1},
			{0, 0, 0, 0, 8, 0, 2, 0, 0},
			{0, 0, 0, 1, 5, 9, 0, 7, 0},
			{8, 0, 0, 3, 0, 4, 0, 0, 5},
			{0, 0, 9, 0, 0, 0, 1, 0, 0},
		},
		{
			{0, 0, 9, 2, 0, 0, 6, 0, 0},
			{5, 1, 8, 0, 0, 4, 0, 0, 0},
			{0, 0, 0, 0, 3, 0, 0, 0, 5},
			{0, 2, 7, 0, 0, 5, 0, 0, 8},
			{0, 8, 5, 0, 1, 0, 9, 4, 0},
			{1, 0, 0, 6, 0, 0, 5, 2, 0},
			{4, 0, 0, 0, 9, 0, 0, 0, 0},
			{0, 0, 0, 5, 0, 0, 4, 3, 9},
			{0, 0, 3, 0, 0, 6, 7, 0, 0}},
		{
			{3, 0, 6, 5, 0, 8, 4, 0, 0},
			{5, 2, 0, 0, 0, 0, 0, 0, 0},
			{0, 8, 7, 0, 0, 0, 0, 3, 1},
			{0, 0, 3, 0, 1, 0, 0, 8, 0},
			{9, 0, 0, 8, 6, 3, 0, 0, 5},
			{0, 5, 0, 0, 9, 0, 6, 0, 0},
			{1, 3, 0, 0, 0, 0, 2, 5, 0},
			{0, 0, 0, 0, 0, 0, 0, 7, 4},
			{0, 0, 5, 2, 0, 6, 3, 0, 0}}}

	result, ok := SudokuSolver(Sudoku[2], 0, 0)
	if !ok {
		fmt.Println("Sorry, I could not solve it")
	} else {
		fmt.Println("Sudoku solved:")
		for _, i2 := range result {
			for _, i3 := range i2 {
				fmt.Printf("%v ", i3)
			}
			fmt.Println(" ")
		}
	}

}
