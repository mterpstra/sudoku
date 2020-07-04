package main

import "testing"

var puzzle PUZZLE

func init() {
	puzzle = [9][9]int{
		{0, 6, 0, 3, 0, 0, 8, 0, 4},
		{5, 3, 7, 0, 9, 0, 0, 0, 0},
		{0, 4, 0, 0, 0, 6, 3, 0, 7},
		{0, 9, 0, 0, 5, 1, 2, 3, 8},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{7, 1, 3, 6, 2, 0, 0, 4, 0},
		{3, 0, 6, 4, 0, 0, 0, 1, 0},
		{0, 0, 0, 0, 6, 0, 5, 2, 3},
		{1, 0, 2, 0, 0, 9, 0, 8, 0},
	}
}

func TestInRow(t *testing.T) {
	if !isValueInRow(&puzzle, 0, 3) {
		t.Errorf("isValueInRow(0,3) failed")
	}
	if isValueInRow(&puzzle, 0, 2) {
		t.Errorf("isValueInRow(0,2) failed")
	}
}

func TestInCol(t *testing.T) {
	if !isValueInCol(&puzzle, 0, 7) {
		t.Errorf("isValueInCol(0,7) failed")
	}
	if isValueInCol(&puzzle, 0, 9) {
		t.Errorf("isValueInCol(0,9) failed")
	}
}

func TestInBox(t *testing.T) {
	values := [][]int{
		{0, 0, 3},
		{6, 0, 2},
		{6, 6, 8},
	}
	for _, v := range values {
		if !isValueInBox(&puzzle, v[0], v[1], v[2]) {
			t.Errorf("isValueInBox failed: %v", v)
		}
	}
	values = [][]int{
		{0, 0, 9},
		{6, 0, 8},
		{6, 6, 4},
	}
	for _, v := range values {
		if isValueInBox(&puzzle, v[0], v[1], v[2]) {
			t.Errorf("isValueInBox failed: %v", v)
		}
	}
}

func TestAllowed(t *testing.T) {
	values := [][]int{
		{0, 0, 2},
		{6, 6, 9},
	}
	for _, v := range values {
		if !isValueAllowed(&puzzle, v[0], v[1], v[2]) {
			t.Errorf("isValueAllowed failed: %v", v)
		}
	}
	values = [][]int{
		{0, 0, 3},
		{0, 0, 1},
		{0, 0, 8},
	}
	for _, v := range values {
		if isValueAllowed(&puzzle, v[0], v[1], v[2]) {
			t.Errorf("isValueAllowed failed: %v", v)
		}
	}
}

func TestEntirePuzzle(t *testing.T) {

	hard := PUZZLE{
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

	solution := PUZZLE{
		{2, 3, 7, 4, 8, 1, 6, 5, 9},
		{6, 9, 4, 5, 2, 7, 1, 3, 8},
		{8, 5, 1, 3, 6, 9, 2, 4, 7},
		{4, 1, 9, 8, 7, 3, 5, 6, 2},
		{7, 8, 5, 6, 1, 2, 4, 9, 3},
		{3, 2, 6, 9, 5, 4, 8, 7, 1},
		{1, 4, 3, 2, 9, 5, 7, 8, 6},
		{5, 7, 8, 1, 3, 6, 9, 2, 4},
		{9, 6, 2, 7, 4, 8, 3, 1, 5},
	}

	solved := solve(hard)
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if solution[i][j] != solved[i][j] {
				t.Errorf("TestEntirePuzzle failed: %v %v", solution, solved)
			}
		}
	}
}
