package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y int
}

type Maze struct {
	matrix          [][]int
	rowsNum, colNum int
}

func NewMaze(n, m int) *Maze {
	maze := Maze{matrix: make([][]int, n), rowsNum: n, colNum: m}
	for i := range maze.matrix {
		maze.matrix[i] = make([]int, m)
	}

	return &maze
}

func (m *Maze) IsZero(p Point) bool {
	if m.matrix[p.X][p.Y] == 0 {
		return true
	}
	return false
}

func (m *Maze) GetValue(p Point) int {
	return m.matrix[p.X][p.Y]
}

func (m *Maze) FindShPath(s, f Point) ([]Point, error) {
	dist := make(map[Point]int)
	dist[s] = 0

	visited := NewSet[Point]()

	component := NewSet[Point]()
	component.Add(s)

	paths := make(map[Point][]Point)
	paths[s] = []Point{s}

	for visited.Len() != component.Len() {
		var p Point
		minDist := math.MaxInt
		for k, v := range dist {
			if v < minDist && !visited.Contains(k) {
				p = k
			}
		}

		steps := m.getStrategy(p)

		for _, step := range steps {
			if visited.Contains(step) || m.IsZero(step) {
				continue
			}

			component.Add(step)

			weight := dist[p] + m.GetValue(step)
			if v, ok := dist[step]; !ok {
				dist[step] = weight
				paths[step] = append(paths[p], step)
			} else if v > weight {
				dist[step] = weight
				paths[step] = append(paths[p], step)
			}
		}

		visited.Add(p)
	}

	if _, ok := dist[f]; ok {
		return paths[f], nil
	} else {
		return []Point{}, fmt.Errorf("unable to reach finish %v", f)
	}
}

func (m *Maze) getStrategy(p Point) []Point {
	neighbours := make([]Point, 0)
	r := m.rowsNum
	c := m.colNum

	switch {
	case p.X > 0 && p.Y > 0 && p.X < (r-1) && p.Y < (c-1):
		neighbours = append(neighbours, Point{p.X - 1, p.Y}, Point{p.X + 1, p.Y},
			Point{p.X, p.Y - 1}, Point{p.X, p.Y + 1})

	case p.X == 0 && p.Y > 0 && p.Y < (c-1):
		neighbours = append(neighbours, Point{p.X + 1, p.Y}, Point{p.X, p.Y - 1}, Point{p.X, p.Y + 1})
	case p.X == 0 && p.Y == 0:
		neighbours = append(neighbours, Point{p.X + 1, p.Y}, Point{p.X, p.Y + 1})
	case p.X == 0 && p.Y == (c-1):
		neighbours = append(neighbours, Point{p.X + 1, p.Y}, Point{p.X, p.Y - 1})

	case p.X == (r-1) && p.Y > 0 && p.Y < (c-1):
		neighbours = append(neighbours, Point{p.X - 1, p.Y}, Point{p.X, p.Y - 1}, Point{p.X, p.Y + 1})
	case p.X == (r-1) && p.Y == 0:
		neighbours = append(neighbours, Point{p.X - 1, p.Y}, Point{p.X, p.Y + 1})
	case p.X == (r-1) && p.Y == (c-1):
		neighbours = append(neighbours, Point{p.X - 1, p.Y}, Point{p.X, p.Y - 1})

	case p.X > 0 && p.X < (r-1) && p.Y == 0:
		neighbours = append(neighbours, Point{p.X - 1, p.Y}, Point{p.X + 1, p.Y}, Point{p.X, p.Y + 1})
	case p.X > 0 && p.X < (r-1) && p.Y == (c-1):
		neighbours = append(neighbours, Point{p.X - 1, p.Y}, Point{p.X + 1, p.Y}, Point{p.X, p.Y - 1})
	}

	return neighbours
}
