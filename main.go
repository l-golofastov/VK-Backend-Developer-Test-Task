package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, m int
	_, err := fmt.Fscanln(reader, &n, &m)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to parse maze size: %v\n", err)
		os.Exit(1)
	}

	maze := NewMaze(n, m)

	for i := 0; i < n; i++ {
		str, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to parse maze on line %d: %v\n", i, err)
			os.Exit(1)
		}

		str = strings.TrimSpace(str)
		split := strings.Split(str, " ")

		if len(split) != m {
			fmt.Fprintf(os.Stderr, "Maze line length %d and number of columns %d does not match\n", len(split), m)
			os.Exit(1)
		}

		for j, strNum := range split {
			num, err := strconv.Atoi(strNum)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Failed to parse symbol %d of line %d in maze: %v\n", j, i, err)
				os.Exit(1)
			}

			maze.matrix[i][j] = num
		}
	}

	var start, finish Point
	_, err = fmt.Fscanln(reader, &start.X, &start.Y, &finish.X, &finish.Y)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to parse start and finish: %v\n", err)
		os.Exit(1)
	}

	path, err := maze.FindShPath(start, finish)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to find shortest path: %v\n", err)
		os.Exit(1)
	}

	for _, p := range path {
		fmt.Fprintf(writer, "%d %d\n", p.X, p.Y)
	}
	fmt.Fprintf(writer, ".\n")
}
