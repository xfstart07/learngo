// Author: xufei
// Date: 2018/11/19

package main

import (
	"fmt"
	"os"
)

func readMaze(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	var row, col int
	fmt.Fscanf(file, "%d %d", &row, &col)

	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}

	return maze
}

type point struct {
	i,j int
}

var dirs = []point{{-1,0}, {0,-1}, {1, 0}, {0, 1}}

func (p point) add(dir point) point {
	return point{p.i+dir.i, p.j+dir.j}
}

func (p point) at(maze [][]int) (int, bool) {
	if p.i < 0 || p.i >= len(maze) {
		return 0, false
	}
	if p.j < 0 || p.j >= len(maze[p.i]) {
		return 0, false
	}

	return maze[p.i][p.j], true
}

func walk(maze [][]int, start, end point) [][]int {
	steps := make([][]int, len(maze))
	for i := range maze {
		steps[i] = make([]int, len(maze[i]))
	}

	Q := []point{start}

	for len(Q) > 0 {
		cur := Q[0]
		Q = Q[1:]

		if cur == end {
			break
		}

		for _, dir := range dirs {
			// cur + dir
			next := cur.add(dir)

			val, ok := next.at(maze)
			if !ok || val == 1 {
				continue
			}

			val, ok = next.at(steps)
			if !ok || val > 0 {
				continue
			}

			Q = append(Q, next)
			val, _ = cur.at(steps)
			steps[next.i][next.j] = val + 1
		}
	}

	return steps
}

func main() {
	maze := readMaze("maze/maze.in")

	steps := walk(maze, point{0,0}, point{len(maze)-1, len(maze[0])-1})

	for _, row := range steps {
		for _, val := range row {
			fmt.Printf("%3d", val)
		}
		fmt.Println()
	}
}
