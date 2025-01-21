package main

import (
	"slices"
	"testing"
)

func TestMaze_FindShPath(t *testing.T) {
	rows := 3
	cols := 3
	maze := NewMaze(rows, cols)

	maze.matrix[0][0] = 1
	maze.matrix[0][1] = 0
	maze.matrix[0][2] = 5
	maze.matrix[1][0] = 1
	maze.matrix[1][1] = 5
	maze.matrix[1][2] = 0
	maze.matrix[2][0] = 9
	maze.matrix[2][1] = 2
	maze.matrix[2][2] = 5

	expected := []Point{{0, 0}, {1, 0}, {1, 1}, {2, 1}}

	result, _ := maze.FindShPath(Point{0, 0}, Point{2, 1})

	if !slices.Equal(expected, result) {
		t.Errorf("Incorrect result. Expected %v, got %v", expected, result)
	}
}

func TestMaze_FindShPathUnreachable(t *testing.T) {
	rows := 3
	cols := 3
	maze := NewMaze(rows, cols)

	maze.matrix[0][0] = 1
	maze.matrix[0][1] = 0
	maze.matrix[0][2] = 5
	maze.matrix[1][0] = 1
	maze.matrix[1][1] = 5
	maze.matrix[1][2] = 0
	maze.matrix[2][0] = 9
	maze.matrix[2][1] = 2
	maze.matrix[2][2] = 5

	var expected []Point

	result, _ := maze.FindShPath(Point{0, 0}, Point{0, 2})

	if !slices.Equal(expected, result) {
		t.Errorf("Incorrect result. Expected %v, got %v", expected, result)
	}
}
