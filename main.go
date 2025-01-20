package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func (m *Maze) Zero(p Point) bool {
	if m.matrix[p.X][p.Y] == 0 {
		return true
	}
	return false
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, m int
	_, err := fmt.Fscanln(reader, &n, &m)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to parse maze size: %v\n", err)
	}

	maze := NewMaze(n, m)

	for i := 0; i < n; i++ {
		str, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to parse maze on line %d: %v\n", i, err)
		}

		str = strings.TrimSpace(str)
		split := strings.Split(str, " ")

		if len(split) != m {
			fmt.Fprintf(os.Stderr, "Maze line length %d and number of columns %d does not match\n", len(split), m)
		}

		for j, strNum := range split {
			num, err := strconv.Atoi(strNum)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Failed to parse symbol %d of line %d in maze: %v\n", j, i, err)
			}

			maze.matrix[i][j] = num
		}
	}

	var start, finish Point
	_, err = fmt.Fscanln(reader, &start.X, &start.Y, &finish.X, &finish.Y)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to parse start and finish: %v\n", err)
	}

	fmt.Println()

	for _, line := range maze.matrix {
		for _, cell := range line {
			fmt.Printf("%d ", cell)
		}
		fmt.Println()
	}

	fmt.Println()
	fmt.Println(start)
	fmt.Println(finish)

}

func (m *Maze) FindShPath(s, f Point) {

	q := NewQueue[Point]()
	q.Push(s)

	visited := NewSet[Point]()

	for q.Len() > 0 {
		p := q.Pop()
		steps := m.getStrategy(p)

		for _, step := range steps {
			if visited.Contains(step) || m.Zero(step) {
				continue
			}
		}
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
