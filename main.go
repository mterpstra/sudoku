package main

type PUZZLE [9][9]int

func printPuzzle(p *PUZZLE) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			print(p[i][j], " ")
		}
		println("")
	}
}

func findNextEmpty(p *PUZZLE) (row, col int) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if p[i][j] == 0 {
				return i, j
			}
		}
	}
	return 0, 0
}

func isValueInRow(p *PUZZLE, row, val int) bool {
	for i := 0; i < 9; i++ {
		if p[row][i] == val {
			return true
		}
	}
	return false
}

func isValueInCol(p *PUZZLE, col, val int) bool {
	for i := 0; i < 9; i++ {
		if p[i][col] == val {
			return true
		}
	}
	return false
}

func isValueInBox(p *PUZZLE, row, col, val int) bool {

	// Find the top left corner of the small box for this point
	r := (row / 3) * 3
	c := (col / 3) * 3

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if p[r+i][c+j] == val {
				return true
			}
		}
	}
	return false
}

func isValueAllowed(p *PUZZLE, row, col, val int) bool {
	return !isValueInRow(p, row, val) &&
		!isValueInCol(p, col, val) &&
		!isValueInBox(p, row, col, val)
}

func isPuzzleSolved(p PUZZLE) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if p[i][j] == 0 {
				return false
			}
		}
	}
	return true
}

func solve(p PUZZLE) *PUZZLE {
	if isPuzzleSolved(p) {
		return &p
	}
	for i := 1; i < 10; i++ {
		r, c := findNextEmpty(&p)
		if isValueAllowed(&p, r, c, i) {
			p[r][c] = i
			if solution := solve(p); solution != nil {
				return solution
			}
			p[r][c] = 0
		}
	}
	return nil
}

// Puzzle setup is sent like this.
// puzzle (row,col,val), (row,col,val), (row,col,val)`
func main() {

	puzzle := PUZZLE{
		{0, 3, 0, 4, 8, 0, 6, 0, 9},
		{0, 0, 0, 0, 2, 7, 0, 0, 0},
		{8, 0, 0, 3, 0, 0, 0, 0, 0},
		{0, 1, 9, 0, 0, 0, 0, 0, 0},
		{7, 8, 0, 0, 0, 2, 0, 9, 3},
		{0, 0, 0, 0, 0, 4, 8, 7, 0},
		{0, 0, 0, 0, 0, 5, 0, 0, 6},
		{0, 0, 0, 1, 3, 0, 0, 0, 0},
		{9, 0, 2, 0, 4, 8, 0, 1, 0},
	}

	println("Start of Puzzle")
	printPuzzle(&puzzle)

	println("Solving...")
	solved := solve(puzzle)

	println("Done.")
	printPuzzle(solved)
}
