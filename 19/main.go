package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Direction int

const (
	UP Direction = iota
	DOWN
	LEFT
	RIGHT
)

type Point struct {
	x, y int
}

func (self Point) Get(grid [][]byte) byte {
	return grid[self.y][self.x]
}

func (self Point) IsValid(grid [][]byte) bool {
	return self.y >= 0 && self.y < len(grid) && self.x >= 0 && self.x < len(grid[self.y]) && self.Get(grid) != ' '
}

func (self Point) Up() Point {
	return Point{
		self.x,
		self.y - 1,
	}
}

func (self Point) Down() Point {
	return Point{
		self.x,
		self.y + 1,
	}
}

func (self Point) Left() Point {
	return Point{
		self.x - 1,
		self.y,
	}
}

func (self Point) Right() Point {
	return Point{
		self.x + 1,
		self.y,
	}
}

func navigate(grid [][]byte) ([]byte, int) {
	letters := make([]byte, 0)
	current := Point{strings.IndexByte(string(grid[0]), byte('|')), 0}

	visited := map[Point]bool{}
	direction := DOWN
	steps := 0
	for {
		visited[current] = true
		stop := false
		char := current.Get(grid)
		switch char {
		case '|', '-':
			current, stop = continueAlongPath(grid, current, direction)
		case '+':
			current, direction, stop = navigateJunction(grid, current, visited)
		default:
			letters = append(letters, char)
			current, stop = continueAlongPath(grid, current, direction)
		}
		if stop {
			break
		} else {
			steps++
		}
	}
	return letters, steps
}

func continueAlongPath(grid [][]byte, current Point, direction Direction) (next Point, stop bool) {
	stop = false
	next = current
	switch direction {
	case UP:
		next = current.Up()
	case DOWN:
		next = current.Down()
	case LEFT:
		next = current.Left()
	case RIGHT:
		next = current.Right()
	}
	if !current.IsValid(grid) {
		stop = true
	}
	return
}

func navigateVertical(grid [][]byte, current Point, visited map[Point]bool) (next Point, direction Direction, stop bool) {
	stop = false
	next = current
	a, b := current.Up(), current.Down()
	if a.IsValid(grid) {
		if b.IsValid(grid) {
			if _, ok := visited[a]; !ok {
				next = a
				direction = UP
			} else if _, ok := visited[b]; !ok {
				next = b
				direction = DOWN
			} else {
				stop = true
			}
		} else {
			if _, ok := visited[a]; !ok {
				next = a
				direction = UP
			} else {
				stop = true
			}
		}
	} else if b.IsValid(grid) {
		if _, ok := visited[b]; !ok {
			next = b
			direction = DOWN
		} else {
			stop = true
		}
	} else {
		stop = true
	}
	return
}

func navigateHorizontal(grid [][]byte, current Point, visited map[Point]bool) (next Point, direction Direction, stop bool) {
	stop = false
	next = current
	a, b := current.Left(), current.Right()
	if a.IsValid(grid) {
		if b.IsValid(grid) {
			if _, ok := visited[a]; !ok {
				next = a
				direction = LEFT
			} else if _, ok := visited[b]; !ok {
				next = b
				direction = RIGHT
			} else {
				stop = true
			}
		} else {
			if _, ok := visited[a]; !ok {
				next = a
				direction = LEFT
			} else {
				stop = true
			}
		}
	} else if b.IsValid(grid) {
		if _, ok := visited[b]; !ok {
			next = b
			direction = RIGHT
		} else {
			stop = true
		}
	} else {
		stop = true
	}
	return
}

func navigateJunction(grid [][]byte, current Point, visited map[Point]bool) (next Point, direction Direction, stop bool) {
	stop = false
	direction = UP
	next = current
	nextVertical, directionVertical, stopVertical := navigateVertical(grid, current, visited)
	nextHorizontal, directionHorizontal, stopHorizontal := navigateHorizontal(grid, current, visited)
	if stopVertical {
		if stopHorizontal {
			stop = true
		} else {
			next = nextHorizontal
			direction = directionHorizontal
		}
	} else if stopHorizontal {
		next = nextVertical
		direction = directionVertical
	} else {
		stop = true
	}
	return
}

func asMaze(input string) [][]byte {
	lines := strings.FieldsFunc(input, func(r rune) bool { return r == '\n' })
	maze := make([][]byte, len(lines))
	for i, line := range lines {
		maze[i] = []byte(line)
	}
	return maze
}

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	maze := asMaze(string(input))
	letters, steps := navigate(maze)
	fmt.Println(string(letters))
	fmt.Println(steps)
}
