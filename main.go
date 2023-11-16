package main

import "fmt"

func SudokuSolver(Sudoku [][]int, i int, j int) ([][]int, bool) {

	if j >= 9 {
		return SudokuSolver(Sudoku, i+1, 0)
	}
	if Sudoku[i][j] != 0 {
		if i >= 8 && j >= 8 {
			return Sudoku, true
		}
		return SudokuSolver(Sudoku, i, j+1)
	}
	guess := 1
	for guess < 10 {

		if (!CheckConditionInRow(Sudoku, i, guess)) || (!CheckConditionInColumn(Sudoku, j, guess)) || (!CheckConditionInSquare(Sudoku, i, j, guess)) {
			guess++

		} else {
			Sudoku[i][j] = guess
			if i >= 8 && j >= 8 {
				return Sudoku, true
			}

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
	Sudoku := [][]int{
		{3, 0, 6, 5, 0, 8, 4, 0, 0},
		{5, 2, 0, 0, 0, 0, 0, 0, 0},
		{0, 8, 7, 0, 0, 0, 0, 3, 1},
		{0, 0, 3, 0, 1, 0, 0, 8, 0},
		{9, 0, 0, 8, 6, 3, 0, 0, 5},
		{0, 5, 0, 0, 9, 0, 6, 0, 0},
		{1, 3, 0, 0, 0, 0, 2, 5, 0},
		{0, 0, 0, 0, 0, 0, 0, 7, 4},
		{0, 0, 5, 2, 0, 6, 3, 0, 0}}
	result, ok := SudokuSolver(Sudoku, 0, 0)
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
